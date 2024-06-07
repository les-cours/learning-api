package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *queryResolver) Documents(ctx context.Context, lessonID string) ([]*models.Document, error) {

	_, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetDocuments(ctx, &learning.IDRequest{
		Id: lessonID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	lessons := gprcToGraph.Documents(res)

	return lessons, nil
}

func (r *queryResolver) Document(ctx context.Context, documentID string) (*models.DocumentLink, error) {

	_, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetDocument(ctx, &learning.IDRequest{
		Id: documentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.DocumentLink{
		DocumentLink: res.DocumentLink,
	}, nil
}

func (r *mutationResolver) DeleteDocument(ctx context.Context, documentID string) (*models.OperationStatus, error) {

	_, err := permisions.CanDelete(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.LearningClient.DeleteDocument(ctx, &learning.IDRequest{
		Id: documentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
