package main

import (
	"fmt"
	"time"

	"github.com/luanfv/teams-match-cs/match"
	"github.com/luanfv/teams-match-cs/team"
)

func main() {
	t1, err := team.NewTeam("1", "Fúria"); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return
	}
	t2, err := team.NewTeam("2", "Vitality"); if err != nil {
		fmt.Println("Erro ao criar time:", err)
		return
	}
	m, err := match.NewMatch("1", t1, t2, time.Now()); if err != nil {
		fmt.Println("Erro ao criar confronto:", err)
		return
	}

	fmt.Printf("%s vs %s: %s\n", m.TeamA().Name(), m.TeamB().Name(), m.Date().Format("02/01/2006 15:04"))
}