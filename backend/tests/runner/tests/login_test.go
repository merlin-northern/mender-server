package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	openapi "github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("0_test_login", mainTestLogin)
}

func mainTestLogin(t *testing.T, settings *TestSettings) error {
	t.Logf("login test starting\n")
	auth := openapi.BasicAuth{
		UserName: settings.Username,
		Password: settings.Password,
	}
	ctx := context.WithValue(context.Background(), openapi.ContextBasicAuth, auth)
	token, r, err := settings.client.UserAdministrationManagementAPIAPI.Login(ctx).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.NotZero(t, len(token))
	assert.Equal(t, 200, r.StatusCode)
	settings.jwt = token
	t.Logf("test passed with %d\n", r.StatusCode)
	return nil
}
