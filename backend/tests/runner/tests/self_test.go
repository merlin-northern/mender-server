package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("test_self", mainTestSelf)
}

func mainTestSelf(t *testing.T, settings *TestSettings) error {
	ctx := context.WithValue(context.Background(), openapi.ContextAccessToken, settings.jwt)
	body, r, err := settings.client.UserAdministrationManagementAPIAPI.ShowOwnUserData(ctx).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.StatusCode)
	assert.Equal(t, settings.Username, body.Email)

	t.Logf("test passed with user data: %+v\n", body)
	return nil
}
