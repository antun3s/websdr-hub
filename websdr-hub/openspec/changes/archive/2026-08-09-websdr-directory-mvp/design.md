## Context

O projeto é um diretório aberto de estações WebSDR com monitoramento de disponibilidade. O catálogo é curado por humanos via Pull Request (arquivos YAML), o status é verificado por um bot 2×/dia, e tudo é publicado como API estática via GitHub Pages. Não há banco de dados — o `status.json` publicado é lido como ponto de partida, recalculado e republicado.

O código atual é um MVP em Go 1.22, organizado em módulos internos (`internal/catalog`, `internal/check`, `internal/status`) com um CLI unificado (`cmd/websdrctl`). A dependência externa única é `go-yaml` para parsing de YAML.

## Goals / Non-Goals

**Goals:**
- Modelo de dados YAML validado rigorosamente, amigável a PRs da comunidade
- Health check HTTP com retry e controle de concorrência, com etiqueta de rede
- Rastreamento de estado entre execuções (`consecutive_failures`) sem banco de dados
- API estática publicável como GitHub Pages (`stations.json` + `status.json`)
- CLI único com subcomandos independentes (validate, build, check)

**Non-Goals:**
- Frontend / UI de consulta ao catálogo (futuro, consumindo a API estática)
- Notificações de offline (pode ser adicionado depois como ação separada)
- Múltiplos vantage points de health check (inicialmente um runner de CI)
- Suporte a outros formatos além de YAML para entrada de dados
- Autenticação ou API com estado no servidor

## Decisions

### YAML como formato de dados de entrada
**Decisão:** Cada estação é um arquivo `.yaml` em `data/stations/`, nomeado `<id>.yaml`.

**Alternativa considerada:** JSON. Rejeitado porque YAML é mais legível para curadores humanos e suporta comentários naturais (embora não usados no schema). A comunidade de rádio amador tem familiaridade com formatos de texto simples.

**Alternativa considerada:** Banco de dados. Rejeitado para o catálogo porque arquivos no repositório permitem revisão em PR, diff limpo no git, e nenhuma infraestrutura adicional. Para o status, um banco criaria um ciclo de deploy desnecessário quando o GitHub Pages funciona como storage gratuito.

### Duas tentativas com 5s entre elas
**Decisão:** Dentro de cada execução de check, cada estação recebe até 2 tentativas HTTP com 5 segundos de intervalo. Se a primeira responder, a segunda é omitida. Só marca offline se ambas falharem.

**Racional:** Com apenas 2 execuções por dia (06:00 e 18:00 UTC), um timeout isolado de link residencial deixaria a estação marcada offline por ~12h. O retry intra-execução mitiga isso sem aumentar a carga de rede (a maioria das estações responde na primeira tentativa).

### Semáforo de concorrência (máx 10 simultâneas)
**Decisão:** O checker usa um channel como semáforo para limitar requisições HTTP simultâneas. Default 10, configurável via `-concurrency`.

**Racional:** Etiqueta de rede — centenas de requisições simultâneas para Raspberry Pis residenciais pareceriam um port scan. O semáforo mantém o throughput razoável (~30s para 300 estações com 2 tentativas) sem sobrecarregar links domésticos.

### status.json como storage (sem banco de dados)
**Decisão:** O estado entre execuções (`consecutive_failures`, `last_online`) é persistido exclusivamente no arquivo `status.json` publicado. A cada execução, o checker lê o arquivo anterior (local ou via HTTP da GitHub Pages), mescla com os resultados novos e escreve atomicamente.

**Racional:** Elimina a necessidade de um banco de dados ou branch de estado. O `git log` da branch `main` permanece limpo (apenas dados de catálogo). O GitHub Pages serve como CDN gratuita para o JSON. Se o Pages estiver fora do ar, a execução parte de estado zerado — pior caso, `consecutive_failures` reseta, o que é aceitável para um MVP.

### Escrita atômica via temp file + rename
**Decisão:** Arquivos JSON são escritos em um arquivo temporário no mesmo diretório e renomeados atomicamente (`os.Rename`).

**Racional:** Em sistemas de arquivos POSIX, `rename` é atômico. Isso evita que um deploy ou leitura concorrente encontre um arquivo JSON truncado pela metade.

### Normalização de URL para detecção de duplicatas
**Decisão:** URLs são comparadas após redução para lowercase e remoção de trailing slash no path. Host + path normalizado é usado como chave de unicidade.

**Racional:** `http://exemplo.com:8073` e `http://exemplo.com:8073/` são a mesma estação. Sem normalização, o sistema permitiria duplicatas que confundiriam o catálogo e o health check.

### Estrutura de saída: dist/v1/
**Decisão:** Os artefatos gerados (`stations.json`, `status.json`) vão em `dist/v1/`, versionados por prefixo de API.

**Racional:** Versionamento de API desde o início permite evolução do schema sem quebrar consumidores. `v1` é o contrato inicial.

### User-Agent identificado
**Decisão:** Toda requisição HTTP de health check inclui `User-Agent: websdr-directory/0.1 (+https://github.com/websdrdir/websdr-directory)`.

**Racional:** Transparência — operadores de WebSDR podem identificar o tráfego nos logs e entrar em contato se houver problemas. É uma prática padrão de etiqueta de rede para bots de monitoramento.

## Risks / Trade-offs

- **[Único vantage point]** O GitHub Actions runner é um único ponto de observação em datacenter. Uma estação bloqueando aquele range de IP aparece como offline mesmo estando funcional. → `consecutive_failures` no status.json sinaliza para investigação humana antes de remover do catálogo.

- **[Estado volátil]** Se o GitHub Pages estiver fora do ar ou o `status.json` for corrompido, o histórico de `consecutive_failures` é perdido naquela execução. → Escrita atômica minimiza corrupção. Perda de estado é tolerável para MVP — o pior caso é reset de contagem de falhas.

- **[Escala do catálogo]** Com milhares de estações, o health check sequencial seria lento. → Concorrência de 10 já cobre ~300 estações em ~30s. Para escalar além, aumentar concorrência ou shard por região geográfica.

- **[Sem notificações]** O sistema detecta offline mas não alerta ninguém. → Fora do escopo do MVP. Pode ser adicionado como ação de CI separada que lê `status.json` e dispara alertas (email, webhook, issue no GitHub).

## Migration Plan

Não aplicável — primeiro ciclo de desenvolvimento. O MVP já está implementado como código inicial. A partir deste ponto, alterações seguem o workflow spec-driven do OpenSpec.