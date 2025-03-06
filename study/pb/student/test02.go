package student

import (
	"fmt"
)

func Test02() {
	student := &Student{
		Name:   "sxy",
		Male:   false,
		Scores: []int32{71, 82, 93},
	}
	fmt.Println(student.GetName())
	fmt.Println(student.GetScores())
}
