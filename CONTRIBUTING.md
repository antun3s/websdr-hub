# Contribuindo

Obrigado pelo interesse em contribuir com o websdr-hub! Este documento explica como adicionar novas estações e contribuir com código.

## Adicionando uma nova estação

1. Crie um arquivo `data/stations/<id>.yaml` seguindo o modelo abaixo:

```yaml
id: br-saopaulo-py2abc           # slug estável e imutável (kebab-case)
name: Nome da Estação
url: http://exemplo.com:8073/
software: websdr                  # websdr|kiwisdr|openwebrx|phantomsdr|other
location:
  country: BR                     # ISO 3166-1 alpha-2
  city: São Paulo
  coordinates: [-23.55, -46.63]   # [lat, lon] — máx. 3 casas decimais
languages: [pt]                   # ISO 639-1
coverage:
  - name: HF
    start_hz: 0
    stop_hz: 30000000
    modes: [AM, LSB, USB, CW]
max_users: 8
operator: PY2XXX                  # apenas indicativo — sem nome ou e-mail
added_at: 2026-08-09
```

2. Execute a validação local:

```bash
go build -o bin/websdrctl ./cmd/websdrctl
./bin/websdrctl validate
```

3. Se a validação passar sem erros, abra um Pull Request.

## Regras do catálogo

- **Privacidade:** Nunca inclua nomes, e-mails ou endereços de operadores. Use apenas o indicativo de radioamador no campo `operator`.
- **Precisão:** Limite coordenadas a ~100 m de precisão (máx. 3 casas decimais).
- **Verificação:** Verifique se a estação está realmente online antes de submeter. O health check automatizado confirma depois.
- **Slug imutável:** O `id` deve ser estável — ele é a âncora do histórico de status. Use o formato `<cidade>-<indicativo>` ou `<país>-<cidade>-<indicativo>`.
- **Frequências:** Sempre use Hz inteiro (`start_hz`, `stop_hz`), nunca MHz em float.
- **Remoção:** Para remover uma estação, abra uma Issue. Não envie PR removendo estações sem justificativa.

## Contribuindo com código

1. Abra uma Issue descrevendo a mudança proposta antes de começar.
2. Crie um branch a partir de `main`.
3. Faça as alterações e execute os testes:

```bash
go test ./...
go build -o bin/websdrctl ./cmd/websdrctl
./bin/websdrctl validate
```

4. Abra um Pull Request.

O CI executa `websdrctl validate` automaticamente em todo PR.

## Como funciona

| Camada | Onde vive | Quem escreve |
|---|---|---|
| Catálogo curado | `data/stations/*.yaml` na `main` | Humanos, via PR |
| Status + API | GitHub Pages (`dist/`) | Bot, 2×/dia (06:00 e 18:00 UTC) |

Não existe banco de dados nem branch de estado. A rodada de health check lê o
`status.json` publicado, recalcula e republica — o site *é* o armazenamento.

## Uso local

```bash
go build -o bin/websdrctl ./cmd/websdrctl

./bin/websdrctl validate                     # valida o catálogo
./bin/websdrctl build                        # gera dist/v1/stations.json
./bin/websdrctl check -prev dist/v1/status.json   # consulta e gera status.json
```

Flags de `check`: `-concurrency` (padrão 10), `-timeout` (10s por tentativa),
`-retry` (5s entre tentativas), `-vantage` (identifica o ponto de observação),
`-prev` (caminho local **ou** URL http).

## Verificação de disponibilidade

Duas rodadas por dia (06:00 e 18:00 UTC). Dentro de **cada rodada**, cada
estação recebe até duas tentativas, com 5s entre elas — basta uma responder
para contar como online. São dois eixos independentes: a frequência (2×/dia)
decide *quando* checar; o retry (2 tentativas) decide se uma falha isolada
dentro daquela janela já marca a estação como offline.

O retry importa mais justamente por rodar só 2×/dia: sem ele, um timeout
isolado de link residencial deixaria a estação marcada errada por ~12h, até
a próxima rodada. Com ele, só falhas persistentes (as duas tentativas
falham) viram `offline`.

O runner do GitHub Actions é um único ponto de observação em datacenter. Uma
estação marcada offline pode estar apenas bloqueando aquele range de IP;
`consecutive_failures` no `status.json` é o sinal para investigar antes de
remover algo do catálogo.

Etiqueta de rede: `User-Agent` identificado com a URL do projeto, no máximo 10
consultas simultâneas, timeout curto. Operador que quiser sair do diretório,
basta abrir uma issue.

## Licenciamento

- Código: Apache-2.0 (`LICENSE`)
- Dados do catálogo: ODbL 1.0 (`LICENSE-DATA`)

Ao contribuir, você concorda em licenciar seu código sob Apache-2.0 e os dados sob ODbL 1.0.