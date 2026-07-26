package team

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	id string
	name string
	externalId int
	isFollowing bool
	createdAt time.Time
	updatedAt time.Time
}

type NewTeamInput struct {
	Name string
	ExternalId  int
	IsFollowing bool
}

type RestoreTeamInput struct {
	Id string
	Name string
	ExternalId int
	IsFollowing bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTeam(input NewTeamInput) (*Team, error) {
	return &Team{
		id: uuid.New().String(),
		name: input.Name,
		externalId: input.ExternalId,
		isFollowing: input.IsFollowing,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RestoreTeam(input RestoreTeamInput) (*Team, error) {
	return &Team{
		id: input.Id,
		name: input.Name,
		externalId: input.ExternalId,
		isFollowing: input.IsFollowing,
		createdAt: input.CreatedAt,
		updatedAt: input.UpdatedAt,
	}, nil
}

func (t *Team) Id() string {
	return t.id
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) ExternalId() int {
	return t.externalId
}

func (t *Team) IsFollowing() bool {
	return t.isFollowing
}

func (t *Team) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Team) UpdatedAt() time.Time {
	return t.updatedAt
}
