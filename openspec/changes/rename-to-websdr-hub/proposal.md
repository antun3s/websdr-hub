## Why

O repositório foi movido de `github.com/websdrdir/websdr-directory` para `github.com/antun3s/websdr-hub`. O nome mais curto (`websdr-hub`) reflete que o projeto é um hub central de receptores WebSDR com curadoria + monitoramento. Todas as referências internas (module path, imports, URLs de CI, docs, footer do dashboard) precisam ser atualizadas para apontar para o novo repositório.

## What Changes

- **go.mod**: module path de `github.com/websdrdir/websdr-directory` → `github.com/antun3s/websdr-hub`
- **Go imports**: todos os `import` que referenciam o module path antigo
- **User-Agent**: string de identificação nas requisições HTTP de health check
- **README.md**: título, badges, URLs de exemplo
- **CONTRIBUTING.md**: nome do projeto
- **CI workflow** (`.github/workflows/check.yml`): URL do status anterior no GitHub Pages
- **HTML** (`web/index.html`, `dist/index.html`): link do footer
- **Specs arquivados**: URLs de exemplo em especificações antigas (apenas cosmético, sem impacto funcional)

Nenhuma mudança de comportamento — apenas identificadores e URLs.

## Capabilities

### New Capabilities

Nenhuma — esta change não introduz novas capacidades.

### Modified Capabilities

Nenhuma — nenhum requisito de especificação existente é alterado.

## Impact

- **15 arquivos** em 6 categorias: module path, imports Go, User-Agent, docs, CI, HTML
- Nenhum impacto em funcionalidade, API pública, schemas de dados ou protocolo
- Após a mudança, `go build` deve continuar funcionando sem alterações
- O GitHub Pages será servido de `https://antun3s.github.io/websdr-hub/...` em vez de `...websdr-directory/...`