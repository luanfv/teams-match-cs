# Fluxo da aplicação

## Funcionalidades planejadas

1. Buscar times do banco de dados
2. Buscar confrontos de uma API externa
3. Cadastrar times novos no banco quando não existirem
4. Validar se os confrontos já estão no banco
5. Cadastrar confrontos novos no banco

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
