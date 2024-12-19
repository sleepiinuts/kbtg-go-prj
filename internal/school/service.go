package school

import (
	"github.com/sleepiinuts/kbtg-go-prj/internal/model"
)

func addStudent(room int, stu model.Student) error {

	// append student to room
	r := model.Room{No: room}
	r.Students = append(r.Students, stu)

	// append room to school
	school.Rooms = append(school.Rooms, r)
	return nil
}
