// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"github.com/99designs/gqlgen/graphql"
)

type Chapter struct {
	ChapterID         string    `json:"chapterID"`
	Title             string    `json:"title"`
	ArabicTitle       string    `json:"arabicTitle"`
	Description       string    `json:"description"`
	ArabicDescription string    `json:"arabicDescription"`
	Lessons           []*Lesson `json:"lessons"`
}

type ClassRoom struct {
	ClassRoomID       string     `json:"classRoomID"`
	Title             string     `json:"title"`
	Image             string     `json:"image"`
	Price             int        `json:"price"`
	Badge             string     `json:"badge"`
	StudentCount      int        `json:"studentCount"`
	Rating            float64    `json:"rating"`
	ArabicTitle       string     `json:"arabicTitle"`
	Description       string     `json:"description"`
	ArabicDescription string     `json:"arabicDescription"`
	Teacher           *Teacher   `json:"teacher"`
	Chapters          []*Chapter `json:"chapters"`
}

type CreateChapterInput struct {
	ClassRoomID string `json:"classRoomID"`
	Title       string `json:"title"`
	ArabicTitle string `json:"arabicTitle"`
	Description string `json:"description"`
}

type CreateLessonInput struct {
	ChapterID   string `json:"chapterID"`
	Title       string `json:"title"`
	ArabicTitle string `json:"arabicTitle"`
	Description string `json:"description"`
}

type Document struct {
	DocumentID        string    `json:"documentID"`
	DocumentType      string    `json:"documentType"`
	Title             string    `json:"title"`
	ArabicTitle       string    `json:"arabicTitle"`
	Description       string    `json:"description"`
	ArabicDescription string    `json:"arabicDescription"`
	Duration          *Duration `json:"duration"`
	LectureNumber     int       `json:"lectureNumber"`
	DocumentLink      string    `json:"documentLink"`
}

type Duration struct {
	Hours       int `json:"hours"`
	Minutes     int `json:"minutes"`
	Seconds     int `json:"seconds"`
	Nanoseconds int `json:"nanoseconds"`
}

type IDRequest struct {
	ID string `json:"id"`
}

type Lesson struct {
	LessonID          string      `json:"lessonID"`
	Title             string      `json:"title"`
	ArabicTitle       string      `json:"arabicTitle"`
	Description       string      `json:"description"`
	ArabicDescription string      `json:"arabicDescription"`
	Documents         []*Document `json:"documents"`
}

type Mutation struct {
}

type OperationStatus struct {
	Succeeded bool `json:"succeeded"`
}

type Query struct {
}

type StudentLesson struct {
	Lesson    *Lesson `json:"lesson"`
	CanAccess bool    `json:"canAccess"`
}

type Teacher struct {
	TeacherID string `json:"teacherID"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UpdateChapterInput struct {
	ChapterID   string `json:"chapterID"`
	Title       string `json:"title"`
	ArabicTitle string `json:"arabicTitle"`
	Description string `json:"description"`
}

type UpdateLessonInput struct {
	LessonID    string `json:"lessonID"`
	ChapterID   string `json:"chapterID"`
	Title       string `json:"title"`
	ArabicTitle string `json:"arabicTitle"`
	Description string `json:"description"`
}

type UploadVideoInput struct {
	LessonID string         `json:"lessonID"`
	Content  graphql.Upload `json:"content"`
}
