## Context

Dashboard atual tem 5 filtros no HTML/JS mas bug no CSS oculta os selects (status, país, banda, idioma). Só o filtro de software (checkboxes) fica visível. O filtro por frequência não existe.

## Goals / Non-Goals

**Goals:**
- Tornar todos os 5 filtros visíveis e funcionais
- Adicionar filtro por frequência (kHz) que intersecciona com as bandas de cada estação

**Non-Goals:**
- Não alterar API, dados ou código Go

## Decisions

### Filtro por frequência como input numérico
**Decisão:** Campo de texto que aceita valores em kHz. Ao digitar, verifica se alguma banda da estação tem `start_hz <= freq_hz <= stop_hz`.

**Racional:** Radioamadores pensam em kHz (ex: 7100 para 40m, 144000 para 2m). O input aceita números inteiros e é combinável com os demais filtros (AND).

### Input nativo (sem slider)
**Decisão:** Input type="number" simples, sem slider de range.

**Alternativa:** Slider com range 0-30000000. Rejeitado porque a escala é muito grande (0 a 30 MHz) para um slider ser preciso — um input numérico é mais prático.

## Risks / Trade-offs

- **[Unidade]** Usuário pode digitar em MHz em vez de kHz. → Colocar sufixo "kHz" e placeholder "ex: 7100" no input.