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
	teamList, err := teamRepo.FindByIsFollowing(true); if err != nil {
		fmt.Println("Erro ao buscar times:", err)
		return
	}

	matchMap := make(map[string]*match.Match)
	newTeamMap := make(map[string]*team.Team)
	for _, t := range teamList {
		ml, err := matchRepo.FindByTeamId(t.Id())
		if err != nil {			
			fmt.Println("Erro ao buscar partidas:", err)
			return
		}
		for _, m := range ml {
			matchMap[m.Id()] = m
			if (t.Id() != m.TeamA().Id()) {
				newTeamMap[m.TeamA().Id()] = m.TeamA()
			}
			if (t.Id() != m.TeamB().Id()) {
				newTeamMap[m.TeamB().Id()] = m.TeamB()
			}
		}
	}

	newTeamList := make([]*team.Team, 0, len(newTeamMap))
	for _, t := range newTeamMap {
		newTeamList = append(newTeamList, t)
	}
	teamRepo.SaveMany(newTeamList)

	fmt.Println("Confrontos agendados:")
	for _, m := range matchMap {
		fmt.Printf("%s - %s vs %s: %s\n", m.TournamentName(), m.TeamA().Name(), m.TeamB().Name(), m.BeginAt().Format("02/01/2006 15:04"))
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
	t1, err := team.NewTeam(team.NewTeamInput{Name: "Fúria", ExternalId: 1, IsFollowing: true}); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return nil, nil
	}
	t2, err := team.NewTeam(team.NewTeamInput{Name: "Vitality", ExternalId: 2, IsFollowing: false}); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return nil, nil
	}
	m, err := match.NewMatch(match.NewMatchInput{TeamA: t1, TeamB: t2, BeginAt: time.Now(), ExternalId: 100, TournamentName: "Major 2026/2"}); if err != nil {
		fmt.Println("Erro ao criar confronto:", err)
		return nil, nil
	}

	teamList := []*team.Team{t1}
	teamRepo.SaveMany(teamList)
	matchList := []*match.Match{m}
	matchRepo.SaveMany(matchList)
	return teamRepo, matchRepo
}