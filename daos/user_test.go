package daos

import (
	"testing"

	"github.com/nettyrnp/go-rest/app"
	"github.com/nettyrnp/go-rest/models"
	"github.com/nettyrnp/go-rest/migrate"
	"github.com/stretchr/testify/assert"
)

func TestUserDAO(t *testing.T) {
	db := migrate.ResetDB()
	dao := NewUserDAO()

	{
		// Get
		testDBCall(db, func(rs app.RequestScope) {
			user, err := dao.Get(rs, 2)
			assert.Nil(t, err)
			if assert.NotNil(t, user) {
				assert.Equal(t, 2, user.Id)
			}
		})
	}

	{
		// Create
		testDBCall(db, func(rs app.RequestScope) {
			user := &models.User{
				Id:   1000,
				Name: "tester",
			}
			err := dao.Create(rs, user)
			assert.Nil(t, err)
			assert.NotEqual(t, 1000, user.Id)
			assert.NotZero(t, user.Id)
		})
	}

	{
		// Delete
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 2)
			assert.Nil(t, err)
		})
	}

	{
		// Delete with error
		testDBCall(db, func(rs app.RequestScope) {
			err := dao.Delete(rs, 99999)
			assert.NotNil(t, err)
		})
	}

	{
		// Query
		testDBCall(db, func(rs app.RequestScope) {
			users, err := dao.Query(rs, 1, 3)
			assert.Nil(t, err)
			assert.Equal(t, 3, len(users))
		})
	}

	{
		// Count
		testDBCall(db, func(rs app.RequestScope) {
			count, err := dao.Count(rs)
			assert.Nil(t, err)
			assert.NotZero(t, count)
		})
	}
}
