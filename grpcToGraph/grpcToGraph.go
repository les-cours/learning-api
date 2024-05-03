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
