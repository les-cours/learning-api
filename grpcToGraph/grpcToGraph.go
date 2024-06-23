package gprcToGraph

import (
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
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
		Order:             int(lesson.LessonOrder),
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

func Document(document *learning.Document) *models.Document {
	var duration = new(models.Duration)
	if document.Duration != nil {
		duration = &models.Duration{
			Hours:       int(document.Duration.Hours),
			Minutes:     int(document.Duration.Minutes),
			Seconds:     int(document.Duration.Seconds),
			Nanoseconds: int(document.Duration.Nanoseconds),
		}

	}
	return &models.Document{
		DocumentID:        document.DocumentID,
		DocumentType:      document.DocumentType,
		Title:             document.Title,
		ArabicTitle:       document.ArabicTitle,
		Description:       document.Description,
		LectureNumber:     int(document.LectureNumber),
		DocumentLink:      document.DocumentLink,
		Duration:          duration,
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

func Comment(c *learning.Comment) *models.Comment {
	return &models.Comment{
		ID: c.Id,
		User: &models.User{
			ID:        c.User.Id,
			Username:  c.User.Username,
			FirstName: c.User.FirstName,
			LastName:  c.User.LastName,
			Avatar:    c.User.Avatar,
		},
		RepliedTo:  c.RepliedTo,
		Content:    c.Content,
		DocumentID: c.DocumentID,
		Timestamp:  int(c.Timestamp),
		Edited:     c.Edited,
		IsTeacher:  c.IsTeacher,
	}

}
func Comments(ds *learning.Comments) []*models.Comment {
	var documents = make([]*models.Comment, 0)

	for _, d := range ds.Comments {
		documents = append(documents, Comment(d))
	}
	return documents
}

func Room(room *learning.Room) *models.Room {
	return &models.Room{
		ID:       room.Id,
		Name:     room.Name,
		Teacher:  User(room.Owner),
		Users:    Users(room.Users),
		Messages: Messages(room.Messages),
	}
}

func Messages(messages []*learning.Message) []*models.Message {
	var grpcMessage = make([]*models.Message, 0)
	for _, message := range messages {
		grpcMessage = append(grpcMessage, Message(message))
	}
	return grpcMessage
}
func Message(message *learning.Message) *models.Message {
	return &models.Message{
		ID:        message.Id,
		RoomID:    message.RoomId,
		Message:   message.Content,
		Timestamp: message.Timestamp,
		IsTeacher: message.IsTeacher,
		Owner:     User(message.User),
	}
}

func Users(users []*learning.User) []*models.UserRoom {
	var grpcUsers = make([]*models.UserRoom, 0)
	for _, user := range users {
		grpcUsers = append(grpcUsers, User(user))
	}
	return grpcUsers
}
func User(user *learning.User) *models.UserRoom {
	return &models.UserRoom{
		ID:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Paid:      user.Paid,
	}
}

func Subscriptions(subscriptions *learning.Subscriptions) []*models.Subscription {
	var grpcSubs = make([]*models.Subscription, 0)
	for _, subscription := range subscriptions.Subscriptions {
		grpcSubs = append(grpcSubs, Subscription(subscription))
	}
	return grpcSubs
}
func Subscription(subscription *learning.Subscription) *models.Subscription {
	return &models.Subscription{
		ID:      subscription.Id,
		MonthID: int(subscription.MonthId),
		PaidAt:  int(subscription.PaidAt),
	}
}
