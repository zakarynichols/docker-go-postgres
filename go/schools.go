package ptp

// ptp -> parent-teacher-portal

type SchoolService interface {
	CreateSchool(school School) error
}

type School struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Type     string `json:"type"`
}
