package handler

func mockupComsciPLO() map[string]PLO {
	return map[string]PLO{
		randomID(): {
			ploName:        "PLO1",
			ploDescription: "balabla",
		},
		randomID(): {
			ploName:        "PLO2",
			ploDescription: "blabla",
		},
	}
}

func mockup102LO() map[string]LO {
	return map[string]LO{
		randomID(): {
			loTitle:      "LO1",
			linkedploIDs: map[string]bool{}, // don't assgin yet
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "LO1 level 1",
				},
				{
					level:            2,
					levelDescription: "LO1 level 2",
				},
			},
		},
		randomID(): {
			loTitle:      "LO2",
			linkedploIDs: map[string]bool{}, // don't assgin yet
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "LO2 level 1",
				},
				{
					level:            2,
					levelDescription: "LO2 level 2",
				},
			},
		},
	}
}

func mockup105LO() map[string]LO {
	return map[string]LO{
		randomID(): {
			loTitle:      "LO1",
			linkedploIDs: map[string]bool{},
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "LO1 level 1",
				},
			},
		},
	}
}

func mockupCourse() map[string]Course {
	return map[string]Course{
		randomID(): {
			courseName: "CSC102 programming",
			semester:   1,
			year:       2021,
			students:   map[string]Student{},
			quizzes:    map[string]Quiz{},
			los:        mockup102LO(),
		},
		randomID(): {
			courseName: "CSC105 web",
			semester:   2,
			year:       2021,
			students:   map[string]Student{},
			quizzes:    map[string]Quiz{},
			los:        mockup105LO(),
		},
	}
}

func mockupProgram() map[string]Program {
	return map[string]Program{
		randomID(): {
			programName:        "Computer Science",
			programDescription: "....",
			courses:            mockupCourse(),
			plos:               mockupComsciPLO(),
		},
	}
}
