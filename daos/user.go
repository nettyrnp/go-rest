package daos

import (
	"github.com/nettyrnp/go-rest/app"
	"github.com/nettyrnp/go-rest/models"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/nettyrnp/go-rest/errors"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get reads the user with the specified ID from the database.
func (dao *UserDAO) Get(rs app.RequestScope, id int) (*models.User, error) {
	var user models.User
	if rs.UserRole() == "user" {
		err := rs.Tx().Select().Where(dbx.HashExp{"role": "user"}).Model(id, &user)
		return &user, err
	} else if rs.UserRole() == "admin" {
		err := rs.Tx().Select().Model(id, &user)
		return &user, err
	} else {
		panic("Unexpected user role:" + rs.UserRole())
	}
}

// Create saves a new user record in the database.
// The User.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(rs app.RequestScope, user *models.User) error {
	user.Id = 0
	return rs.Tx().Model(user).Insert()
}

// Delete deletes a user with the specified ID from the database.
func (dao *UserDAO) Delete(rs app.RequestScope, id int) error {
	if rs.UserRole() == "admin" {
		user, err := dao.Get(rs, id)
		if err != nil {
			return err
		}
		return rs.Tx().Model(user).Delete()
	} else {
		return errors.Unauthorized("Only admins are allowed for this operation")
	}
}

// Query retrieves the user records with the specified offset and limit from the database.
func (dao *UserDAO) Query(rs app.RequestScope, offset, limit int) ([]models.User, error) {
	users := []models.User{}
	err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&users)
	return users, err
}

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count(rs app.RequestScope) (int, error) {
	var count int
	err := rs.Tx().Select("COUNT(*)").From("user").Row(&count)
	return count, err
}
