package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
	"log"
)

func (r *queryResolver) ClassRooms(ctx context.Context, subjectID string) ([]*models.ClassRoom, error) {

	res, err := r.LearningClient.GetClassRooms(ctx, &learning.IDRequest{
		Id: subjectID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.ClassRooms(res), nil
}

func (r *queryResolver) ClassRoomsTeacher(ctx context.Context, teacherID string) ([]*models.ClassRoom, error) {

	_, err := permisions.CanRead(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetClassRoomsByTeacher(ctx, &learning.IDRequest{
		Id: teacherID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.ClassRooms(res), nil
}

func (r *queryResolver) MyClassRoomsTeacher(ctx context.Context) ([]*models.ClassRoom, error) {

	user, err := permisions.CanRead(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.LearningClient.GetClassRoomsByTeacher(ctx, &learning.IDRequest{
		Id: user.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	for _, classroom := range res.Classrooms {
		for _, ch := range classroom.Chapters.Chapters {
			for _, l := range ch.Lessons.Lessons {
				log.Println(l.Documents)

			}
		}
	}
	return gprcToGraph.ClassRooms(res), nil
}

func (r *queryResolver) ClassRoom(ctx context.Context, ClassRoomID string) (*models.ClassRoom, error) {
	log.Println("ClassRoom Started ....")
	ClassRoom, err := r.LearningClient.GetClassRoom(ctx, &learning.IDRequest{
		Id: ClassRoomID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.ClassRoom(ClassRoom), nil

}

func (r *mutationResolver) DeleteClassRoom(ctx context.Context, in models.IDRequest) (*models.OperationStatus, error) {

	//ctx.Value()
	_, err := r.LearningClient.DeleteClassRoom(ctx, &learning.IDRequest{
		Id: in.ID,
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

func (r *queryResolver) MyClassRooms(ctx context.Context, subjectID string) ([]*models.ClassRoom, error) {

	student, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.GetMyClassRooms(ctx, &learning.IDRequest{
		Id:     subjectID,
		UserID: student.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.ClassRooms(res), nil
}
