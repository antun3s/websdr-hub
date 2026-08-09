## 1. Modelo de Dados e Catálogo

- [x] 1.1 Implementar struct Station com todos os campos obrigatórios (id, name, url, software, location, coverage, etc.)
- [x] 1.2 Implementar validação de estações (slug, URL, ISO país, coordenadas, bandas, data, campos desconhecidos)
- [x] 1.3 Implementar carregamento YAML do diretório `data/stations/` com verificação de nome de arquivo vs id
- [x] 1.4 Implementar detecção de duplicatas (ID e URL normalizada)
- [x] 1.5 Criar arquivos YAML de exemplo (3 estações: NL, DE, BR)

## 2. Health Check

- [x] 2.1 Implementar probe HTTP com timeout configurável, coleta de status code e latência
- [x] 2.2 Implementar lógica de retry (2 tentativas com intervalo de 5s entre elas)
- [x] 2.3 Implementar controle de concorrência via semáforo (default 10 simultâneas)
- [x] 2.4 Configurar User-Agent identificado com URL do projeto
- [x] 2.5 Configurar HTTP client com limite de 5 redirects
- [x] 2.6 Escrever testes unitários para cenários: online na primeira, retry recupera, offline após duas falhas, host inexistente, concorrência respeitada

## 3. Status Tracker

- [x] 3.1 Implementar load do status.json anterior (arquivo local e HTTP URL)
- [x] 3.2 Implementar merge de resultados: online reseta consecutive_failures, offline incrementa
- [x] 3.3 Implementar preservação de last_online entre execuções
- [x] 3.4 Implementar escrita atômica (temp file + rename)

## 4. API Estática

- [x] 4.1 Implementar geração de `dist/v1/stations.json` a partir do catálogo
- [x] 4.2 Implementar geração de `dist/v1/status.json` a partir do merge de resultados
- [x] 4.3 Implementar criação automática do diretório `dist/v1/` se não existir
- [x] 4.4 Serializar todos os campos do modelo Station no JSON de saída

## 5. CLI

- [x] 5.1 Implementar subcomando `validate` — lê catálogo, valida, reporta todos os erros
- [x] 5.2 Implementar subcomando `build` — gera `stations.json`
- [x] 5.3 Implementar subcomando `check` — health check + merge + `status.json`
- [x] 5.4 Implementar flags: `-concurrency`, `-timeout`, `-retry`, `-vantage`, `-prev`
- [x] 5.5 Implementar semântica de exit code: check exit 0 mesmo com estações offline

## 6. Setup do Projeto

- [x] 6.1 Adicionar arquivo `LICENSE` (Apache-2.0)
- [x] 6.2 Adicionar arquivo `LICENSE-DATA` (ODbL 1.0)

## 7. Automação CI/CD

- [x] 7.1 Criar workflow do GitHub Actions para rodar `websdrctl validate` em todo PR
- [x] 7.2 Criar workflow do GitHub Actions com cron 2×/dia (06:00 e 18:00 UTC) para rodar `check` e publicar no GitHub Pages
- [x] 7.3 Configurar GitHub Pages para servir o diretório `dist/` como raiz do site estático

## 8. Catálogo de Estações

- [x] 8.1 Substituir estações de exemplo (`-example`) por estações reais verificadas
- [x] 8.2 Adicionar `CONTRIBUTING.md` com instruções para contribuir novas estações via PR