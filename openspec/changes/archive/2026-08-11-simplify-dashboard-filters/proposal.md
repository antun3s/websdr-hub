## Why

O dashboard tem dois problemas de usabilidade:

1. **Filtros de frequência (kHz) e linguagem** são pouco usados e poluem a interface. Status, software, país e banda já cobrem a maioria dos casos de uso.
2. **Menu de filtros colapsa no scroll** — comportamento pensado para mobile, mas em desktop (viewport > 768px) é irritante: o usuário perde os filtros de vista ao rolar.

## What Changes

- **Remover** filtro "Freq (kHz)" do HTML, JS e CSS
- **Remover** filtro "Language" do HTML, JS e CSS
- **Ocultação no scroll** removida para viewports ≥ 768px; mantida apenas em mobile (< 768px)

## Capabilities

### New Capabilities

Nenhuma — apenas remoção de funcionalidade.

### Modified Capabilities

- `web-dashboard`: requisitos de filtro por frequência e linguagem removidos; comportamento de ocultação dos filtros restrito a mobile

## Impact

- Apenas `web/index.html` (fonte) e `dist/index.html` (regenerado)
- Nenhuma API, schema de dados ou backend afetado