package student

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
	default:
		grade = "A"
	}

	return grade
}
