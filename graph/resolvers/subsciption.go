package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *queryResolver) MySubscription(ctx context.Context, classroomID string) (*models.CurrentSubscription, error) {

	student, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetCurrentSubscription(ctx, &learning.IDRequest{
		Id:     classroomID,
		UserID: student.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	if !res.Status {
		return &models.CurrentSubscription{
			Status: res.Status,
		}, nil
	}
	return &models.CurrentSubscription{
		Status:   res.Status,
		RestDays: int(res.RestDays),
		Subscription: &models.Subscription{
			ID:      res.Subscription.Id,
			MonthID: int(res.Subscription.MonthId),
			PaidAt:  int(res.Subscription.PaidAt),
		},
	}, nil
}

func (r *queryResolver) GetSubscriptions(ctx context.Context, classroomID string) ([]*models.Subscription, error) {

	student, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetSubscriptions(ctx, &learning.IDRequest{
		Id:     classroomID,
		UserID: student.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Subscriptions(res), nil

}

func (r *queryResolver) GetSubscriptionsByStudentID(ctx context.Context, classroomID string, studentID string) ([]*models.Subscription, error) {

	_, err := permisions.Admin(ctx)
	if err != nil {
		return nil, ErrApi(err)
	}

	res, err := r.LearningClient.GetSubscriptions(ctx, &learning.IDRequest{
		Id:     classroomID,
		UserID: studentID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Subscriptions(res), nil

}
