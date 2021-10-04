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
					levelDescription: "LO1 Level1 Describe basic of data structures",
				},
				{
					level:            2,
					levelDescription: "LO1 Level2 Describe basic of the relevant algorithms",
				},
				{
					level:            3,
					levelDescription: "LO1 Level3 Discuss how data structures are related to its algorithms",
				},
			},
		},
		randomID(): {
			loTitle:      "LO2 Describe and discuss the use of built-in data structures",
			linkedploIDs: map[string]bool{}, // don't assgin yet
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "LO2 Level1 Describe the basic of built-in data structures",
				},
				{
					level:            2,
					levelDescription: "LO2 Level2 Discuss what is the best use of built-in data structures",
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
					levelDescription: "LO1 Level1 Able to tell how to solved the problem",
				},
			},
		},
		randomID(): {
			loTitle:      "LO2 Create algorithms for solving simple problems and use programming language to implement the algorithm of solution",
			linkedploIDs: map[string]bool{},
			levels: []LOLevel{
				{
					level:            1,
					levelDescription: "LO2 Level1 Create algorithm with Python",
				},
				{
					level:            2,
					levelDescription: "LO2 Level2 Create algorithm with Java",
				},
				{
					level:            3,
					levelDescription: "LO2 Level3 Able to improve create algorithms",
				},
			},
		},
	}
}

func mockupCourse() map[string]Course {
	return map[string]Course{
		randomID(): {
			courseName: "CSC209 Data Structures",
			courseDescription: "Abstract data type in Java, pointer and vector in Java, running time and complexity, linked-lists," +
				" stacks, queues, trees, recursion, numerical case studies, trees, graph, binary heap, tree algorithms, sorting case studies," +
				" hash table, data compression, string matching, event-driven programming.",
			semester: 2,
			year:     2021,
			students: map[string]Student{},
			quizzes:  map[string]Quiz{},
			los:      mockup209LO(),
		},
		randomID(): {
			courseName: "CSC102 Intro to Programming",
			courseDescription: "Fundamental concepts of programming, basic computation, simple I/O, standard conditional and" +
				" iterative structures, the definition of functions, and parameter passing, arrays, programming style and documentation, " +
				"program testing and debugging, basic algorithms and sorting, basic type systems, fundamental object-oriented programming.",
			semester: 1,
			year:     2021,
			students: map[string]Student{},
			quizzes:  map[string]Quiz{},
			los:      mockup102LO(),
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
