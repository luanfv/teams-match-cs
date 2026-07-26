package match

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/luanfv/teams-match-cs/team"
)

type Match struct {
	id string
	externalId int
	beginAt time.Time
	tournamentName string
	teamA *team.Team
	teamB *team.Team
	createdAt time.Time
	updatedAt time.Time
}

type NewMatchInput struct {
	ExternalId int
	BeginAt time.Time
	TournamentName string
	TeamA *team.Team
	TeamB *team.Team
}

type RestoreMatchInput struct {
	Id string
	ExternalId int
	BeginAt time.Time
	TournamentName string
	TeamA *team.Team
	TeamB *team.Team
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewMatch(input NewMatchInput) (*Match, error) {
	if input.TeamA.Id() == input.TeamB.Id() {
		return nil, errors.New("Times não podem entrar em uma partida contra si mesmos")
	}

	// TODO - Aplica evento de domínio
	return &Match{
		id: uuid.New().String(),
		externalId: input.ExternalId,
		beginAt: input.BeginAt,
		tournamentName: input.TournamentName,
		teamA: input.TeamA,
		teamB: input.TeamB,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RestoreMatch(input RestoreMatchInput) (*Match, error) {
	return &Match{
		id: input.Id,
		externalId: input.ExternalId,
		beginAt: input.BeginAt,
		tournamentName: input.TournamentName,
		teamA: input.TeamA,
		teamB: input.TeamB,
		createdAt: input.CreatedAt,
		updatedAt: input.UpdatedAt,
	}, nil
}

func (m *Match) Id() string {
	return m.id
}

func (m *Match) ExternalId() int {
	return m.externalId
}

func (m *Match) BeginAt() time.Time {
	return m.beginAt
}

func (m *Match) TournamentName() string {
	return m.tournamentName
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
