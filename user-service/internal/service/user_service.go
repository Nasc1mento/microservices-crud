package service

import (
	"context"

	"microservices-crud/user-service/internal/db/repo"
	user "microservices-crud/user-service/internal/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	q *repo.Queries
	user.UnimplementedUserServiceServer
}

func NewServer(q *repo.Queries) *UserServer {
	return &UserServer{q: q}
}

func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.UserResponse, error) {
	created, err := s.q.CreateUser(ctx, repo.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &user.UserResponse{
		Id:    created.ID.String(),
		Name:  created.Name,
		Email: created.Email,
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse id: %v", err)
	}

	err = s.q.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{
		Id: id.String(),
	}, nil
}

func (s *UserServer) GetUsers(ctx context.Context, req *user.GetUsersRequest) (*user.GetUsersResponse, error) {
	users, err := s.q.ListUsers(ctx, repo.ListUsersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	if err != nil {
		return nil, err
	}

	var resp user.GetUsersResponse
	for _, u := range users {
		resp.Users = append(resp.Users, &user.UserResponse{
			Id:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return &resp, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.UserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse id: %v", err)
	}

	u, err := s.q.GetUserById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user with id %s not found", id)
	}

	return &user.UserResponse{
		Id:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not parse id: %v", err)
	}

	u, err := s.q.GetUserById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user with id %s not found", id)
	}

	err = s.q.UpdateUser(ctx, repo.UpdateUserParams{
		ID:       u.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update user: %v", err)
	}

	return &user.UserResponse{
		Id:    u.ID.String(),
		Name:  req.Name,
		Email: req.Email,
	}, nil
}
