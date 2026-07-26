package match

type MatchRepository interface {
	Save(match *Match) error
	SaveMany(matches []*Match) error
	FindByTeamId(id string) ([]*Match, error)
}