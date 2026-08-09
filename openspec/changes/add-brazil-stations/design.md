## Context

Catálogo atual com 1 estação (NL). O site websdr.com.br lista ~35 estações brasileiras de 3 sistemas (WebSDR Clássico, KiwiSDR, OpenWebSDR) com dados de cidade, URL, software e bandas aproximadas. Callsigns dos operadores não estão disponíveis no site.

## Goals / Non-Goals

**Goals:**
- Adicionar ~35 estações brasileiras ao catálogo via arquivos YAML
- Extrair dados manualmente: cidade, UF, URL, software, bandas
- Callsign vazio onde não disponível, para preenchimento posterior

**Non-Goals:**
- Nenhuma alteração em código Go, HTML, CSS ou workflows
- Nenhuma scraping automatizado (extração manual)

## Decisions

### Extração manual (não automatizada)
**Decisão:** Cada estação é inserida manualmente como arquivo YAML.

**Racional:** Apenas ~35 estações, e a extração automatizada demandsria heurísticas complexas para traduzir descrições em português ("80m e 40m") para faixas de frequência estruturadas. O trabalho manual garante precisão.

### Nomenclatura dos arquivos
**Decisão:** `br-<cidade>-<indicativo>.yaml` quando o indicativo é conhecido. Quando não, `br-<cidade>-<uf>-<software>.yaml` (ex: `br-saopaulo-sp-websdr.yaml`). Se houver múltiplas estações na mesma cidade, desambiguar por software ou nome DNS (ex: `br-pardinho-sp-websdr.yaml`, `br-pardinho-sp-kiwi1.yaml`).

### Coordenadas aproximadas
**Decisão:** Usar coordenadas da cidade (centro aproximado) com 1-2 casas decimais, não da estação exata (~10 km de precisão).

**Racional:** O site não fornece coordenadas. Privacidade do operador é prioridade.

### Bandas de frequência
**Decisão:** Traduzir descrições like "80m e 40m" para coverage com start_hz/stop_hz aproximados usando tabela de bandas de radioamador. Quando a descrição é vaga ("Scanner com varias bandas", "VHF, SAT e AVIAÇÃO"), usar a banda mais próxima conhecida ou cobrir a faixa genérica.

## Risks / Trade-offs

- **[Coordenadas imprecisas]** Coordenadas são da cidade, não da estação. → Precisão de ~10 km é aceitável para um diretório global; operadores podem ajustar depois via PR.
- **[Bandas aproximadas]** "80m" pode ser 3.5-4.0 MHz ou 3.5-3.8 MHz dependendo da região. → Usar a faixa padrão brasileira/local quando possível.
- **[Callsign ausente]** Sem callsign, a estação é menos útil para identificação. → Campo vazio permite preenchimento futuro via PR.