package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *mutationResolver) CreatePDF(ctx context.Context, in models.CreatePDFInput) (*models.Document, error) {

	user, err2 := permisions.CanCreate(ctx)
	if err2 != nil {
		return nil, err2
	}
	res, err := r.LearningClient.CreatePdf(ctx, &learning.CreatePdfRequest{
		UserID:            user.ID,
		LessonID:          in.LessonID,
		Title:             in.Title,
		ArabicTitle:       in.ArabicTitle,
		Description:       in.Description,
		ArabicDescription: in.ArabicDescription,
		LectureNumber:     int32(in.LectureNumber),
		Url:               in.URL,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.Document(res), nil
}
