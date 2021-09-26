package database

import (
	"backend/database/course"
	"backend/database/lo"
	"backend/database/lolevel"
	"backend/database/plo"
	"backend/database/plolink"
	"backend/database/ploversion"
	"backend/database/program"
	"backend/database/question"
	"backend/database/questionlink"
	"backend/database/questionresult"
	"backend/database/quiz"
	"backend/database/student"
	"backend/database/user"
)

func Connect() {
	course.Connect()
	lo.Connect()
	lolevel.Connect()
	plo.Connect()
	plolink.Connect()
	ploversion.Connect()
	program.Connect()
	question.Connect()
	questionlink.Connect()
	questionresult.Connect()
	quiz.Connect()
	student.Connect()
	user.Connect()
}
