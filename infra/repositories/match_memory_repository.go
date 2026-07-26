package repositories

import "github.com/luanfv/teams-match-cs/match"

type matchMemoryRepository struct {
    list []*match.Match
}

func NewMatchMemoryRepository() (match.MatchRepository, error) {
    return &matchMemoryRepository{}, nil
}

func (m *matchMemoryRepository) SaveMany(matches []*match.Match) error {
    m.list = append(m.list, matches...)
    return nil
}

func (m *matchMemoryRepository) Save(match *match.Match) error {
    m.list = append(m.list, match)
    return nil
}

func (m *matchMemoryRepository) FindByTeamId(teamId string) ([]*match.Match, error) {
    var results []*match.Match
    for _, match := range m.list {
        if match.TeamA().Id() == teamId || match.TeamB().Id() == teamId {
            results = append(results, match)
        }
    }
    return results, nil
}
