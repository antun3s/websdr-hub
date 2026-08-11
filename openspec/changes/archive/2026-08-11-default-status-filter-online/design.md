## Context

O dashboard é HTML+CSS+JS vanilla em arquivo único (`web/index.html`). O filtro de status é um `<select>` com opções "All" (value `""`), "Online" (`"online"`) e "Offline" (`"offline"`). Atualmente o `<option value="">` tem o atributo `selected` implícito (primeira opção).

## Goals / Non-Goals

**Goals:**
- `<select id="f-status">` inicia com "Online" selecionado por padrão
- `renderStations()` filtra por `statusVal === "online"` na primeira execução
- O contador de filtros ativos reflete que o filtro de status está ativo (1 filtro) desde o início

**Non-Goals:**
- Nenhuma mudança no backend, schema de dados, outros filtros, layout ou comportamento de scroll
- Nenhuma persistência da escolha do usuário (ex.: localStorage) — apenas o padrão na carga da página

## Decisions

**Decisão:** adicionar `selected` no `<option value="online">` em vez de setar `fStatus.value = "online"` no JS. Motivo: mais simples, zero JS extra, o valor correto já está no DOM desde o parse inicial.

**Decisão:** não alterar a lógica de contagem de filtros ativos. Com "Online" como padrão, `statusVal` será `"online"` (truthy), então `activeFilters` começará em 1 — o que é correto, pois o filtro está ativo.

## Risks / Trade-offs

- Usuários que queriam ver "All" agora precisam mudar o filtro manualmente. Como a maioria quer "Online" (motivação da mudança), o trade-off é aceitável.
- `dist/index.html` precisa ser regenerado com `go run ./cmd/websdrctl build`