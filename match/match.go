package match

import (
	"errors"
	"time"

	"github.com/luanfv/teams-match-cs/team"
)

type Match struct {
	id string
	date time.Time
	teamA *team.Team
	teamB *team.Team
	createdAt time.Time
	updatedAt time.Time
}

func NewMatch(id string, teamA *team.Team, teamB *team.Team, date time.Time) (*Match, error) {
	if teamA.Id() == teamB.Id() {
		return nil, errors.New("Times não podem entrar em uma partida contra si mesmos")
	}

	// TODO - Aplica evento de domínio
	return &Match{
		id: id,
		date: date,
		teamA: teamA,
		teamB: teamB,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RestoreMatch(id string, teamA *team.Team, teamB *team.Team, date time.Time, createdAt time.Time, updatedAt time.Time) (*Match, error) {
	return &Match{
		id: id,
		date: date,
		teamA: teamA,
		teamB: teamB,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (m *Match) Id() string {
	return m.id
}

func (m *Match) Date() time.Time {
	return m.date
}

func (m *Match) TeamA() *team.Team {
	return m.teamA
}

func (m *Match) TeamB() *team.Team {
	return m.teamB
}

func (m *Match) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Match) UpdatedAt() time.Time {
	return m.updatedAt
}
