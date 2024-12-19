package school

import (
	"fmt"

	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

type RedisRepos model.School

// AddStudentToRoom implements Repos.
func (rd *RedisRepos) AddStudentToRoom(name string, room int, score int) error {
	var r model.Room
	index := -1

	// check if room exist
	for i, d := range rd.Rooms {
		if d.No == room {
			index = i
			r = d
			break
		}
	}

	r.Students = append(r.Students, model.Student{Name: name, Score: score})

	// room not exist
	if index == -1 {
		r.No = room
		rd.Rooms = append(rd.Rooms, r)
	} else {
		rd.Rooms[index] = r
	}

	fmt.Println("Rooms: ", rd.Rooms)

	return nil
}

// GetStudentByName implements Repos.
func (rd *RedisRepos) GetStudentByName(name string) (*model.Student, error) {
	for _, d := range rd.Rooms {
		for _, s := range d.Students {
			if s.Name == name {
				return &s, nil
			}
		}
	}
	return &model.Student{}, nil
}

// GetStudentByRoom implements Repos.
func (rd *RedisRepos) GetStudentByRoom(room int) ([]model.Student, error) {
	for _, r := range rd.Rooms {
		if r.No == room {
			return r.Students, nil
		}
	}

	return nil, fmt.Errorf("room not found")
}

var _ Repos = &RedisRepos{}
