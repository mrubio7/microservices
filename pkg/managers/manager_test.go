package managers_test

import (
	testutil "ibercs/internal/test"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	schemas := []string{"matches", "players", "users", "teams"}

	testutil.SetupTestDBs(schemas)

	code := m.Run()

	testutil.CleanupTestDBs()

	os.Exit(code)
}
