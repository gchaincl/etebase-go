package etebase

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	envAccount  = "ETEBASE_TEST_ACCOUNT"
	envUsername = "ETEBASE_TEST_USERNAME"
	envEmail    = "ETEBASE_TEST_EMAIL"
	envPassword = "ETEBASE_TEST_PASSWORD"
)

func envSkip(t *testing.T, env string) {
	t.Skipf("Define %s to run this test", env)
}

func testUserFromEnv(t *testing.T) (User, string) {
	user, pass := User{
		Username: os.Getenv(envUsername),
		Email:    os.Getenv(envEmail),
	}, os.Getenv(envPassword)

	if user.Username == "" {
		envSkip(t, envUsername)
	}

	if user.Email == "" {
		envSkip(t, envEmail)
	}

	if pass == "" {
		envSkip(t, envPassword)
	}

	return user, pass
}

func TestClient(t *testing.T) {
	accountName := os.Getenv(envAccount)
	if accountName == "" {
		envSkip(t, envAccount)
	}
	user, password := testUserFromEnv(t)

	var (
		acc = NewAccount(
			NewClient(accountName, DefaultClientOptions),
		)
	)

	t.Run("Signup", func(t *testing.T) {
		assert.NoError(t,
			acc.Signup(user, password),
		)
	})

	t.Run("Login", func(t *testing.T) {
		require.NoError(t,
			acc.Login(user.Username, password),
		)
	})

	t.Run("Play", func(t *testing.T) {
		require.NoError(t, acc.Play())
	})
}
