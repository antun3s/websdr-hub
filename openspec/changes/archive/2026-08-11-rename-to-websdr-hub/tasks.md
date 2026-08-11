## 1. Module Path e Imports Go

- [x] 1.1 Atualizar `go.mod`: `module github.com/websdrdir/websdr-directory` → `github.com/antun3s/websdr-hub`
- [x] 1.2 Atualizar imports em `cmd/websdrctl/main.go` (3 imports)
- [x] 1.3 Atualizar imports em `internal/check/check.go` (1 import + User-Agent)
- [x] 1.4 Atualizar imports em `internal/check/check_test.go` (1 import)
- [x] 1.5 Atualizar imports em `internal/status/status.go` (1 import)
- [x] 1.6 Executar `go mod tidy` para regenerar `go.sum`

## 2. Documentação

- [x] 2.1 Atualizar `README.md`: título, URLs de catálogo/status
- [x] 2.2 Atualizar `CONTRIBUTING.md`: nome do projeto

## 3. CI e HTML

- [x] 3.1 Atualizar `.github/workflows/check.yml`: URL do GitHub Pages
- [x] 3.2 Atualizar `web/index.html`: link do footer
- [x] 3.3 Regenerar `dist/`: `go build && go run ./cmd/websdrctl build` (para sincronizar `dist/index.html`)

## 4. Especificações Arquivadas (cosmético)

- [x] 4.1 Atualizar `openspec/specs/web-dashboard/spec.md`: URLs de exemplo
- [x] 4.2 Atualizar `openspec/specs/health-checker/spec.md`: User-Agent
- [x] 4.3 Atualizar arquivos em `openspec/changes/archive/` (3 arquivos com URLs/User-Agent)

## 5. Verificação

- [x] 5.1 `go vet ./...` sem erros
- [x] 5.2 `go build ./cmd/websdrctl` bem-sucedido
- [x] 5.3 `git diff` revisado — apenas as substituições esperadas
- [x] 5.4 `git status` limpo após o commit final