package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"log"
)

func (r *queryResolver) ClassRoomsBySubject(ctx context.Context, subjectID string) ([]*models.ClassRoom, error) {

	var ClassRooms []*models.ClassRoom

	res, err := r.LearningClient.GetClassRoomsBySubject(ctx, &learning.GetByUUIDRequest{
		Id: subjectID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	for _, ClassRoom := range res.Classrooms {
		ClassRooms = append(ClassRooms, gprcToGraph.ClassRoom(ClassRoom))
	}

	return ClassRooms, nil
}

func (r *queryResolver) ClassRoom(ctx context.Context, ClassRoomID string) (*models.ClassRoom, error) {

	log.Println("ClassRoom ... ")
	ClassRoom, err := r.LearningClient.GetClassRoom(ctx, &learning.GetByUUIDRequest{
		Id: ClassRoomID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.ClassRoom(ClassRoom), nil

}

func (r *mutationResolver) DeleteClassRoom(ctx context.Context, classRoomID string) (*models.OperationStatus, error) {

	//ctx.Value()
	_, err := r.LearningClient.DeleteClassRoom(ctx, &learning.GetByUUIDRequest{
		Id: classRoomID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return &models.OperationStatus{
		Succeeded: true,
	}, nil
}
