## Context

15 arquivos referenciam `websdr-directory` ou `websdrdir` como module path, User-Agent, URL de CI, docs ou footer do dashboard. A mudança é puramente mecânica — sed/awk em cada categoria.

## Goals / Non-Goals

**Goals:**
- Module path e imports Go apontarem para `github.com/antun3s/websdr-hub`
- User-Agent identificar o novo repositório
- URLs de CI e docs refletirem o novo nome
- `go build` funcional após a mudança

**Non-Goals:**
- Nenhuma mudança de comportamento, lógica, schemas ou API
- Nenhuma migração de dados

## Decisions

**Decisão:** substituição mecânica com `sed` por categoria de arquivo, não um `sed` global. Motivo: cada categoria tem padrão de substituição diferente (ex: imports Go são `github.com/websdrdir/websdr-directory/...` → `github.com/antun3s/websdr-hub/...`, já URLs de CI são `websdr-directory` → `websdr-hub`).

**Decisão:** `go.mod` editado manualmente (module path), `go.sum` regenerado com `go mod tidy`. Motivo: `go.sum` contém hashes — não é seguro usar sed.

## Risks / Trade-offs

- **`dist/index.html` é gerado a partir de `web/index.html`** — após alterar a fonte, o `dist/` precisa ser regenerado com `go build && go run ./cmd/websdrctl build` para manter consistência.
- **Arquivos de especificações arquivadas** (`openspec/changes/archive/...`) serão atualizados por consistência, mas não afetam o build.