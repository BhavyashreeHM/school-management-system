package handlers

import pb "employee-management-system/proto/gen"

type Server struct {
	pb.UnimplementedExecsServiceServer
	pb.UnimplementedTeachersServiceServer
	pb.UnimplementedStudentsServiceServer
}
