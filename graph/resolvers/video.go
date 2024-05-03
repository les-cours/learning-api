package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	"io/ioutil"
)

func (r *mutationResolver) UploadVideo(ctx context.Context, in models.UploadVideoInput) (*models.OperationStatus, error) {

	fileContent, _ := ioutil.ReadAll(in.Content.File)
	_, err := r.LearningClient.UploadVideo(ctx, &learning.UploadVideoRequest{
		LessonID: in.LessonID,
		Content:  fileContent,
		Filename: in.Content.Filename,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
