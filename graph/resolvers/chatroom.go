package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
)

func (r *queryResolver) Room(ctx context.Context, roomID string) (*models.Room, error) {

	//user, err := permisions.User(ctx)
	//if err != nil {
	//	return nil, err
	//}
	//ctx.Value()
	room, err := r.LearningClient.GetChatRoom(ctx, &learning.IDRequest{
		Id:     roomID,
		UserID: "", //user.Id
	})

	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.Room(room), nil
}

func (r *mutationResolver) CreateMessage(ctx context.Context, in models.MessageInput) (*models.OperationStatus, error) {

	//user, err := permisions.Student(ctx)
	//if err != nil {
	//	user, err = permisions.Teacher(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	_, err := r.LearningClient.AddMessageToChatRoom(ctx, &learning.AddMessage{
		RoomID:    in.RoomID,
		Content:   in.Message,
		UserID:    in.OwnerID,
		IsTeacher: in.IsTeacher,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return &models.OperationStatus{Succeeded: true}, nil

}
