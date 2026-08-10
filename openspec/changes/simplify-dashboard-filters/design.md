## Context

O dashboard é HTML+CSS+JS vanilla em arquivo único (`web/index.html`). Filtros e comportamento de scroll são implementados inline no JS com manipulação direta de DOM. Não há framework, build step nem dependências.

## Goals / Non-Goals

**Goals:**
- Remover elementos do filtro de frequência do HTML (`#f-freq`)
- Remover elementos do filtro de linguagem do HTML (`#f-lang`)
- Remover lógica JS que processa esses filtros
- Remover contagem dos filtros no contador de filtros ativos
- Restringir ocultação do menu `.filters` no scroll apenas para viewports < 768px

**Non-Goals:**
- Nenhuma mudança no backend, schema de dados ou outros filtros (status, software, país, banda)

## Decisions

**Decisão:** remoção manual (excluir nós HTML + blocos JS) em vez de CSS-only (esconder). Motivo: remover o código morto evita manutenção futura e confusão.

**Decisão:** ocultação no scroll condicionada a `window.innerWidth < 768`, mantendo a classe `.collapsed` e a transição CSS existentes para mobile.

## Risks / Trade-offs

- Após remover os filtros, `populateDropdowns()` e `populateSoftwareCheckboxes()` continuam funcionando inalterados
- `dist/index.html` precisa ser regenerado com `go run ./cmd/websdrctl build`