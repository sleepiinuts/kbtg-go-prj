package student_test

import (
	"testing"

	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
	"github.com/sleepiinuts/kbtg-go-prj/internal/school"
	"github.com/sleepiinuts/kbtg-go-prj/internal/student"
)

func TestCalculateGradeXxx(t *testing.T) {
	s := student.NewService(nil)

	if s.CalculateGrade(49) != "F" {
		t.Errorf("invalid score, expected \"F\", but got %s", s.CalculateGrade(49))
	}
}

type mockSchoolRepos struct{}

// AddStudentToRoom implements school.Repos.
func (m *mockSchoolRepos) AddStudentToRoom(name string, room int, score int) error {
	panic("unimplemented")
}

// GetStudentByName implements school.Repos.
func (m *mockSchoolRepos) GetStudentByName(name string) (*model.Student, error) {
	var student *model.Student
	switch name {
	case "John":
		student = &model.Student{Name: "John", Score: 80}
	default:
		student = &model.Student{Name: "default", Score: 0}

	}
	return student, nil
}

// GetStudentByRoom implements school.Repos.
func (m *mockSchoolRepos) GetStudentByRoom(room int) ([]model.Student, error) {
	panic("unimplemented")
}

var _ school.Repos = &mockSchoolRepos{}

func TestCalculateGradeByStudentName(t *testing.T) {
	m := mockSchoolRepos{}
	s := student.NewService(&m)

	g1 := s.CalculateGradeByStudentName("John")
	if g1 != "A" {
		t.Errorf("invalid score, expected \"A\", but got %s", g1)
	}

	g2 := s.CalculateGradeByStudentName("Susan")
	if g2 != "F" {
		t.Errorf("invalid score, expected \"F\", but got %s", g2)
	}
}
