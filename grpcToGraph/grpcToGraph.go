package gprcToGraph

import (
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
)

func ClassRoom(classRoom *learning.ClassRoom) *models.ClassRoom {
	return &models.ClassRoom{
		ClassRoomID:  classRoom.ClassRoomID,
		Title:        classRoom.Title,
		Image:        classRoom.Image,
		Price:        int(classRoom.Price),
		Badge:        classRoom.Badge,
		StudentCount: int(classRoom.StudentCount),
		Rating:       float64(classRoom.Rating),
	}

}

func ClassRooms(cs *learning.ClassRooms) []*models.ClassRoom {
	var classRooms []*models.ClassRoom

	for _, c := range cs.Classrooms {
		classRooms = append(classRooms, ClassRoom(c))
	}
	return classRooms
}

func Chapter(chapter *learning.Chapter) *models.Chapter {
	return &models.Chapter{
		ChapterID:   chapter.ChapterID,
		Title:       chapter.Title,
		ArabicTitle: chapter.ArabicTitle,
		Description: chapter.Description,
	}
}

func Lesson(lesson *learning.Lesson) *models.Lesson {
	return &models.Lesson{
		LessonID:    lesson.LessonID,
		Title:       lesson.Title,
		ArabicTitle: lesson.ArabicTitle,
	}

}

func Lessons(ls *learning.Lessons) []*models.Lesson {
	var lessons []*models.Lesson

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
	var lessons []*models.StudentLesson

	for _, l := range ls.Lessons {
		lessons = append(lessons, StudentLesson(l))
	}
	return lessons
}

func Document(document *learning.Document) *models.Document {
	return &models.Document{
		DocumentID:    document.DocumentID,
		DocumentType:  document.DocumentType,
		Title:         document.Title,
		ArabicTitle:   document.ArabicTitle,
		Description:   document.Description,
		Duration:      int(document.Duration),
		LectureNumber: int(document.LectureNumber),
		DocumentLink:  document.DocumentLink,
	}

}

func Documents(ds *learning.Documents) []*models.Document {
	var documents []*models.Document

	for _, d := range ds.Documents {
		documents = append(documents, Document(d))
	}
	return documents
}
