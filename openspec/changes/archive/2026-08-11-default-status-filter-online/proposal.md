## Why

Ao carregar o dashboard, o filtro de status inicia como "All", mostrando todas as estações — inclusive as offline. Como a maioria dos usuários quer ver apenas estações online, o padrão deveria ser "Online" para uma experiência mais útil desde o primeiro carregamento.

## What Changes

- Alterar o valor padrão do `<select id="f-status">` de `""` (All) para `"online"`
- Atualizar a variável/constante que armazena o valor inicial do filtro de status
- Garantir que `renderStations()` seja chamado com o filtro padrão "online" já ativo na primeira renderização
- Nenhuma mudança no backend, schema de dados ou outros filtros

## Capabilities

### New Capabilities

Nenhuma.

### Modified Capabilities

- `web-dashboard`: o filtro de status inicia como "online" por padrão em vez de "all"

## Impact

- Apenas `web/index.html` (fonte) e `dist/index.html` (regenerado)
- Nenhuma API, schema de dados ou backend afetado