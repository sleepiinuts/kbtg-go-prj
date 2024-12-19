package school

import (
	"fmt"

	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

type MongoRepos model.School

// AddStudentToRoom implements Repos.
func (m *MongoRepos) AddStudentToRoom(name string, room int, score int) error {
	if len(m.Rooms) == 0 {
		m.Rooms = make([]model.Room, 1)
		m.Rooms[0].No = room
	}

	// add every student to single room
	m.Rooms[0].Students = append(m.Rooms[0].Students,
		model.Student{Name: name, Score: score})

	fmt.Println("Rooms: ", m.Rooms)
	return nil
}

// GetStudentByName implements Repos.
func (m *MongoRepos) GetStudentByName(name string) (*model.Student, error) {
	if len(m.Rooms) == 0 {
		return &model.Student{}, nil
	}

	for _, s := range m.Rooms[0].Students {
		if s.Name == name {
			return &s, nil
		}
	}
	return &model.Student{}, nil
}

// GetStudentByRoom implements Repos.
func (m *MongoRepos) GetStudentByRoom(room int) ([]model.Student, error) {
	for _, r := range m.Rooms {
		if r.No == room {
			return r.Students, nil
		}
	}

	return nil, fmt.Errorf("room not found")
}

var _ Repos = &MongoRepos{}
