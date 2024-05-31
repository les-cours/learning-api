package gprcToGraph

import (
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	"log"
)

func ClassRoom(classRoom *learning.ClassRoom) *models.ClassRoom {
	if classRoom.Teacher == nil {
		classRoom.Teacher = &learning.Teacher{}
	}
	if classRoom.Chapters == nil {
		classRoom.Chapters = &learning.Chapters{}
	}
	return &models.ClassRoom{
		ClassRoomID:       classRoom.ClassRoomID,
		Title:             classRoom.Title,
		Image:             classRoom.Image,
		Price:             int(classRoom.Price),
		Badge:             classRoom.Badge,
		StudentCount:      int(classRoom.StudentCount),
		Rating:            float64(classRoom.Rating),
		ArabicTitle:       classRoom.ArabicTitle,
		Description:       classRoom.Description,
		ArabicDescription: classRoom.ArabicDescription,
		Teacher: &models.Teacher{
			TeacherID: classRoom.Teacher.TeacherID,
			Firstname: classRoom.Teacher.Firstname,
			Lastname:  classRoom.Teacher.Lastname,
		},
		Chapters: Chapters(classRoom.Chapters),
	}

}
func ClassRooms(cs *learning.ClassRooms) []*models.ClassRoom {
	var classRooms = make([]*models.ClassRoom, 0)

	for _, c := range cs.Classrooms {
		classRooms = append(classRooms, ClassRoom(c))
	}
	return classRooms
}

func Chapter(chapter *learning.Chapter) *models.Chapter {
	if chapter.Lessons == nil {
		chapter.Lessons = &learning.Lessons{}
	}
	return &models.Chapter{
		ChapterID:         chapter.ChapterID,
		Title:             chapter.Title,
		ArabicTitle:       chapter.ArabicTitle,
		Description:       chapter.Description,
		ArabicDescription: chapter.ArabicDescription,
		Lessons:           Lessons(chapter.Lessons),
	}
}
func Chapters(chapters *learning.Chapters) []*models.Chapter {
	log.Println("ch : ", chapters)
	var cc = make([]*models.Chapter, 0)

	for _, c := range chapters.Chapters {
		cc = append(cc, Chapter(c))
	}
	return cc
}

func Lesson(lesson *learning.Lesson) *models.Lesson {
	if lesson.Documents == nil {
		lesson.Documents = &learning.Documents{}
	}
	return &models.Lesson{
		LessonID:          lesson.LessonID,
		Title:             lesson.Title,
		ArabicTitle:       lesson.ArabicTitle,
		Description:       lesson.Description,
		ArabicDescription: lesson.ArabicDescription,
		Documents:         Documents(lesson.Documents),
	}

}
func Lessons(ls *learning.Lessons) []*models.Lesson {
	var lessons = make([]*models.Lesson, 0)

	for _, l := range ls.Lessons {
		lessons = append(lessons, Lesson(l))
	}
	return lessons
}

func StudentLesson(lesson *learning.StudentLesson) *models.StudentLesson {

	return &models.StudentLesson{
		Lesson: &models.Lesson{
			LessonID:    lesson.Lesson.LessonID,
			Title:       lesson.Lesson.Title,
			ArabicTitle: lesson.Lesson.ArabicTitle,
		},
		CanAccess: lesson.CanAccess,
	}

}
func StudentLessons(ls *learning.StudentLessons) []*models.StudentLesson {
	var lessons = make([]*models.StudentLesson, 0)

	for _, l := range ls.Lessons {
		lessons = append(lessons, StudentLesson(l))
	}
	return lessons
}

func Document(document *learning.Document) *models.Document {
	return &models.Document{
		DocumentID:   document.DocumentID,
		DocumentType: document.DocumentType,
		Title:        document.Title,
		ArabicTitle:  document.ArabicTitle,
		Description:  document.Description,
		Duration: &models.Duration{
			Hours:       int(document.Duration.Hours),
			Minutes:     int(document.Duration.Minutes),
			Seconds:     int(document.Duration.Seconds),
			Nanoseconds: int(document.Duration.Nanoseconds),
		},
		LectureNumber:     int(document.LectureNumber),
		DocumentLink:      document.DocumentLink,
		ArabicDescription: document.ArabicDescription,
	}

}
func Documents(ds *learning.Documents) []*models.Document {
	var documents = make([]*models.Document, 0)

	for _, d := range ds.Documents {
		documents = append(documents, Document(d))
	}
	return documents
}
