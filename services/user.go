package services

import (
	"github.com/nettyrnp/go-rest/app"
	"github.com/nettyrnp/go-rest/models"
)

// userDAO specifies the interface of the user DAO needed by UserService.
type userDAO interface {
	// Get returns the user with the specified user ID.
	Get(rs app.RequestScope, id int) (*models.User, error)
	// Query returns the list of users with the given offset and limit.
	Query(rs app.RequestScope, offset, limit int) ([]models.User, error)
	// Create saves a new user in the storage.
	Create(rs app.RequestScope, user *models.User) error
	// Delete removes the user with given ID from the storage.
	Delete(rs app.RequestScope, id int) error
	// Count returns the number of users.
	Count(rs app.RequestScope) (int, error)
}

// UserService provides services related with users.
type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get returns the user with the specified the user ID.
func (s *UserService) Get(rs app.RequestScope, id int) (*models.User, error) {
	return s.dao.Get(rs, id)
}

// Create creates a new user.
func (s *UserService) Create(rs app.RequestScope, model *models.User) (*models.User, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.Id)
}

// Delete deletes the user with the specified ID.
func (s *UserService) Delete(rs app.RequestScope, id int) (*models.User, error) {
	user, err := s.dao.Get(rs, id)
	if err != nil {
		return nil, err
	}
	err = s.dao.Delete(rs, id)
	return user, err
}

// Query returns the users with the specified offset and limit.
func (s *UserService) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	return s.dao.Query(rs, offset, limit)
}

// Count returns the number of users.
func (s *UserService) Count(rs app.RequestScope) (int, error) {
	return s.dao.Count(rs)
}
