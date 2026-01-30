package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("test_self", mainTestSelf)
}

func mainTestSelf(t *testing.T, settings *TestSettings) error {
	ctx := context.Background()
	c, err := client.NewClientWithResponses(
		settings.ServerURL,
		client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+settings.jwt)
			return nil
		}),
	)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	r, err := c.ShowOwnUserDataWithResponse(
		ctx,
	)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 200, r.HTTPResponse.StatusCode)
	assert.Equal(t, settings.Username, r.JSON200.Email)

	t.Logf("test passed with user data: %+v\n", r.JSON200)
	return nil
}
