package localdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"totality/users/internal/files"
	"totality/users/totality"
)

var (
	notFound = errors.New("not found")
)

// Local represents local(users.json) sotrage.
type Local struct {
	users map[int64]totality.User
}

var _ totality.UserManager = (*Local)(nil)

func (l *Local) GetUserByID(ctx context.Context, id int64) (totality.User, error) {
	if _, ok := l.users[id]; ok {
		return l.users[id], nil
	}

	return totality.User{}, notFound
}

func (l *Local) GetUsers(ctx context.Context, ids []int64) ([]totality.User, error) {
	var result []totality.User
	for _, id := range ids {
		if _, ok := l.users[id]; ok {
			result = append(result, l.users[id])
		}
	}

	if len(result) <= 0 {
		return nil, notFound
	}

	return result, nil
}

func NewLocalDB() (totality.UserManager, error) {
	var bytes, err = files.ReadFromFile("dao/localdb/users.json")
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

	return &Local{
		users: m,
	}, nil
}
