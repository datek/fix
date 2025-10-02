package examples_test

import (
	"testing"

	"github.com/datek/fix"
	"github.com/datek/fix/examples"
)

func TestUserStore(t *testing.T) {
	t.Run("Creates user", func(t *testing.T) {
		// given
		name := "John"
		store := fixtureUserStore(t)

		// when
		err := store.CreateUser(name)

		// then
		equal(t, nil, err)
		statements := *fixtureStatements(t)
		equal(t, 1, len(statements))
		equal(t, "INSERT "+name, statements[0])
	})

	t.Run("Deletes user", func(t *testing.T) {
		// given
		name := fixtureExistingUser(t)
		store := fixtureUserStore(t)

		// when
		err := store.DeleteUser(name)

		// then
		equal(t, nil, err)
		statements := *fixtureStatements(t)
		equal(t, 1, len(statements))
		equal(t, "DELETE "+name, statements[0])
	})
}

var fixtureExistingUser = fix.New(func(t *testing.T) string {
	t.Helper()
	userStore := fixtureUserStore(t)

	username := "Doe"
	err := userStore.CreateUser(username)

	if err != nil {
		t.Fatalf("Error when creating user, %v", err)
		return ""
	}

	statements := fixtureStatements(t)
	*statements = []string{}

	return username
})

var fixtureUserStore = fix.New(func(t *testing.T) examples.UserStore {
	t.Helper()
	db := fixtureMockDB(t)

	return examples.NewUserStore(db)
})

var fixtureMockDB = fix.New(func(t *testing.T) examples.DB {
	t.Helper()
	statements := fixtureStatements(t)

	return examples.NewMockDB(func(statement string) error {
		*statements = append(*statements, statement)
		return nil
	})
})

var fixtureStatements = fix.New(func(t *testing.T) *[]string {
	t.Helper()
	return &[]string{}
})

func equal[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
