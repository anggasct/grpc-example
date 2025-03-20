package service

import (
	"context"
	"fmt"

	pb "github.com/anggasct/post-service/pb/post"

	userPb "github.com/anggasct/post-service/pb/user"

	"github.com/anggasct/post-service/internal/db"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PostService struct {
	pb.UnimplementedPostServiceServer
	db             *db.Database
	userServiceUrl string
}

func NewPostService(database *db.Database, userServiceUrl string) *PostService {
	return &PostService{
		db:             database,
		userServiceUrl: userServiceUrl,
	}
}

func (s *PostService) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := s.db.GetPost(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %v", err)
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	conn, err := grpc.Dial(s.userServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	userClient := userPb.NewUserServiceClient(conn)
	userResp, err := userClient.GetUser(ctx, &userPb.GetUserRequest{Id: post.UserID})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:      post.ID,
			UserId:  post.UserID,
			Content: post.Content,
			User:    userResp.User,
		},
	}, nil
}
