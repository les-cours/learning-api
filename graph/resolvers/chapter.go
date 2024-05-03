package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
)

func (r *mutationResolver) CreateChapter(ctx context.Context, in models.CreateChapterInput) (*models.Chapter, error) {

	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}

	chapter, err := r.LearningClient.CreateChapter(ctx, &learning.CreateChapterRequest{
		//	UserID:      user.AccountID,
		ClassRoomID: in.ClassRoomID,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		Description: in.Description,
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Chapter(chapter), nil
}
