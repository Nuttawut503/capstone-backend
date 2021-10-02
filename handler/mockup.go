package handler

func mockupComsciPLO() map[string]PLO {
	return map[string]PLO{
		"0000011": {
			ploName:        "PLO1",
			ploDescription: "An ability to apply knowledge of computer sciences appropriate to the discipline.",
		},
		randomID(): {
			ploName: "PLO2",
			ploDescription: "An ability to use appropriate system design notations and apply" +
				"software engineering process in order to plan, design, and implement software systems of varying complexity.",
		},
		randomID(): {
			ploName: "PLO3",
			ploDescription: "An ability to demonstrate a depth of knowledge appropriate to graduate study and/or" +
				"lifelong learning in a self-selected area in computing.",
		},
	}
}

func mockup209LO() map[string]LO {
	return map[string]LO{
		randomID(): {
			loTitle: "LO1 Describe and discuss fundamental data structures and the relevant algorithms",
			linkedploIDs: map[string]bool{
				"0000011": true,
			}, // don't assgin yet
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "Describe basic of data structures",
				},
				{
					level:            2,
					levelDescription: "Describe basic of the relevant algorithms",
				},
				{
					level:            3,
					levelDescription: "Discuss how data structures are related to its algorithms",
				},
			},
		},
		randomID(): {
			loTitle:      "LO2 Describe and discuss the use of built-in data structures",
			linkedploIDs: map[string]bool{}, // don't assgin yet
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "Describe the basic of built-in data structures",
				},
				{
					level:            2,
					levelDescription: "Discuss what is the best use of built-in data structures",
				},
			},
		},
	}
}

func mockup102LO() map[string]LO {
	return map[string]LO{
		randomID(): {
			loTitle:      "LO1 Discuss how a problem may be solved by multiple algorithms",
			linkedploIDs: map[string]bool{},
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "Able to tell how to solved the problem",
				},
			},
		},
		randomID(): {
			loTitle:      "LO2 Create algorithms for solving simple problems and use programming language to implement the algorithm of solution",
			linkedploIDs: map[string]bool{},
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "Create algorithm with Python",
				},
				{
					level:            2,
					levelDescription: "Create algorithm with Java",
				},
				{
					level:            3,
					levelDescription: "Able to improve create algorithms",
				},
			},
		},
	}
}

func mockupCourse() map[string]Course {
	return map[string]Course{
		randomID(): {
			courseName: "CSC209 Data Structures",
			semester:   2,
			year:       2021,
			students:   map[string]Student{},
			quizzes:    map[string]Quiz{},
			los:        mockup209LO(),
		},
		randomID(): {
			courseName: "CSC102 Intro to Programming",
			semester:   1,
			year:       2021,
			students:   map[string]Student{},
			quizzes:    map[string]Quiz{},
			los:        mockup102LO(),
		},
	}
}

func mockupProgram() map[string]Program {
	return map[string]Program{
		randomID(): {
			programName:        "Computer Science",
			programDescription: "Computer science is the study of algorithmic processes, computational machines and computation itself.",
			courses:            mockupCourse(),
			plos:               mockupComsciPLO(),
		},
		randomID(): {
			programName: "Information Technology",
			programDescription: "Information technology is the use of any computers, storage, networking and other physical devices," +
				"infrastructure and processes to create, process, store, secure and exchange all forms of electronic data.",
			courses: mockupCourse(),
			plos:    mockupComsciPLO(),
		},
	}
}
