package imbotwebhookadapter

import (
	"os"
	"testing"
)

func IntegrationTest(t *testing.T) {
	t.Helper()

	if os.Getenv("INTEGRATION_TEST_ENABLED") != "true" {
		t.Skip("skipping integration tests, set environment variable INTEGRATION")
	}
}
