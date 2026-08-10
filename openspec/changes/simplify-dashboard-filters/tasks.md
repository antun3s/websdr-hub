## 1. Remover filtro de frequência (HTML + JS)

- [x] 1.1 Remover o grupo do filtro de frequência do HTML (`#f-freq` e seu `<div class="filter-group">`)
- [x] 1.2 Remover variável `fFreq` e listener `input` no JS
- [x] 1.3 Remover lógica de contagem do filtro de frequência em `renderStations()` (contagem de `activeFilters`)
- [x] 1.4 Remover lógica de filtragem por frequência em `renderStations()` (bloco `if (freqHz > 0 ...)`)

## 2. Remover filtro de linguagem (HTML + JS)

- [x] 2.1 Remover o grupo do filtro de linguagem do HTML (`#f-lang` e seu `<div class="filter-group">`)
- [x] 2.2 Remover variável `fLang` e listener `change` no JS
- [x] 2.3 Remover lógica de contagem do filtro de linguagem em `renderStations()`
- [x] 2.4 Remover lógica de filtragem por linguagem em `renderStations()` (bloco `if (langVal && ...)`)
- [x] 2.5 Remover `populateDropdowns()` do idioma (loop `langs`)

## 3. Ajustar ocultação do menu no scroll

- [x] 3.1 Adicionar condição `window.innerWidth >= 768` para não aplicar colapso em desktop
- [x] 3.2 Regenerar `dist/` com `go run ./cmd/websdrctl build`

## 4. Verificação

- [x] 4.1 Servidor em :8080 exibe dashboard sem quebras
- [x] 4.2 Filtros de frequência e linguagem não aparecem na UI
- [x] 4.3 Menu de filtros não colapsa ao scrollar em desktop (≥ 768px)
- [x] 4.4 Git status limpo