package auth

import "fmt"

type MemoryRepository struct {
	users map[string]*User
}

func NewMemoryRepository(users []*User) *MemoryRepository {
	repo := &MemoryRepository{
		users: make(map[string]*User),
	}

	for _, user := range users {
		repo.users[user.Username] = user
	}

	return repo
}

func (r *MemoryRepository) GetUser(username string) (*User, error) {
	found := r.users[username]
	if found == nil {
		return nil, fmt.Errorf("user not found")
	}

	return found, nil
}
