## Why

A API estática (`stations.json` + `status.json`) já existe e é publicada via GitHub Pages, mas não há interface visual para consultá-la. Um operador ou entusiasta que acessa o site vê apenas JSON bruto — não consegue filtrar por país, tipo de software, banda de frequência ou status online/offline. Uma dashboard web simples, que consuma a própria API estática, transforma o diretório de arquivos em um catálogo navegável.

## What Changes

- Nova página HTML estática (`index.html`) com JavaScript vanilla que busca `stations.json` e `status.json` da mesma origem e renderiza uma tabela/cards filtrável
- Filtros por: status (online/offline), tipo de software (websdr/kiwisdr/openwebrx/...), país, banda de frequência e idioma
- Interface responsiva e leve, sem dependências de build ou frameworks
- Publicada junto com a API estática em `dist/` (raiz do GitHub Pages)

## Capabilities

### New Capabilities
- `web-dashboard`: Interface web single-page que consome a API estática, mescla dados de catálogo e status, e oferece filtros para navegação

### Modified Capabilities
<!-- Nenhuma — a API estática existente já supre todos os dados necessários -->

## Impact

- **Novo arquivo**: `dist/index.html` (dashboard estática, publicada como raiz do GitHub Pages)
- **Build tool**: O subcomando `build` do CLI pode opcionalmente copiar/gerar o `index.html` para `dist/`
- **Deploy**: CI de health check já publica `dist/` via GitHub Pages — a dashboard é servida automaticamente
- **Dependências**: Nenhuma externa — HTML + CSS + JavaScript vanilla