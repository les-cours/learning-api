package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
)

func (r *queryResolver) Documents(ctx context.Context, lessonID string) ([]*models.Document, error) {

	////get user
	//var user *types.UserToken
	//if user, _ = ctx.Value("user").(*types.UserToken); user == nil {
	//	return nil, ErrPermissionDenied
	//}
	//
	//if !user.Read.LEARNING {
	//	return nil, ErrPermissionDenied
	//}

	res, err := r.LearningClient.GetDocumentsByLesson(ctx, &learning.IDRequest{
		Id: lessonID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	lessons := gprcToGraph.Documents(res)

	return lessons, nil
}
