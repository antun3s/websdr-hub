## Why

Os filtros do dashboard existem no HTML/JS (status, país, banda, idioma, software) mas um bug no CSS (`display: none` em selects que nunca é revertido) os torna invisíveis — o usuário vê apenas o filtro de software. Além disso, um filtro por frequência em kHz (ex: "7100" mostra estações que cobrem 40m) é viável e útil para radioamadores.

## What Changes

- **BUGFIX:** Corrigir CSS para exibir os dropdowns de status, país, banda e idioma
- **MELHORIA:** Adicionar campo de input para filtrar por frequência (kHz) — verifica se a estação tem alguma banda que cobre a frequência informada
- **MELHORIA:** Melhorar labels dos filtros com ícones/prefixos visuais

## Capabilities

### New Capabilities
<!-- Nenhuma — apenas modificação no web-dashboard existente -->

### Modified Capabilities
- `web-dashboard`: Adicionar requisito de filtro por frequência

## Impact

- **Código**: Apenas `web/index.html` (CSS + JS) — sem mudanças em Go, API ou dados