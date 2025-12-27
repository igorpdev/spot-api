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

```bash
go run ./cmd/api
```

API em `http://localhost:8080`.

Endpoints básicos:
- `GET /health`
- `GET /places`
- `GET /places/{slug}`

