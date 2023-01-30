package postgresql

import (
	"fmt"

	ptp "github.com/zakarynichols/parent-teacher-portal"
)

type schoolService struct {
	psql *psqlService
}

func NewSchoolService(psql *psqlService) *schoolService {
	return &schoolService{psql}
}

func (ss schoolService) CreateSchool(school ptp.School) error {
	_, err := ss.psql.db.Exec("INSERT INTO schools (name, location, type) VALUES ($1, $2, $3)", school.Name, school.Location, school.Type)
	if err != nil {
		return fmt.Errorf("failed to insert new school: %s", err.Error())
	}
	return nil
}

func (ss schoolService) GetSchool(id string) (ptp.School, error) {
	var school ptp.School
	err := ss.psql.db.QueryRow("SELECT * FROM schools WHERE id = $1", id).Scan(&school.ID, &school.Name, &school.Location, &school.Type)
	if err != nil {
		return school, fmt.Errorf("failed to retrieve school: %s", err.Error())
	}
	return school, nil
}

func (ss schoolService) GetAllSchools() ([]ptp.School, error) {
	rows, err := ss.psql.db.Query("SELECT * FROM schools")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve schools: %s", err.Error())
	}
	defer rows.Close()

	var schools []ptp.School
	for rows.Next() {
		var school ptp.School
		if err := rows.Scan(&school.ID, &school.Name, &school.Location, &school.Type); err != nil {
			return nil, fmt.Errorf("failed to retrieve school: %s", err.Error())
		}
		schools = append(schools, school)
	}
	return schools, nil
}

func (ss schoolService) UpdateSchool(id string, school ptp.School) error {
	_, err := ss.psql.db.Exec("UPDATE schools SET name = $1, location = $2, type = $3 WHERE id = $4", school.Name, school.Location, school.Type, id)
	if err != nil {
		return fmt.Errorf("failed to update school: %s", err.Error())
	}
	return nil
}

func (ss schoolService) DeleteSchool(id string) error {
	_, err := ss.psql.db.Exec("DELETE FROM schools WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete school: %s", err.Error())
	}
	return nil
}
