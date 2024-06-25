package resolvers

import (
	"context"
	"github.com/les-cours/learning-api/api/learning"
	"github.com/les-cours/learning-api/graph/models"
	gprcToGraph "github.com/les-cours/learning-api/grpcToGraph"
	"github.com/les-cours/learning-api/permisions"
)

func (r *queryResolver) ClassRooms(ctx context.Context, subjectID string) ([]*models.ClassRoom, error) {

	_, err := permisions.User(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.LearningClient.GetClassRooms(ctx, &learning.IDRequest{
		Id: subjectID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.ClassRooms(res), nil
}

func (r *queryResolver) ClassRoomsTeacher(ctx context.Context, teacherID string) ([]*models.ClassRoom, error) {

	user, err := permisions.Teacher(ctx)
	if err != nil {
		return nil, err
	}
	if !permisions.CanRead(user) {
		return nil, permisions.ErrPermissionDenied
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

	user, err := permisions.Teacher(ctx)
	if err != nil {
		return nil, err
	}
	res, err := r.LearningClient.GetClassRoomsByTeacher(ctx, &learning.IDRequest{
		Id: user.ID,
	})
	if err != nil {
		return nil, ErrApi(err)
	}

	return gprcToGraph.ClassRooms(res), nil
}

func (r *queryResolver) ClassRoom(ctx context.Context, ClassRoomID string) (*models.ClassRoom, error) {

	user, err := permisions.Student(ctx)
	if err != nil {
		user, err = permisions.Teacher(ctx)
		if err != nil {
			return nil, err
		}
	}

	ClassRoom, err := r.LearningClient.GetClassRoom(ctx, &learning.IDRequest{
		Id:     ClassRoomID,
		UserID: user.ID,
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

func (r *mutationResolver) UpdateClassRoom(ctx context.Context, in models.UpdateClassRoomInput) (*models.ClassRoom, error) {

	user, err := permisions.Teacher(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.LearningClient.UpdateClassRoom(ctx, &learning.UpdateClassRoomRequest{
		TeacherID:         user.ID,
		ClassRoomID:       in.ClassRoomID,
		Title:             in.Title,
		Image:             in.Image,
		Price:             int32(in.Price),
		ArabicTitle:       in.ArabicTitle,
		Description:       in.Description,
		ArabicDescription: in.ArabicDescription,
	})
	if err != nil {
		return nil, ErrApi(err)
	}
	return gprcToGraph.ClassRoom(res), nil
}

/*
STUDENTS
*/

func (r *queryResolver) MyClassRooms(ctx context.Context, subjectID string) ([]*models.ClassRoom, error) {

	student, err := permisions.Student(ctx)
	if err != nil {
		return nil, err
	}

	var res = new(learning.ClassRooms)
	if subjectID == "" {
		res, err = r.LearningClient.GetMyChatRoom(ctx, &learning.IDRequest{
			Id:     student.ID,
			UserID: student.ID,
		})
		if err != nil {
			return nil, ErrApi(err)
		}
	} else {
		res, err = r.LearningClient.GetMyClassRooms(ctx, &learning.IDRequest{
			Id:     subjectID,
			UserID: student.ID,
		})
		if err != nil {
			return nil, ErrApi(err)
		}
	}

	return gprcToGraph.ClassRooms(res), nil
}
