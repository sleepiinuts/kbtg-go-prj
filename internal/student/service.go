package student

import "github.com/sleepiinuts/kbtg-go-prj/internal/school"

type Service struct {
	schoolRepos school.Repos
}

func NewService(repos school.Repos) *Service {
	return &Service{schoolRepos: repos}
}

func (s *Service) CalculateGrade(score int) string {
	return "F"
}

func (s *Service) CalculateGradeByStudentName(name string) string {
	stu, err := s.schoolRepos.GetStudentByName(name)
	if err != nil {
		return ""
	}

	return studentCalculateGrade(stu.Score)
}

func studentCalculateGrade(score int) string {

	grade := ""
	switch {
	case score < 50:
		grade = "F"
	case score < 60:
		grade = "D"
	case score < 70:
		grade = "C"
	case score < 80:
		grade = "B"
	case score <= 100:
		grade = "A"
	}

	return grade
}
