package handlers

import (
	"context"
	"employee-management-system/internal/models"
	"employee-management-system/internal/reposioriy/mongodb"
	"employee-management-system/pkg/utils"
	pb "employee-management-system/proto/gen"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) AddTeachers(ctx context.Context, req *pb.Teachers) (*pb.Teachers, error) {
	client, err := mongodb.CreateMongoClient()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Internal Server Error")
	}
	defer client.Disconnect(ctx)

	newTeachers := make([]*models.Teacher, len(req.GetTeachers()))
	for i, pbteacher := range req.GetTeachers() {
		modelTeacher := models.Teacher{}
		pbVal := reflect.ValueOf(pbteacher).Elem()
		// fmt.Println("pbVal:", pbVal)
		modelVal := reflect.ValueOf(&modelTeacher).Elem()

		//* this is how we can set the value of modelTeacher's FirstName field to the value of pbteacher's FirstName field using reflection
		// fmt.Println("ModelVal:", modelVal)
		// fmt.Println("pbVal num field",pbVal.NumField())
		// fmt.Println("pbVal field[0]",pbVal.Type().Field(0).Name)
		// fmt.Println("pbVal field[4]",pbVal.Type().Field(4).Name)
		// fmt.Println("pbVal field[5]",pbVal.Type().Field(5).Name)
		// fmt.Println("pbVal field[6]",pbVal.Type().Field(6).Name)
		// fmt.Println("pbVal field[7]",pbVal.Type().Field(7).Name)
		// fmt.Println("pbVal field[8]",pbVal.Type().Field(8).Name)
		// fmt.Println("ModelVal FieldByName:",modelVal.FieldByName("FirstName"))
		// modelVal.FieldByName(pbVal.Type().Field(4).Name).Set(pbVal.Field(4))
		// fmt.Println("ModelVal NewValue",modelVal.FieldByName(pbVal.Type().Field(4).Name))

		//map protobf to model using reflection
		for i := 0; i < pbVal.NumField(); i++ {
			pbField := pbVal.Field(i)
			fieldName := pbVal.Type().Field(i).Name

			modelField := modelVal.FieldByName(fieldName)
			if modelField.IsValid() && modelField.CanSet() {
				modelField.Set(pbField)
			}
			// else{
			// 	fmt.Println("Field %s is not valid or cannot be set",fieldName)
			// }
		}
		// newTeachers = append(newTeachers, &modelTeacher)
		// ?or
		newTeachers[i] = &modelTeacher
	}

	var addedTeachers []*pb.Teacher
	for _, teacher := range newTeachers {
		result, err := client.Database("school").Collection("teachers").InsertOne(ctx, teacher)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Failed to insert teacher")
		}
		objectId, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			teacher.Id = objectId.Hex()
		}

		//map back to protobuf model
		pbTeacher := &pb.Teacher{}
		modelVal := reflect.ValueOf(*teacher)
		pbVal := reflect.ValueOf(pbTeacher).Elem()

		for i := 0; i < modelVal.NumField(); i++ {
			modelField := modelVal.Field(i)
			modelfieldType := modelVal.Type().Field(i)
			pbField := pbVal.FieldByName(modelfieldType.Name)
			if pbField.IsValid() && pbField.CanSet() {
				pbField.Set(modelField)
			}
		}

		addedTeachers = append(addedTeachers, pbTeacher)
	}

	return &pb.Teachers{Teachers: addedTeachers}, nil

}
