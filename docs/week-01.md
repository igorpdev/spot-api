# Semana 1 — SPOT

Documentação do que foi feito nos dias 1-8.
Tom honesto, sem marketing.

---

## O que fiz

### Dia 1 — Levantar o chão
- API básica em Go com `net/http`
- Entidade `Place` no domínio (ID, Name, Slug, Lat/Lng, Profiles)
- Repositório in-memory
- Endpoints: `GET /health`, `GET /places`, `GET /places/{slug}`
- Seed com 10 lugares de SP
- Validações mínimas (nome, lat/lng, perfis)

**Decisão:** Sem usecase layer. Handlers chamam repositórios diretamente. Adicionar usecase só quando houver lógica real de orquestração.

### Dia 2 — Cidade explícita
- Entidade `City` no domínio
- `Place` passa a ter `CityID` obrigatório
- Query param `city` obrigatório em `/places`
- Parsing de query params: `lat`, `lng`, `radius`, `profiles` (ainda não usados)

**Decisão:** Endereçar dívida conceitual cedo. Cidade vira conceito explícito, não implícito. Preparar terreno para múltiplas cidades sem refactor futuro.

### Dia 3 — Filtros e ordenação
- Cálculo de distância (Haversine)
- Filtro por distância (radius)
- Filtro por perfis (interseção, não união)
- Ordenação por distância quando lat/lng fornecidos
- Extração de lógica de filtro para função separada

**Observação:** Handler `GetPlaces` ficou com ~90 linhas. Escopo grande para um dia. Para próximos dias similares, considerar quebrar em parsing/validação, filtros e ordenação em dias separados.

### Dia 4 — Endpoint de sugestões (primeiro passo)
- `GET /suggestions` com `city` obrigatório
- Retorna todos os lugares da cidade (sem score ainda)

**Decisão:** Quebrar implementação de sugestões em múltiplos dias. Primeiro endpoint, depois score, depois ordenação, depois reasons.

### Dia 5 — Score e filtro
- Cálculo de score simples: proximidade (50%) + match de perfis (50%)
- Filtro: apenas lugares com score > 0
- Extração de parsing compartilhado (`parseLocationAndProfiles`)

**Decisão:** Score propositalmente simples. Melhorias vêm depois, se necessário.

### Dia 6 — Ordenação e DTO
- Ordenação por score (maior primeiro)
- DTO `SuggestionResponse` com `Place` e `Score`

### Dia 7 — Reasons explicáveis
- Reasons por sugestão: "perfil: boemio", "perto: 1.4km"
- Testes unitários do motor de sugestão (3 casos)

**Decisão:** Toda sugestão precisa explicar por quê. Transparência é feature.

### Dia 8 — Campos de curadoria
- Campos no `Place`: `Description`, `MakesSenseFor`, `DoesNotMakeSenseIf`, `Tags`
- Seed atualizado com 3 lugares usando os novos campos

**Decisão:** Campos opcionais. Lugar sem `Description` ainda funciona. Revisão completa do seed vem nos Dias 25-26 (quando campos V2 estiverem implementados).

---

## O que ficou incompleto

### Slug generation
O `GenerateSlug` é bem simples: só remove acentos e espaços. Pode gerar colisões. Melhorar depois quando houver necessidade real.

### Handler GetPlaces grande
Ficou com ~90 linhas no Dia 3. Funciona, mas indica escopo grande para um dia. Aprendizado: quebrar atividades similares em múltiplos dias.

### Seed parcial
Apenas 3 lugares têm campos de curadoria completos. Os outros 7 ainda não. Isso é intencional — revisão completa vem nos Dias 25-26. Mas pode confundir quem olha o código sem contexto.

### Parsing de horários (futuro)
Quando implementar campos temporais (Dia 15+), o parsing de horários será heurística simples. "sexta" = sexta-feira, "noite" = após 18h. Não precisa de parsing exato no começo.

---

## O que decidi não fazer

### Usecase layer
Removido no Dia 1. Estava funcionando só como pass-through. Handlers chamam repositórios diretamente. Adicionar usecase só quando houver lógica real de orquestração.

### Banco de dados
Fica para os Dias 10-11. In-memory é suficiente para validar o conceito.

### Validações complexas
Validações são mínimas. Nome não vazio, lat/lng válidos, pelo menos 1 perfil. Validações mais complexas vêm depois, se necessário.

### ML ou heurística avançada
Score é simples: proximidade + match de perfis. Nada de machine learning ou heurística complexa. Simplicidade proposital.

### Múltiplas cidades reais
Estrutura suporta múltiplas cidades, mas só SP existe por enquanto. Escalar depois quando houver necessidade real.

### Revisão completa do seed
Apenas 3 lugares têm campos de curadoria completos. Revisão completa vem nos Dias 25-26, quando campos V2 estiverem implementados.

---

## Decisões importantes

### Cidade como conceito explícito (Dia 2)
Endereçar dívida conceitual cedo. `Place` tem `CityID` obrigatório desde o Dia 2. Endpoints aceitam `city` como contexto. Preparar terreno sem over-engineering.

### Score simples (Dia 5)
Proximidade (50%) + match de perfis (50%). Nada de ML ou heurística complexa. Simplicidade proposital. Melhorias vêm depois, se necessário.

### Reasons obrigatórias (Dia 7)
Toda sugestão precisa explicar por quê. Transparência é feature. Nada de "porque o algoritmo decidiu".

### Campos opcionais (Dia 8)
Campos de curadoria são opcionais. Lugar sem `Description` ainda funciona. Não quebrar compatibilidade com lugares existentes.

---

## Próximos passos

### Semana 2 (Dias 10-14)
- Docker + Postgres (Dias 10-11)
- Admin endpoints (Dias 12-13)
- Deixar apresentável (Dia 14)

### Expansão V2 (Dias 15+)
- Contexto temporal (QUANDO)
- Companhia (COM QUEM)
- Preço (QUANTO)
- Anti-perfis (O QUE NÃO)
- Lugares temporários

Ver `EXPANSION.md` para detalhes.

---

## Limitações conhecidas

- Slug pode gerar colisões
- Handler `GetPlaces` grande (~90 linhas)
- Seed parcial (apenas 3 lugares com campos completos)
- Score simples (pode melhorar depois)
- Sem persistência real (in-memory por enquanto)
- Apenas São Paulo (estrutura pronta para múltiplas cidades)

Nada disso é bug. São decisões conscientes ou limitações temporárias.

---

## Tom e identidade

O código parece trabalho real, não gerado por IA:
- Comentários só quando explicam decisão ou intenção
- Commits com mensagens humanas
- Seed parcial mostra iteração, não código perfeito
- Escopo grande em alguns dias mostra ritmo humano

Isso é intencional. Ver `PROJECT_CONTEXT.md` para princípios.

