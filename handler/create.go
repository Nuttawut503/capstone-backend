package handler

import "github.com/google/uuid"

func (db *Database) createNewProgram(programName, programDescription string) {
	db.programs[uuid.New().String()] = Program{
		programName:        programName,
		programDescription: programDescription,
		courses:            map[string]Course{},
		plos:               map[string]PLO{},
	}
}

func (db *Database) createNewPLO(programID, ploName, ploDescription string) {
	db.programs[programID].plos[uuid.New().String()] = PLO{
		ploName:        ploName,
		ploDescription: ploDescription,
	}
}

func (db *Database) createNewCourse(programID, courseName string, semester, year int) {
	db.programs[programID].courses[uuid.New().String()] = Course{
		courseName: courseName,
		semester:   semester,
		year:       year,
		students:   map[string]Student{},
		los:        map[string]LO{},
		quizzes:    map[string]Quiz{},
	}
}

func (db *Database) addNewStudent(programID, courseID, studentID, studentEmail, studentName, studentSurname string) {
	db.programs[programID].courses[courseID].students[studentID] = Student{
		studentEmail:   studentEmail,
		studentName:    studentName,
		studentSurname: studentSurname,
	}
}

func (db *Database) addNewLO(programID, courseID, loTitle string, initLevel int, description string) {
	db.programs[programID].courses[courseID].los[uuid.New().String()] = LO{
		loTitle:      loTitle,
		levels:       []LOLevel{{level: initLevel, levelDescription: description}},
		linkedploIDs: map[string]bool{},
	}
}

func (db *Database) addNewLOLevel(programID, courseID, loID string, level int, description string) {
	cp := db.programs[programID].courses[courseID].los[loID]
	delete(db.programs[programID].courses[courseID].los, loID)
	db.programs[programID].courses[courseID].los[loID] = LO{
		loTitle:      cp.loTitle,
		levels:       append(cp.levels, LOLevel{level: level, levelDescription: description}),
		linkedploIDs: cp.linkedploIDs,
	}
}

func (db *Database) addPLOLink(programID, courseID, ploID, loID string) {
	db.programs[programID].courses[courseID].los[loID].linkedploIDs[ploID] = true
}

func (db *Database) addQuiz(programID, courseID, quizName string) {
	db.programs[programID].courses[courseID].quizzes[uuid.New().String()] = Quiz{
		quizName:  quizName,
		questions: map[string]Question{},
	}
}

func (db *Database) addNewQuestion(programID, courseID, quizID string, questionExcel []QuestionExcel) {
	questionIDs := map[string]string{}
	questionMaxScores := map[string]int{}
	questionMapResults := map[string][]QuestionResult{}
	for _, v := range questionExcel {
		id, added := questionIDs[v.QuestionTitle]
		if !added {
			id = uuid.New().String()
			questionIDs[v.QuestionTitle] = id
			questionMaxScores[id] = v.Maxscore
			questionMapResults[id] = []QuestionResult{}
		}
		questionMapResults[id] = append(questionMapResults[id], QuestionResult{
			studentID:    v.StudentID,
			studentScore: v.StudentScore,
		})
	}
	for question, id := range questionIDs {
		db.programs[programID].courses[courseID].quizzes[quizID].questions[id] = Question{
			questionTitle: question,
			maxScore:      questionMaxScores[id],
			results:       questionMapResults[id],
		}
	}
}

func (db *Database) addLOLink(programID, courseID, quizID, questionID, loID string, level int) {
	db.programs[programID].courses[courseID].quizzes[quizID].questions[questionID].linkedloIDs[loID] = level
}
