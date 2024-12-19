package school

import (
	"fmt"

	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

type Service struct {
	repos Repos
}

func NewService(repos Repos) *Service {
	return &Service{repos: repos}
}

type Repos interface {
	GetStudentByName(name string) (*model.Student, error)
	AddStudentToRoom(name string, room, score int) error
	GetStudentByRoom(room int) ([]model.Student, error)
}

func (s *Service) AddStudentToDB(room int, stu model.Student) error {

	// // append student to room
	// r := model.Room{No: room}
	// r.Students = append(r.Students, stu)

	// // append room to school
	// school.Rooms = append(school.Rooms, r)

	// required: student name
	if stu.Name == "" {
		return fmt.Errorf("student name is required")
	}

	// student is NOT allow to be in more than 1 room
	st, err := s.repos.GetStudentByName(stu.Name)
	if err != nil {
		return err
	}

	if st.Name != "" {
		return fmt.Errorf("student is already exist")
	}

	if err := s.repos.AddStudentToRoom(stu.Name, room, stu.Score); err != nil {
		return err
	}
	return nil
}

// GetStudentByRoom implements Schooler.
func (s *Service) GetStudentByRoom(room int) ([]model.Student, error) {
	return s.repos.GetStudentByRoom(room)
}

var _ Schooler = &Service{}
