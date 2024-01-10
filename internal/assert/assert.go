package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %v; got %v", expected, actual)
	}
}

func StringContains(t *testing.T, actual, expectedSubsting string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubsting) {
		t.Errorf("expected %q to contain %q", actual, expectedSubsting)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("expected nil error; got %v", actual)
	}
}
