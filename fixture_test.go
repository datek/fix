package fix_test

import (
	"testing"

	"github.com/datek/fix"
)

func TestFixture(topT *testing.T) {
	topT.Run("Returns fixture result", func(t *testing.T) {
		// given
		fixture := fix.New(func(t *testing.T) string {
			return "haha"
		})

		// when
		result := fixture.Value(t)

		// then
		equal(t, "haha", result)
	})

	topT.Run("Runs fixture only once per test context", func(t *testing.T) {
		// given
		fixtureCreationCount := 0
		fixture := fix.New(func(t *testing.T) int {
			fixtureCreationCount++
			return 1
		})

		// when
		fixture.Value(t)
		fixture.Value(t)

		// then
		equal(t, 1, fixtureCreationCount)
	})

	topT.Run("Fixture is calling the dependency fixture", func(t *testing.T) {
		// given
		fixture1 := fix.New(func(t *testing.T) int {
			return 1
		})

		fixture2 := fix.New(func(t *testing.T) int {
			return fixture1.Value(t) + 2
		})

		// when
		result := fixture2.Value(t)

		// then
		equal(t, 3, result)
	})

	topLevelFixtureCallCount := 0
	topLevelFixture := fix.New(func(t *testing.T) string {
		topLevelFixtureCallCount++
		return "!yay!"
	})

	topT.Run("Mixed scope fixtures are working", func(t *testing.T) {
		// given
		fixture1 := fix.New(func(t *testing.T) string {
			return topLevelFixture.Value(topT) + " this works too "
		})

		fixture2 := fix.New(func(t *testing.T) string {
			return fixture1.Value(t) + topLevelFixture.Value(topT)
		})

		// when
		result := fixture2.Value(t)

		// then
		equal(t, "!yay! this works too !yay!", result)
		equal(t, 1, topLevelFixtureCallCount)
	})

	counter := 0
	topT.Run("Cleanup - preparation", func(t *testing.T) {
		fixture := fix.New(func(t *testing.T) any {
			t.Cleanup(func() { counter++ })
			return nil
		})

		fixture.Value(t)
		equal(t, 0, counter)
	})

	topT.Run("Cleanup - test", func(t *testing.T) {
		equal(t, 1, counter)
	})
}

func equal[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
}
