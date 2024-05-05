package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *mutationResolver) CreateChapter(ctx context.Context, in models.CreateChapterInput) (*models.Chapter, error) {

	_, err := permisions.CanCreate(ctx)
	if err != nil {
		return nil, ErrApi(err)
	}

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

func (r *queryResolver) Chapters(ctx context.Context, classRoomID string) ([]*models.Chapter, error) {

	var chapters []*models.Chapter

	res, err := r.LearningClient.GetChaptersByClassRoom(ctx, &learning.IDRequest{
		Id: classRoomID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	for _, chapter := range res.Chapters {
		chapters = append(chapters, gprcToGraph.Chapter(chapter))
	}

	return chapters, nil
}

func (r *mutationResolver) UpdateChapter(ctx context.Context, in models.UpdateChapterInput) (*models.Chapter, error) {

	_, err := permisions.CanUpdate(ctx)
	if err != nil {
		return nil, ErrApi(err)
	}

	_, err = r.LearningClient.UpdateChapter(ctx, &learning.UpdateChapterRequest{
		ChapterID: in.ChapterID,
		//UserID:  user.Id   ,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		Description: in.Description,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.Chapter{
		ChapterID:   in.ChapterID,
		Title:       in.Title,
		ArabicTitle: in.ArabicTitle,
		Description: in.Description,
	}, nil
}

func (r *mutationResolver) DeleteChapter(ctx context.Context, in models.IDRequest) (*models.OperationStatus, error) {

	user, err := permisions.CanDelete(ctx)
	if err != nil {
		return nil, ErrApi(err)
	}

	_, err = r.LearningClient.DeleteChapter(ctx, &learning.IDRequest{
		Id:     in.ID,
		UserID: user.ID,
	})

	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}

/*
STUDENTS
*/
