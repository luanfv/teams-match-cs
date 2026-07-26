package main

import (
	"fmt"
	"time"

	"github.com/luanfv/teams-match-cs/infra/repositories"
	"github.com/luanfv/teams-match-cs/match"
	"github.com/luanfv/teams-match-cs/team"
)

func main() {
	teamRepo, matchRepo := settings()
	if teamRepo == nil || matchRepo == nil {
		fmt.Println("Erro ao configurar repositórios")
		return
	}
	teamList, err := teamRepo.FindAll(); if err != nil {
		fmt.Println("Erro ao buscar times:", err)
		return
	}

	matchMap := make(map[string]*match.Match)
	for _, t := range teamList {
		ml, err := matchRepo.FindByTeamId(t.Id())
		if err != nil {
			fmt.Println("Erro ao buscar partidas:", err)
			return
		}
		for _, m := range ml {
			matchMap[m.Id()] = m
		}
	}

	fmt.Println("Confrontos agendados:")
	for _, m := range matchMap {
		fmt.Printf("%s vs %s: %s\n", m.TeamA().Name(), m.TeamB().Name(), m.Date().Format("02/01/2006 15:04"))
	}
}

func settings() (team.TeamRepository, match.MatchRepository) {
	teamRepo, err := repositories.NewTeamMemoryRepository(); if err != nil {
		fmt.Println("Erro ao criar repositório de times:", err)
		return nil, nil
	}
	matchRepo, err := repositories.NewMatchMemoryRepository(); if err != nil {
		fmt.Println("Erro ao criar repositório de partidas:", err)
		return nil, nil
	}
	t1, err := team.NewTeam("1", "Fúria"); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return nil, nil
	}
	t2, err := team.NewTeam("2", "Vitality"); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return nil, nil
	}
	m, err := match.NewMatch("1", t1, t2, time.Now()); if err != nil {
		fmt.Println("Erro ao criar confronto:", err)
		return nil, nil
	}

	teamList := []*team.Team{t1}
	teamRepo.SaveMany(teamList)
	matchList := []*match.Match{m}
	matchRepo.SaveMany(matchList)
	return teamRepo, matchRepo
}