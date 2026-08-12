## Context

O README atual tem 120+ linhas e mistura onboarding de contribuidores com documentação técnica de infraestrutura (health check, GitHub Actions, flags de CLI, licenças). Isso dificulta que um novo contribuidor encontre rapidamente o que precisa para adicionar uma estação.

## Goals / Non-Goals

**Goals:**
- README sucinto (~30-50 linhas) focado em: o que é o projeto, demonstração ao vivo, tutorial de adição de estação
- Manter modelo de dados (`<id>.yaml`) no README de forma enxuta
- Mover seções de infraestrutura, licenças, health check para CONTRIBUTING.md

**Non-Goals:**
- Alterar qualquer código, pipeline, ou schema de dados
- Mudar a estrutura ou validação do catálogo
- Adicionar/remover seções do CONTRIBUTING.md existente (apenas realocar conteúdo)

## Decisions

1. **README como landing page de contribuição** — O README deve responder "o que é isso?" e "como adiciono minha estação?" em segundos. Detalhes operacionais vão para CONTRIBUTING.md.
2. **Link de demonstração em destaque** — `https://websdr.antunes.pro/` no topo, para que o visitante veja o projeto funcionando antes de contribuir.
3. **Tutorial passo a passo** — Sequência curta: criar arquivo YAML, rodar validação, abrir PR. Sem explicação de flags ou infra.
4. **Modelo de dados enxuto** — Apenas os campos essenciais com exemplo, sem notas de privacidade ou validação (isso fica no CONTRIBUTING.md).

## Risks / Trade-offs

- **Conteúdo duplicado entre README e CONTRIBUTING.md** → Risco baixo, pois CONTRIBUTING.md já existe e receberá o conteúdo movido. Manteremos consistência na revisão.
- **Usuário perde detalhes de validação** → O passo "rode `validate`" no README redireciona para o CONTRIBUTING.md para detalhes, então o fluxo é progressivo.