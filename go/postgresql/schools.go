package postgresql

import (
	"fmt"
)

type School struct {
	Name     string
	Location string
	Type     string
}

type schoolService struct {
	psql *psqlService
}

func NewSchoolService(psql *psqlService) *schoolService {
	return &schoolService{psql}
}

func (ss schoolService) CreateSchool(school School) error {
	_, err := ss.psql.db.Exec("INSERT INTO schools (name, location, type) VALUES ($1, $2, $3)", school.Name, school.Location, school.Type)
	if err != nil {
		return fmt.Errorf("failed to insert new school: %v", err)
	}
	return nil
}
