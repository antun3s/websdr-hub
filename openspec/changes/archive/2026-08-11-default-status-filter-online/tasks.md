## 1. Alterar o padrão do filtro de status para "online"

- [x] 1.1 Adicionar `selected` no `<option value="online">` do `<select id="f-status">` no HTML
- [x] 1.2 Verificar que `renderStations()` é chamada após `init()` e filtra corretamente com o novo padrão
- [x] 1.3 Verificar que o contador de filtros ativos mostra "1 filter active" ao carregar

## 2. Verificação

- [x] 2.1 Servidor em :8080 exibe dashboard com filtro "Online" selecionado por padrão
- [x] 2.2 Apenas estações online aparecem na primeira renderização
- [x] 2.3 Mudar para "All" exibe todas as estações
- [x] 2.4 Regenerar `dist/` com `go run ./cmd/websdrctl build`
- [x] 2.5 Git status limpo