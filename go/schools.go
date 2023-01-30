package ptp

type School struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
}

type SchoolService interface {
	CreateSchool(school School) error
	GetSchool(id string) (School, error)
	GetAllSchools() ([]School, error)
	UpdateSchool(id string, school School) error
	DeleteSchool(id string) error
}
