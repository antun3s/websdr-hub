## Why

O README atual é extenso e mistura documentação técnica de infraestrutura (health check, GitHub Actions, modelo de dados completo) com o guia de contribuição. Para um visitante que quer apenas adicionar uma estação, a informação relevante está enterrada. Precisamos de um README sucinto que foque em ensinar pessoas a contribuir com novas estações, deixando detalhes operacionais para CONTRIBUTING.md.

## What Changes

- Reescrever README.md com foco em: o que é o projeto, link para demonstração ao vivo, e tutorial passo a passo para adicionar uma nova estação
- Mover detalhes de verificação de disponibilidade, health check, GitHub Actions e licenças para CONTRIBUTING.md (que já existe)
- Manter modelo de dados no README, mas de forma mais enxuta e direta
- Remover/reduzir seções de uso local (`go build`, flags) que não são relevantes para contribuidores de catálogo
- Adicionar link para https://websdr.antunes.pro/ como demonstração ao vivo

## Capabilities

### New Capabilities

- `documentation`: Documentação do projeto — README focado em onboarding de contribuidores

### Modified Capabilities

<!-- Nenhuma spec existente tem requisitos sendo alterados. É apenas documentação. -->

## Impact

- README.md será reescrito
- CONTRIBUTING.md pode receber conteúdo movido do README atual
- Nenhuma mudança em código, API, esquema de dados ou infraestrutura