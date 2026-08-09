## Context

O projeto já gera API estática (`dist/v1/stations.json` + `dist/v1/status.json`) publicada via GitHub Pages. O próximo passo é oferecer uma interface visual para que qualquer pessoa possa navegar pelo catálogo sem precisar ler JSON bruto. A dashboard será um arquivo HTML estático servido na raiz do GitHub Pages (`dist/index.html`), consumindo a API da mesma origem.

## Goals / Non-Goals

**Goals:**
- Single-page application em HTML + CSS + JS vanilla, sem dependências externas ou build step
- Exibir lista de estações com dados do catálogo + status de disponibilidade mesclados
- Filtros combináveis: status (todos/online/offline), tipo de software, país, banda de frequência, idioma
- Layout responsivo (desktop e mobile)
- Publicada automaticamente junto com a API estática

**Non-Goals:**
- Autenticação ou favoritos
- Mapas interativos (pode ser adicionado depois)
- Ordenação por coluna clicável (pode ser adicionado depois)
- Paginação server-side (não necessário para o tamanho atual do catálogo)

## Decisions

### Vanilla HTML + CSS + JS (sem framework)
**Decisão:** Arquivo único `index.html` com CSS e JS embutidos. Nenhuma dependência npm, nenhum build step.

**Alternativa considerada:** React/Vue/Svelte. Rejeitado porque adicionaria complexidade de build e dependências para uma interface de ~200 linhas de JS. O catálogo tem dezenas/centenas de itens, não milhares — performance de renderização com vanilla JS é mais que suficiente.

**Alternativa considerada:** HTMX. Considerado mas rejeitado porque os filtros são puramente client-side (os dados já estão no browser após o primeiro fetch) e HTMX brilharia mais com interações server-side.

### Fetch da mesma origem
**Decisão:** A dashboard faz `fetch('/v1/stations.json')` e `fetch('/v1/status.json')` da mesma origem (GitHub Pages). Sem CORS necessário.

**Racional:** Ambos os arquivos são servidos do mesmo domínio GitHub Pages. Isso simplifica o deploy e elimina configuração de CORS.

### Filtros como dropdowns/checkboxes acima da lista
**Decisão:** Barra de filtros no topo com:
- Status: dropdown (Todos / Online / Offline)
- Software: checkboxes (WebSDR, KiwiSDR, OpenWebRX, PhantomSDR, Other)
- País: dropdown populado dinamicamente a partir dos dados
- Banda: dropdown populado dinamicamente a partir dos dados
- Idioma: dropdown populado dinamicamente a partir dos dados

**Racional:** Dropdowns dinâmicos evitam opções fantasmas (ex: mostrar "KiwiSDR" se nenhuma estação usa). Checkboxes para software porque um operador pode querer ver WebSDR + OpenWebRX simultaneamente.

### Cards em grid responsivo
**Decisão:** Layout de cards (não tabela) com grid CSS responsivo (1 coluna mobile, 2 tablets, 3+ desktop).

**Alternativa considerada:** Tabela HTML. Rejeitada porque cards são mais legíveis em mobile e permitem mostrar mais informações (bandas, idiomas) sem scroll horizontal.

### Estado "carregando" e "erro"
**Decisão:** Enquanto os fetches estão em andamento, exibir indicador de loading. Se o fetch falhar (ex: Pages fora do ar), exibir mensagem de erro com instrução para recarregar.

**Racional:** Feedback visual é essencial para não parecer quebrado. Como é um site estático, não há como "cair" — mas o fetch pode falhar por rede.

## Risks / Trade-offs

- **[Performance com catálogo grande]** Com milhares de estações, o JSON pode ser grande (~500KB+) e o DOM com cards pode ficar pesado. → Para MVP com dezenas/centenas de estações, não é problema. Se escalar, adicionar virtual scrolling ou paginação.
- **[Sem atualização em tempo real]** A dashboard mostra o snapshot do último health check (2×/dia). O status pode estar desatualizado por até 12h. → O timestamp `generated_at` é exibido no rodapé para transparência.
- **[Acessibilidade]** Vanilla JS sem framework pode negligenciar a11y se não for intencional. → Usar elementos semânticos ( `<main>`, `<nav>`, `<article>`) e labels em filtros.