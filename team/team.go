package team

import "time"

type Team struct {
	id string
	name string
	createdAt time.Time
	updatedAt time.Time
}

func NewTeam(id string, name string) (*Team, error) {
	return &Team{
		id: id,
		name: name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func RestoreTeam(id string, name string, createdAt time.Time, updatedAt time.Time) (*Team, error) {
	return &Team{
		id: id,
		name: name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (t *Team) Id() string {
	return t.id
}

func (t *Team) Name() string {
	return t.name
}

func (t *Team) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Team) UpdatedAt() time.Time {
	return t.updatedAt
}
