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
		IsTeacher:  in.IsTeacher,
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
		IsTeacher:  in.IsTeacher,
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{Succeeded: true}, nil
}

func (r *queryResolver) Comments(ctx context.Context, documentID string, replied bool) ([]*models.Comment, error) {

	user, err := permisions.User(ctx)
	if err != nil {
		return nil, err
	}

	var comments = new(learning.Comments)
	if replied {
		comments, err = r.LearningClient.GetRepliedComments(ctx, &learning.IDRequest{
			Id:     documentID,
			UserID: user.ID,
		})
		if err != nil {
			return nil, ErrApi(err)
		}

	} else {
		comments, err = r.LearningClient.GetComments(ctx, &learning.IDRequest{
			Id:     documentID,
			UserID: user.ID,
		})
	}

	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.Comments(comments), nil
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
