package repositories

import "github.com/luanfv/teams-match-cs/team"

type teamMemoryRepository struct {
	list []*team.Team
}

func NewTeamMemoryRepository() (team.TeamRepository, error) {
	return &teamMemoryRepository{}, nil
}

func (t *teamMemoryRepository) Save(team *team.Team) error {
	t.list = append(t.list, team)
	return nil
}

func (t *teamMemoryRepository) SaveMany(teams []*team.Team) error {
	t.list = append(t.list, teams...)
	return nil
}

func (t *teamMemoryRepository) FindById(id string) (*team.Team, error) {
	for _, team := range t.list {
		if team.Id() == id {
			return team, nil
		}
	}
	return nil, nil
}

func (t *teamMemoryRepository) FindByIsFollowing(isFollowing bool) ([]*team.Team, error) {
	var result []*team.Team
	for _, team := range t.list {
		if team.IsFollowing() == isFollowing {
			result = append(result, team)
		}
	}
	return result, nil
}
