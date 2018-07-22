package services

import (
	"errors"
	"testing"

	"github.com/restful/starter-kit/app"
	"github.com/restful/starter-kit/models"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	dao := newMockUserDAO()
	s := NewUserService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestUserService_Get(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Get(nil, 1)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "aaa", user.Name)
	}

	user, err = s.Get(nil, 100)
	assert.NotNil(t, err)
}

func TestUserService_Create(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Create(nil, &models.User{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 4, user.Id)
		assert.Equal(t, "ddd", user.Name)
	}

	// dao error
	_, err = s.Create(nil, &models.User{
		Id:   100,
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Create(nil, &models.User{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestUserService_Update(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Update(nil, 2, &models.User{
		Name: "ddd",
	})
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 2, user.Id)
		assert.Equal(t, "ddd", user.Name)
	}

	// dao error
	_, err = s.Update(nil, 100, &models.User{
		Name: "ddd",
	})
	assert.NotNil(t, err)

	// validation error
	_, err = s.Update(nil, 2, &models.User{
		Name: "",
	})
	assert.NotNil(t, err)
}

func TestUserService_Delete(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Delete(nil, 2)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, 2, user.Id)
		assert.Equal(t, "bbb", user.Name)
	}

	_, err = s.Delete(nil, 2)
	assert.NotNil(t, err)
}

func TestUserService_Query(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	result, err := s.Query(nil, 1, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, 2, len(result))
	}
}

func newMockUserDAO() userDAO {
	return &mockUserDAO{
		records: []models.User{
			{Id: 1, Name: "aaa"},
			{Id: 2, Name: "bbb"},
			{Id: 3, Name: "ccc"},
		},
	}
}

type mockUserDAO struct {
	records []models.User
}

func (m *mockUserDAO) Get(rs app.RequestScope, id int) (*models.User, error) {
	for _, record := range m.records {
		if record.Id == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockUserDAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	return m.records[offset : offset+limit], nil
}

func (m *mockUserDAO) Count(rs app.RequestScope) (int, error) {
	return len(m.records), nil
}

func (m *mockUserDAO) Create(rs app.RequestScope, user *models.User) error {
	if user.Id != 0 {
		return errors.New("Id cannot be set")
	}
	user.Id = len(m.records) + 1
	m.records = append(m.records, *user)
	return nil
}

func (m *mockUserDAO) Update(rs app.RequestScope, id int, user *models.User) error {
	user.Id = id
	for i, record := range m.records {
		if record.Id == id {
			m.records[i] = *user
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockUserDAO) Delete(rs app.RequestScope, id int) error {
	for i, record := range m.records {
		if record.Id == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}
