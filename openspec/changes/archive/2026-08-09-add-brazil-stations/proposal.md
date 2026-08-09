## Why

O catálogo atual tem apenas 1 estação (NL). O Brasil possui dezenas de WebSDRs ativos listados em https://www.websdr.com.br/ que não estão no diretório. Adicioná-los aumenta significativamente a cobertura do catálogo, especialmente na América do Sul, sem exigir alterações no código — apenas novos arquivos YAML seguindo o schema existente.

## What Changes

- Adicionar ~35 arquivos YAML em `data/stations/` para estações brasileiras extraídas manualmente do websdr.com.br
- Cada arquivo segue o modelo existente: id kebab-case, name, url, software, location (cidade/UF/coords), coverage com bandas aproximadas, operator vazio quando não disponível
- Estações já existentes no catálogo não são alteradas

## Capabilities

### New Capabilities
<!-- Nenhuma — apenas dados, sem novas funcionalidades -->

### Modified Capabilities
<!-- Nenhuma — o schema do catálogo não muda -->

## Impact

- **Dados**: ~35 novos arquivos em `data/stations/` (somente estações BR)
- **Código**: Nenhuma alteração em Go, CSS, HTML ou workflows
- **API**: `stations.json` e `dist/index.html` passarão a exibir as novas estações após o próximo `build`