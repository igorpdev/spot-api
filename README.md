# SPOT

Mapa curado de São Paulo baseado em afinidade, não em ranking ou nota.

A pergunta central: "onde faz sentido ir, para alguém como você, agora?"

## A ideia

SPOT não é Google Maps nem TripAdvisor. Não tenta mapear tudo nem ser objetivo. Foca em poucos lugares bem descritos, com opinião clara.

Cada lugar responde três coisas:
1. O que esse lugar é?
2. Para quem faz sentido?
3. Quando NÃO faz sentido?

Descrições são honestas, com limites explícitos. Se um lugar é barulhento, diz que é barulhento. Se é caro, diz que é caro. Sem marketing, sem adjetivos vazios.

## Perfis (lentes, não categorias)

Perfis são formas de ver a cidade, não caixas rígidas. Um lugar pode ter vários perfis, e você pode escolher vários também.

Perfis disponíveis:
- boemio
- baixa_gastro
- alta_gastro
- artes
- contemplativo
- noturno
- caotico_urbano

A ideia é que você escolhe seus perfis e o sistema sugere lugares que fazem sentido para você, explicando o porquê de cada sugestão.

## Evolução

Começou simples: lugares em memória, sem filtros. Depois veio proximidade e filtro por perfis. Em seguida, sugestões explicáveis (cada sugestão vem com reasons: "perfil: boemio", "perto: 1.2km").

O motor de sugestão é propositalmente simples no começo. Proximidade + match de perfis. Sem ML, sem score mágico. Tudo explicável.

Próximos passos: persistência, admin para manter os dados, depois melhorias no motor de sugestão conforme o uso mostrar o que falta.

## Como rodar

### Setup inicial

1. Suba o Postgres com docker-compose:
```bash
docker-compose up -d
```

O docker-compose usa valores padrão para as credenciais. Se quiser customizar, crie um arquivo `.env` na raiz do projeto (veja `.env.example` como referência).

2. Rode a API:
```bash
go run ./cmd/api
```

API em `http://localhost:8080`.

### Variáveis de ambiente

O docker-compose requer um arquivo `.env` com as seguintes variáveis:
- `POSTGRES_USER` — usuário do Postgres
- `POSTGRES_PASSWORD` — senha do Postgres
- `POSTGRES_DB` — nome do banco
- `POSTGRES_PORT` — porta do Postgres (default: 5432)

Crie um arquivo `.env` na raiz do projeto com essas variáveis. Veja `.env.example` como referência (se existir).

O código Go usa `os.Getenv()` para ler essas variáveis quando necessário (repo Postgres será implementado no Dia 11).

## Endpoints

### `GET /health`
Healthcheck.

### `GET /places?city=sp`
Lista lugares de uma cidade. Query params opcionais:
- `lat`, `lng` — localização do usuário
- `radius` — raio em km (default: 50)
- `profiles` — filtro por perfis (ex: `profiles=boemio,noturno`)

### `GET /places/{slug}`
Detalhes de um lugar pelo slug.

### `GET /suggestions?city=sp`
Sugestões personalizadas. Query params opcionais:
- `lat`, `lng` — localização do usuário
- `profiles` — perfis desejados (ex: `profiles=boemio`)

Retorna lugares ordenados por score, com reasons explicando cada sugestão:
```json
{
  "place": { ... },
  "score": 0.75,
  "reasons": ["perto: 1.2km", "perfil: boemio"]
}
```

