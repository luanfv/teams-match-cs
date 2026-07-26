package team

type TeamRepository interface {
	Save(team *Team) error
	SaveMany(teams []*Team) error
	FindById(id string) (*Team, error)
	FindAll() ([]*Team, error)
}