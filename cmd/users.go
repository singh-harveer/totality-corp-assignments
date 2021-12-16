package main

import (
	"context"
	"log"
	"totality/users/client"
	"totality/users/totality"
)

func main() {
	var userClient, err = client.NewUserClient(context.Background(), 9001, "localhost")
	if err != nil {
		log.Fatalf("failed to create client :%v", err)
	}

	var user totality.User
	user, err = userClient.GetUserByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("failed to get user by id : %v", err)
	}

	log.Printf("user details for userID: %d", user.ID)

	var users []totality.User
	users, err = userClient.GetUsers(context.Background(), []int64{1, 2, 3, 4})
	if err != nil {
		log.Fatalf("failed to get user by id : %v", err)
	}

	log.Printf("users details:\n %v", users)
}
