package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *mutationResolver) CreateComment(ctx context.Context, in models.CreateCommentInput) (*models.OperationStatus, error) {

	user, err := permisions.User(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.LearningClient.CreateComment(ctx, &learning.CreateCommentRequest{
		UserID:     user.ID,
		Content:    in.Content,
		DocumentID: in.DocumentID,
		IsTeacher:  false,
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{Succeeded: true}, nil
}
func (r *mutationResolver) CreateReply(ctx context.Context, in models.CreateReplyInput) (*models.OperationStatus, error) {

	user, err := permisions.User(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.LearningClient.CreateComment(ctx, &learning.CreateCommentRequest{
		UserID:     user.ID,
		Content:    in.Content,
		DocumentID: in.DocumentID,
		RepliedTo:  in.RepliedTo,
		IsTeacher:  false,
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{Succeeded: true}, nil
}

func (r *queryResolver) Comments(ctx context.Context, documentID string) ([]*models.Comment, error) {

	res, err := r.LearningClient.GetComments(ctx, &learning.IDRequest{
		Id:     documentID,
		UserID: "2b836dd5-760c-47b6-9f0a-ddf39293afca",
	})
	if err != nil {
		return nil, err
	}

	return gprcToGraph.Comments(res), nil
}

//
//func (r *mutationResolver) UpdateChapter(ctx context.Context, in models.UpdateChapterInput) (*models.Chapter, error) {
//
//	user, err := permisions.Teacher(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = r.LearningClient.UpdateChapter(ctx, &learning.UpdateChapterRequest{
//		ChapterID:   in.ChapterID,
//		UserID:      user.ID,
//		Title:       in.Title,
//		ArabicTitle: in.ArabicTitle,
//		Description: in.Description,
//	})
//	if err != nil {
//		return nil, ErrApi(err)
//	}
//
//	return &models.Chapter{
//		ChapterID:   in.ChapterID,
//		Title:       in.Title,
//		ArabicTitle: in.ArabicTitle,
//		Description: in.Description,
//	}, nil
//}
//
//func (r *mutationResolver) DeleteChapter(ctx context.Context, in models.IDRequest) (*models.OperationStatus, error) {
//
//	user, err := permisions.Teacher(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = r.LearningClient.DeleteChapter(ctx, &learning.IDRequest{
//		Id:     in.ID,
//		UserID: user.ID,
//	})
//
//	if err != nil {
//		return nil, ErrApi(err)
//	}
//
//	return &models.OperationStatus{
//		Succeeded: true,
//	}, nil
//}

/*
STUDENTS
*/
