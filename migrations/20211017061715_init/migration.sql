-- CreateTable
CREATE TABLE "User" (
    "id" VARCHAR(11) NOT NULL,
    "email" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "surname" TEXT NOT NULL,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Student" (
    "id" VARCHAR(11) NOT NULL,

    CONSTRAINT "Student_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Teacher" (
    "id" VARCHAR(11) NOT NULL,

    CONSTRAINT "Teacher_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Program" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,

    CONSTRAINT "Program_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PLOgroup" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "programID" TEXT NOT NULL,

    CONSTRAINT "PLOgroup_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "PLO" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "ploGroupID" TEXT NOT NULL,

    CONSTRAINT "PLO_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Course" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "semester" INTEGER NOT NULL,
    "year" INTEGER NOT NULL,
    "programID" TEXT NOT NULL,
    "ploGroupID" TEXT NOT NULL,

    CONSTRAINT "Course_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "LO" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "courseID" TEXT NOT NULL,

    CONSTRAINT "LO_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "LOlevel" (
    "level" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "loID" TEXT NOT NULL,

    CONSTRAINT "LOlevel_pkey" PRIMARY KEY ("loID","level")
);

-- CreateTable
CREATE TABLE "LOlink" (
    "loID" TEXT NOT NULL,
    "ploID" TEXT NOT NULL,

    CONSTRAINT "LOlink_pkey" PRIMARY KEY ("loID","ploID")
);

-- CreateTable
CREATE TABLE "Quiz" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "courseID" TEXT NOT NULL,

    CONSTRAINT "Quiz_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Question" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "maxScore" INTEGER NOT NULL,
    "quizID" TEXT NOT NULL,

    CONSTRAINT "Question_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "QuestionResult" (
    "questionID" TEXT NOT NULL,
    "studentID" VARCHAR(11) NOT NULL,
    "score" INTEGER NOT NULL,

    CONSTRAINT "QuestionResult_pkey" PRIMARY KEY ("questionID","studentID")
);

-- CreateTable
CREATE TABLE "QuestionLink" (
    "questionID" TEXT NOT NULL,
    "loID" TEXT NOT NULL,
    "level" INTEGER NOT NULL,

    CONSTRAINT "QuestionLink_pkey" PRIMARY KEY ("questionID","loID","level")
);

-- CreateIndex
CREATE UNIQUE INDEX "Program_name_key" ON "Program"("name");

-- AddForeignKey
ALTER TABLE "Student" ADD CONSTRAINT "Student_id_fkey" FOREIGN KEY ("id") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Teacher" ADD CONSTRAINT "Teacher_id_fkey" FOREIGN KEY ("id") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PLOgroup" ADD CONSTRAINT "PLOgroup_programID_fkey" FOREIGN KEY ("programID") REFERENCES "Program"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "PLO" ADD CONSTRAINT "PLO_ploGroupID_fkey" FOREIGN KEY ("ploGroupID") REFERENCES "PLOgroup"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Course" ADD CONSTRAINT "Course_programID_fkey" FOREIGN KEY ("programID") REFERENCES "Program"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Course" ADD CONSTRAINT "Course_ploGroupID_fkey" FOREIGN KEY ("ploGroupID") REFERENCES "PLOgroup"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LO" ADD CONSTRAINT "LO_courseID_fkey" FOREIGN KEY ("courseID") REFERENCES "Course"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LOlevel" ADD CONSTRAINT "LOlevel_loID_fkey" FOREIGN KEY ("loID") REFERENCES "LO"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LOlink" ADD CONSTRAINT "LOlink_loID_fkey" FOREIGN KEY ("loID") REFERENCES "LO"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "LOlink" ADD CONSTRAINT "LOlink_ploID_fkey" FOREIGN KEY ("ploID") REFERENCES "PLO"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Quiz" ADD CONSTRAINT "Quiz_courseID_fkey" FOREIGN KEY ("courseID") REFERENCES "Course"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question" ADD CONSTRAINT "Question_quizID_fkey" FOREIGN KEY ("quizID") REFERENCES "Quiz"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "QuestionResult" ADD CONSTRAINT "QuestionResult_questionID_fkey" FOREIGN KEY ("questionID") REFERENCES "Question"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "QuestionResult" ADD CONSTRAINT "QuestionResult_studentID_fkey" FOREIGN KEY ("studentID") REFERENCES "Student"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "QuestionLink" ADD CONSTRAINT "QuestionLink_questionID_fkey" FOREIGN KEY ("questionID") REFERENCES "Question"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "QuestionLink" ADD CONSTRAINT "QuestionLink_loID_level_fkey" FOREIGN KEY ("loID", "level") REFERENCES "LOlevel"("loID", "level") ON DELETE CASCADE ON UPDATE CASCADE;
