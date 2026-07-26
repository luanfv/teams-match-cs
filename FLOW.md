# Fluxo da aplicação

## Funcionalidades planejadas

- [x] Buscar times do banco de dados
- [ ] Buscar confrontos de uma API externa
- [x] Cadastrar times novos no banco quando não existirem
- [ ] Validar se os confrontos já estão no banco
- [ ] Cadastrar confrontos novos no banco

## Infraestrutura

- [ ] Substituir repositório em memória por banco de dados
- [ ] Implementar conexão com API externa (http client)
- [ ] Adicionar eventos de domínio
- [ ] Trigger via consumer (mensageria) ou http request

## Estrutura atual

```
/team
    team.go   → struct Team (id, name, createdAt, updatedAt)
               NewTeam: cria um novo time
               RestoreTeam: reconstrói um time vindo do banco

/match
    match.go  → struct Match (id, date, teamA, teamB, createdAt, updatedAt)
               NewMatch: cria um novo confronto (requer os dois times)
               RestoreMatch: reconstrói um confronto vindo do banco
```

## Dependência entre pacotes

```
match → conhece team
team  → não conhece match
```

## Ligação entre Match e Team

`Match` guarda os objetos `*Team` completos em memória.  
No banco, apenas os IDs serão persistidos — a reconstrução do objeto completo ficará no repository (futuramente).
