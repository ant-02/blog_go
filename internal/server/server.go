package server

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"zll.blog.com/internal/pb"
)

func RegisterAllServer(grpcServer *grpc.Server, db *gorm.DB) {
	pb.RegisterArticleServiceServer(grpcServer, NewArticleServer(db))
	pb.RegisterCategoryServiceServer(grpcServer, NewCategoryServer(db))
	pb.RegisterArticleCategoryServiceServer(grpcServer, NewArticleCategoryServer(db))
	pb.RegisterUserServiceServer(grpcServer, NewUserServer(db))
	pb.RegisterUserFollowServiceServer(grpcServer, NewUserFollowServer(db))
	pb.RegisterTagServiceServer(grpcServer, NewTagServer(db))
	pb.RegisterArticleTagServiceServer(grpcServer, NewArticleTagServer(db))
}
