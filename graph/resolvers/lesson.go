package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
)

func (r *mutationResolver) CreateLesson(ctx context.Context, in models.CreateLessonInput) (*models.Lesson, error) {
	// get user
	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}

	//if !user.Read.LEARNING {
	//	return nil, ErrPermissionDenied
	//}

	chapter, err := r.LearningClient.CreateLesson(ctx, &learning.CreateLessonRequest{
		//UserID:      user.ID,
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
