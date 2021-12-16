package service

/**
This package will contains all GRPC service implementations.
**/

import (
	"context"
	"totality/users/totality"
	"totality/users/userpb"
)

// UserService implements UserServiceServicer.
type UserService struct {
	UserDB totality.UserManager
	userpb.UnimplementedUserServiceServer
}

// GetUserByID retrieves user for given userID.
func (u *UserService) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.User, error) {
	var user, err = u.UserDB.GetUserByID(context.Background(), req.Id)
	if err != nil {
		return &userpb.User{}, err
	}

	return &userpb.User{
		Id:      user.ID,
		Name:    user.Name,
		Height:  user.Height,
		Phone:   user.Phone,
		Married: user.Married,
	}, nil
}

// GetUsers retrives all users.
func (u *UserService) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	var users, err = u.UserDB.GetUsers(context.Background(), req.Ids)
	if err != nil {
		return nil, err
	}

	var result []*userpb.User
	for _, user := range users {
		result = append(result, &userpb.User{
			Id:      user.ID,
			Name:    user.Name,
			Height:  user.Height,
			Phone:   user.Phone,
			Married: user.Married,
		})
	}

	return &userpb.GetUsersResponse{
		Users: result,
	}, nil
}
