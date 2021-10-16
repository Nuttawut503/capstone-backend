-- CreateEnum
CREATE TYPE "Role" AS ENUM ('STUDENT', 'TEACHER', 'LEADER');

-- CreateTable
CREATE TABLE "User" (
    "id" VARCHAR(11) NOT NULL,
    "email" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "surname" TEXT NOT NULL,
    "role" "Role" NOT NULL DEFAULT E'STUDENT',

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Program" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,

    CONSTRAINT "Program_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Plo_plan" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "programID" TEXT NOT NULL,

    CONSTRAINT "Plo_plan_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Plo_info" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "planID" TEXT NOT NULL,

    CONSTRAINT "Plo_info_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Course" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "semester" INTEGER NOT NULL,
    "year" INTEGER NOT NULL,
    "programID" TEXT NOT NULL,
    "planID" TEXT NOT NULL,

    CONSTRAINT "Course_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Lo" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "courseID" TEXT NOT NULL,

    CONSTRAINT "Lo_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Lo_level" (
    "level" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "loID" TEXT NOT NULL,

    CONSTRAINT "Lo_level_pkey" PRIMARY KEY ("loID","level")
);

-- CreateTable
CREATE TABLE "Lo_link" (
    "loID" TEXT NOT NULL,
    "ploID" TEXT NOT NULL,

    CONSTRAINT "Lo_link_pkey" PRIMARY KEY ("loID","ploID")
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
CREATE TABLE "Question_result" (
    "questionID" TEXT NOT NULL,
    "studentID" VARCHAR(11) NOT NULL,
    "score" INTEGER NOT NULL,

    CONSTRAINT "Question_result_pkey" PRIMARY KEY ("questionID","studentID")
);

-- CreateTable
CREATE TABLE "Question_link" (
    "questionID" TEXT NOT NULL,
    "loID" TEXT NOT NULL,
    "level" INTEGER NOT NULL,

    CONSTRAINT "Question_link_pkey" PRIMARY KEY ("questionID","loID","level")
);

-- CreateIndex
CREATE UNIQUE INDEX "Program_name_key" ON "Program"("name");

-- AddForeignKey
ALTER TABLE "Plo_plan" ADD CONSTRAINT "Plo_plan_programID_fkey" FOREIGN KEY ("programID") REFERENCES "Program"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Plo_info" ADD CONSTRAINT "Plo_info_planID_fkey" FOREIGN KEY ("planID") REFERENCES "Plo_plan"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Course" ADD CONSTRAINT "Course_programID_fkey" FOREIGN KEY ("programID") REFERENCES "Program"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Course" ADD CONSTRAINT "Course_planID_fkey" FOREIGN KEY ("planID") REFERENCES "Plo_plan"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Lo" ADD CONSTRAINT "Lo_courseID_fkey" FOREIGN KEY ("courseID") REFERENCES "Course"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Lo_level" ADD CONSTRAINT "Lo_level_loID_fkey" FOREIGN KEY ("loID") REFERENCES "Lo"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Lo_link" ADD CONSTRAINT "Lo_link_loID_fkey" FOREIGN KEY ("loID") REFERENCES "Lo"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Lo_link" ADD CONSTRAINT "Lo_link_ploID_fkey" FOREIGN KEY ("ploID") REFERENCES "Plo_info"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Quiz" ADD CONSTRAINT "Quiz_courseID_fkey" FOREIGN KEY ("courseID") REFERENCES "Course"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question" ADD CONSTRAINT "Question_quizID_fkey" FOREIGN KEY ("quizID") REFERENCES "Quiz"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question_result" ADD CONSTRAINT "Question_result_questionID_fkey" FOREIGN KEY ("questionID") REFERENCES "Question"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question_result" ADD CONSTRAINT "Question_result_studentID_fkey" FOREIGN KEY ("studentID") REFERENCES "User"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question_link" ADD CONSTRAINT "Question_link_questionID_fkey" FOREIGN KEY ("questionID") REFERENCES "Question"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Question_link" ADD CONSTRAINT "Question_link_loID_level_fkey" FOREIGN KEY ("loID", "level") REFERENCES "Lo_level"("loID", "level") ON DELETE CASCADE ON UPDATE CASCADE;
