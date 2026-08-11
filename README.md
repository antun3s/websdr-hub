# websdr-hub

Diretório aberto de receptores WebSDR / KiwiSDR / OpenWebRX com monitoramento
de disponibilidade. O catálogo é curado por Pull Request, o status é verificado
duas vezes por dia e tudo é publicado como API estática.

- Catálogo: `https://<user>.github.io/websdr-hub/v1/stations.json`
- Status: `https://<user>.github.io/websdr-hub/v1/status.json`

## Como funciona

| Camada | Onde vive | Quem escreve |
|---|---|---|
| Catálogo curado | `data/stations/*.yaml` na `main` | Humanos, via PR |
| Status + API | GitHub Pages (`dist/`) | Bot, 2×/dia (06:00 e 18:00 UTC) |

Não existe banco de dados nem branch de estado. A rodada de health check lê o
`status.json` publicado, recalcula e republica — o site *é* o armazenamento.
Isso mantém o `git log` limpo para quem contribui com dados.

## Uso local

```bash
go build -o bin/websdrctl ./cmd/websdrctl

./bin/websdrctl validate                     # valida o catálogo
./bin/websdrctl build                        # gera dist/v1/stations.json
./bin/websdrctl check -prev dist/v1/status.json   # consulta e gera status.json
```

Flags relevantes de `check`: `-concurrency` (padrão 10), `-timeout` (10s por
tentativa), `-retry` (5s entre tentativas), `-vantage` (identifica o ponto de
observação), `-prev` (caminho local **ou** URL http).

## Modelo de dados

Um arquivo por estação, nomeado `<id>.yaml`:

```yaml
id: nl-enschede-pa3fwm            # slug estável e imutável — âncora do histórico
name: University of Twente WebSDR
url: http://websdr.ewi.utwente.nl:8901/
software: websdr                  # websdr|kiwisdr|openwebrx|phantomsdr|other
location:
  country: NL                     # ISO 3166-1 alpha-2
  city: Enschede
  coordinates: [52.239, 6.857]    # [lat, lon] — máx. 3 casas decimais
languages: [nl, en]               # ISO 639-1
coverage:
  - name: HF
    start_hz: 0                   # sempre Hz inteiro, nunca MHz em float
    stop_hz: 29160000
    modes: [AM, LSB, USB, CW]
max_users: 300
operator: PA3FWM                  # apenas indicativo — sem nome ou e-mail
added_at: 2026-08-08
```

**Privacidade:** limite as coordenadas a ~100 m de precisão e nunca inclua
dado pessoal do operador além do indicativo. Muitos receptores ficam na
residência de quem os opera.

## Verificação de disponibilidade

Duas rodadas por dia (06:00 e 18:00 UTC). Dentro de **cada rodada**, cada
estação recebe até duas tentativas, com 5s entre elas — basta uma responder
para contar como online. São dois eixos independentes: a frequência (2×/dia)
decide *quando* checar; o retry (2 tentativas) decide se uma falha isolada
dentro daquela janela já marca a estação como offline.

O retry importa mais justamente por rodar só 2×/dia: sem ele, um timeout
isolado de link residencial deixaria a estação marcada errada por ~12h, até
a próxima rodada. Com ele, só falhas persistentes (as duas tentativas
falham) viram `offline`.

O runner do GitHub Actions é um único ponto de observação em datacenter. Uma
estação marcada offline pode estar apenas bloqueando aquele range de IP;
`consecutive_failures` no `status.json` é o sinal para investigar antes de
remover algo do catálogo.

Etiqueta de rede: `User-Agent` identificado com a URL do projeto, no máximo 10
consultas simultâneas, timeout curto. Operador que quiser sair do diretório,
basta abrir uma issue.

## Contribuindo

1. Crie `data/stations/<id>.yaml` seguindo o modelo acima.
2. Rode `./bin/websdrctl validate` — o CI roda exatamente o mesmo comando.
3. Abra a PR.

A validação verifica formato de slug, URL, ISO de país, faixa de coordenadas,
`start_hz < stop_hz`, data, campos desconhecidos (pega typo em chave),
id/URL duplicados e se o nome do arquivo bate com o `id`. Todos os problemas
saem de uma vez.

## Licenças

- Código: Apache-2.0 (`LICENSE`)
- Dados do catálogo: ODbL 1.0 (`LICENSE-DATA`)

Licenciamento duplo é intencional: sem ele ninguém sabe se pode reusar o
catálogo, que é justamente o problema dos diretórios existentes.

