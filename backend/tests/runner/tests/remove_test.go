package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("test_login", mainTestRemove)
}

func mainTestRemove(t *testing.T, settings *TestSettings) error {
	ctx := context.WithValue(context.Background(), openapi.ContextAccessToken, settings.jwt)
	r, err := settings.client.UserAdministrationManagementAPIAPI.RemoveUser(ctx, "id").Execute()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 204, r.StatusCode)

	return nil
}
