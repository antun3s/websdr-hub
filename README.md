# websdr-hub

Diretório aberto de receptores WebSDR / KiwiSDR / OpenWebRX com monitoramento
de disponibilidade. O catálogo é curado por Pull Request e o status é verificado
duas vezes por dia.

**Demonstração ao vivo:** [https://websdr.antunes.pro/](https://websdr.antunes.pro/)

- Catálogo: `https://<user>.github.io/websdr-hub/v1/stations.json`
- Status: `https://<user>.github.io/websdr-hub/v1/status.json`

## Adicionar uma estação

1. Crie `data/stations/<id>.yaml` seguindo o modelo abaixo
2. Valide localmente: `./bin/websdrctl validate`
3. Abra um Pull Request

O CI valida automaticamente em todo PR. Para detalhes sobre infraestrutura,
health check e licenças, veja [CONTRIBUTING.md](CONTRIBUTING.md).

## Modelo de dados

```yaml
id: nl-enschede-pa3fwm
name: University of Twente WebSDR
url: http://websdr.ewi.utwente.nl:8901/
software: websdr
location:
  country: NL
  city: Enschede
  coordinates: [52.239, 6.857]
languages: [nl, en]
coverage:
  - name: HF
    start_hz: 0
    stop_hz: 29160000
    modes: [AM, LSB, USB, CW]
max_users: 300
operator: PA3FWM
added_at: 2026-08-08
```