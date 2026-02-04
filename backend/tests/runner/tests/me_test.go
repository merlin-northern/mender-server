package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	openapi "github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("test_me", mainTestMe)
}

func mainTestMe(t *testing.T, settings *TestSettings) error {
	t.Logf("login test starting\n")
	ctx := context.WithValue(context.Background(), openapi.ContextAccessToken, settings.jwt)
	body, r, err := settings.client.UserAdministrationManagementAPIAPI.ShowMyUserSettings(ctx).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.NotNil(t, body)
	assert.Equal(t, 200, r.StatusCode)
	t.Logf("test passed with data: %+v\n", body)
	return nil
}
