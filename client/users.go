package client

import (
	"context"
	"fmt"
	"log"
	"totality/users/totality"
	"totality/users/userpb"

	"google.golang.org/grpc"
)

// Client represents totality's GRPC Client.
type Client struct {
	grpcConn userpb.UserServiceClient
}

// GetUserByID retrieves user for given userID.
func (c *Client) GetUserByID(ctx context.Context, id int64) (totality.User, error) {
	var user, err = c.grpcConn.GetUserByID(ctx, &userpb.GetUserByIDRequest{Id: id})
	if err != nil {
		// TODO- handle error using GRPC error code.
		return totality.User{}, err
	}

	return totality.User{
		ID:      user.Id,
		Name:    user.Name,
		Height:  user.Height,
		City:    user.City,
		Phone:   user.Phone,
		Married: user.Married,
	}, nil
}

// GetUsers retrives all users.
func (c *Client) GetUsers(ctx context.Context, ids []int64) ([]totality.User, error) {
	var users, err = c.grpcConn.GetUsers(ctx, &userpb.GetUsersRequest{Ids: ids})
	if err != nil {
		// TODO- handle error using GRPC error code.
		return nil, err
	}

	var result []totality.User
	for _, user := range users.Users {
		result = append(result, totality.User{
			ID:      user.Id,
			Name:    user.Name,
			Height:  user.Height,
			Phone:   user.Phone,
			Married: user.Married,
		})
	}

	return result, nil
}

// NewUserClient creates new User grpc service client.
func NewUserClient(ctx context.Context, port int, host string) (*Client, error) {
	var conn, err = grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Client{
		grpcConn: userpb.NewUserServiceClient(conn),
	}, nil
}
