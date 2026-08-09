## Why

Diretórios de WebSDR existem mas não são abertos, versionados ou monitorados ativamente. Operadores e entusiastas de rádio precisam de um catálogo confiável, que aceite contribuições da comunidade via Pull Request e indique se uma estação está realmente online — não apenas listada. Com licenciamento claro (código Apache-2.0, dados ODbL), qualquer um pode reusar o catálogo sem ambiguidade legal.

## What Changes

- Modelo de dados YAML para estações WebSDR/KiwiSDR/OpenWebRX com validação rigorosa (slug, URL, ISO de país, coordenadas, bandas de frequência, campos desconhecidos, duplicatas)
- Health check HTTP com duas tentativas por estação (retry de 5s) e controle de concorrência (máx 10 simultâneas) — etiqueta de rede embutida (User-Agent identificado, timeout curto)
- Rastreamento de falhas consecutivas entre execuções (`consecutive_failures` persiste no `status.json`)
- Geração de API estática: `stations.json` (catálogo) e `status.json` (disponibilidade), ambos publicáveis via GitHub Pages
- CLI único (`websdrctl`) com subcomandos: `validate`, `build`, `check`

## Capabilities

### New Capabilities
- `station-catalog`: Modelo de dados, validação e carregamento de estações WebSDR a partir de arquivos YAML no diretório `data/stations/`
- `health-checker`: Motor de health check HTTP com retry, controle de concorrência e coleta de latência/status code
- `status-tracker`: Mesclagem e persistência de status entre execuções, com rastreamento de falhas consecutivas e escrita atômica
- `static-api`: Geração dos arquivos JSON estáticos — catálogo de estações e status de disponibilidade — na estrutura `dist/v1/`
- `cli-tool`: Interface de linha de comando unificada com subcomandos `validate`, `build` e `check`

### Modified Capabilities
<!-- Nenhuma — primeiro ciclo de desenvolvimento -->

## Impact

- **Código existente**: `cmd/websdrctl/`, `internal/catalog/`, `internal/check/`, `internal/status/` — código Go do MVP já implementado
- **Dados**: `data/stations/*.yaml` — 3 estações de exemplo (1 real + 2 de exemplo)
- **Build**: Go 1.22, dependência única `github.com/goccy/go-yaml`
- **Publicação**: Saída em `dist/v1/` pronta para deploy via GitHub Pages
- **CI futuro**: GitHub Actions com cron 2×/dia (06:00 e 18:00 UTC) — fluxo de CI ainda não implementado