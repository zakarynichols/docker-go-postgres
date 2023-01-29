package ptp

// ptp -> parent-teacher-portal

import "github.com/zakarynichols/parent-teacher-portal/postgresql"

type SchoolService interface {
	CreateSchool(school postgresql.School) error
}
