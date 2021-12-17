package mocks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"totality/users/internal/files"
	"totality/users/totality"
)

var (
	notFound = errors.New("not found")
)

//  Client represent mocks client.
type Client struct {
	usersMap map[int64]totality.User
	mutex    *sync.RWMutex
}

//  GetUserByID retrives User details based on given ID.
func (c *Client) GetUserByID(ctx context.Context, id int64) (totality.User, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var v, ok = c.usersMap[id]
	if !ok {
		return totality.User{}, nil
	}

	return v, nil
}

// GetUsers retrives list of users.
func (c *Client) GetUsers(ctx context.Context, ids []int64) ([]totality.User, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var result []totality.User
	for _, id := range ids {
		var v, ok = c.usersMap[id]
		if ok {
			result = append(result, v)
		}
	}

	if len(result) <= 0 {
		return nil, notFound
	}

	return result, nil
}

// NewUserClient create new mocks client.
func NewUserClient() (totality.UserManager, error) {
	var bytes, err = files.ReadFromFile("/home/harvir/godev/totality-corp-assignments/mocks/users.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read users from file: %w", err)
	}

	var users []totality.User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	var m = make(map[int64]totality.User)
	for _, user := range users {
		m[user.ID] = user
	}

	return &Client{
		usersMap: m,
		mutex:    &sync.RWMutex{},
	}, nil

}
