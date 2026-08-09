## 1. Estrutura HTML

- [x] 1.1 Criar `dist/index.html` com estrutura base: header, barra de filtros e container de cards
- [x] 1.2 Adicionar seção de filtros: status (dropdown), software (checkboxes), país (dropdown), banda (dropdown), idioma (dropdown)
- [x] 1.3 Adicionar contador de resultados ("N estações / M filtros ativos")
- [x] 1.4 Adicionar indicador de loading e mensagem de erro com botão de retry

## 2. CSS e Layout

- [x] 2.1 Estilizar barra de filtros com layout horizontal (empilha vertical em mobile)
- [x] 2.2 Criar layout de cards em grid responsivo: 1 col mobile, 2 cols tablet, 3+ cols desktop
- [x] 2.3 Estilizar cards: nome, software badge, bandeira, bandas com modos, indicador online/offline, latência
- [x] 2.4 Estilizar indicador de status: verde (online), vermelho (offline), cinza (unknown)
- [x] 2.5 Aplicar tema escuro (dark mode) para conforto visual em uso de rádio

## 3. Lógica de Dados

- [x] 3.1 Implementar fetch de `./v1/stations.json` e `./v1/status.json` via URLs relativas (sem hardcode de origem/host/porta) em paralelo na inicialização
- [x] 3.2 Implementar merge dos dados por station ID (join catalog + status)
- [x] 3.3 Tratar estados: loading, erro (com mensagem + retry), sucesso

## 4. Filtros

- [x] 4.1 Implementar filtro por status (todos / online / offline)
- [x] 4.2 Implementar filtro por software via checkboxes (WebSDR, KiwiSDR, OpenWebRX, PhantomSDR, Other)
- [x] 4.3 Implementar filtro por país via dropdown populado dinamicamente
- [x] 4.4 Implementar filtro por banda de frequência via dropdown populado dinamicamente
- [x] 4.5 Implementar filtro por idioma via dropdown populado dinamicamente
- [x] 4.6 Implementar lógica AND entre filtros (todos os filtros ativos são aplicados simultaneamente)

## 5. Renderização

- [x] 5.1 Implementar função de render que gera cards HTML a partir da lista filtrada
- [x] 5.2 Exibir emoji de bandeira do país com fallback para código ISO
- [x] 5.3 Exibir bandas de frequência com modos em cada card
- [x] 5.4 Exibir status (online/offline) com indicador colorido, latência e timestamp
- [x] 5.5 Atualizar contador de resultados a cada mudança de filtro

## 6. Integração

- [x] 6.1 Garantir que a dashboard funcione independentemente da origem de rede: localhost, IP LAN, domínio público ou GitHub Pages (todas as URLs de API são relativas)
- [x] 6.2 Copiar `index.html` da fonte para `dist/` no subcomando `build` do CLI (para desenvolvimento local)
- [x] 6.3 Testar dashboard localmente com `dist/v1/stations.json` + `dist/v1/status.json` gerados