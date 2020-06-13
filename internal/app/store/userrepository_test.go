package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yykhomenko/http-rest-api/internal/app/model"
	"github.com/yykhomenko/http-rest-api/internal/app/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
