package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
	"github.com/les-cours/learning-api/types"
)

func (r *mutationResolver) CreateLesson(ctx context.Context, in models.CreateLessonInput) (*models.Lesson, error) {
	//get user
	var user *types.UserToken
	if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
		return nil, permisions.ErrPermissionDenied
	}

	if !user.Create.LEARNING {
		return nil, permisions.ErrPermissionDenied
	}

	chapter, err := r.LearningClient.CreateLesson(ctx, &learning.CreateLessonRequest{
		UserID:      user.ID,
		ChapterID:   in.ChapterID,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		Description: in.Description,
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Lesson(chapter), nil
}

func (r *queryResolver) Lessons(ctx context.Context, chapterID string) ([]*models.Lesson, error) {

	user, err := permisions.Teacher(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.LearningClient.GetLessonsByChapter(ctx, &learning.IDRequest{
		Id:     chapterID,
		UserID: user.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	lessons := gprcToGraph.Lessons(res)

	return lessons, nil
}

func (r *mutationResolver) UpdateLesson(ctx context.Context, in models.UpdateLessonInput) (*models.Lesson, error) {

	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}
	_, err := r.LearningClient.UpdateLesson(ctx, &learning.UpdateLessonRequest{
		LessonID: in.LessonID,
		//UserID:  user.Id   ,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		ChapterID:   in.ChapterID,
		Description: in.Description,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.Lesson{
		LessonID: in.LessonID,
		//UserID:  user.Id   ,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		Description: in.Description,
	}, nil
}

func (r *mutationResolver) DeleteLesson(ctx context.Context, in models.IDRequest) (*models.OperationStatus, error) {

	_, err := r.LearningClient.DeleteLesson(ctx, &learning.IDRequest{
		Id: in.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
