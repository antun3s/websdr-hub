# Contribuindo

Obrigado pelo interesse em contribuir com o websdr-directory! Este documento explica como adicionar novas estações e contribuir com código.

## Adicionando uma nova estação

1. Crie um arquivo `data/stations/<id>.yaml` seguindo o modelo abaixo:

```yaml
id: br-saopaulo-py2abc           # slug estável e imutável (kebab-case)
name: Nome da Estação
url: http://exemplo.com:8073/
software: websdr                  # websdr|kiwisdr|openwebrx|phantomsdr|other
location:
  country: BR                     # ISO 3166-1 alpha-2
  city: São Paulo
  coordinates: [-23.55, -46.63]   # [lat, lon] — máx. 3 casas decimais
languages: [pt]                   # ISO 639-1
coverage:
  - name: HF
    start_hz: 0
    stop_hz: 30000000
    modes: [AM, LSB, USB, CW]
max_users: 8
operator: PY2XXX                  # apenas indicativo — sem nome ou e-mail
added_at: 2026-08-09
```

2. Execute a validação local:

```bash
go build -o bin/websdrctl ./cmd/websdrctl
./bin/websdrctl validate
```

3. Se a validação passar sem erros, abra um Pull Request.

## Regras do catálogo

- **Privacidade:** Nunca inclua nomes, e-mails ou endereços de operadores. Use apenas o indicativo de radioamador no campo `operator`.
- **Precisão:** Limite coordenadas a ~100 m de precisão (máx. 3 casas decimais).
- **Verificação:** Verifique se a estação está realmente online antes de submeter. O health check automatizado confirma depois.
- **Slug imutável:** O `id` deve ser estável — ele é a âncora do histórico de status. Use o formato `<cidade>-<indicativo>` ou `<país>-<cidade>-<indicativo>`.
- **Frequências:** Sempre use Hz inteiro (`start_hz`, `stop_hz`), nunca MHz em float.
- **Remoção:** Para remover uma estação, abra uma Issue. Não envie PR removendo estações sem justificativa.

## Contribuindo com código

1. Abra uma Issue descrevendo a mudança proposta antes de começar.
2. Crie um branch a partir de `main`.
3. Faça as alterações e execute os testes:

```bash
go test ./...
go build -o bin/websdrctl ./cmd/websdrctl
./bin/websdrctl validate
```

4. Abra um Pull Request.

O CI executa `websdrctl validate` automaticamente em todo PR.

## Licenciamento

- Código: Apache-2.0 (`LICENSE`)
- Dados do catálogo: ODbL 1.0 (`LICENSE-DATA`)

Ao contribuir, você concorda em licenciar seu código sob Apache-2.0 e os dados sob ODbL 1.0.