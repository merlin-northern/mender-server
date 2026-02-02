package tests

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/mendersoftware/mender-server/tests/runner/client"
	"github.com/stretchr/testify/assert"
)

func init() {
	AddTestCase("test_z_me", mainTestMe)
}

func mainTestMe(t *testing.T, settings *TestSettings) error {
	t.Logf("login test starting\n")
	ctx := context.Background()
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	c, err := client.NewClientWithResponses(settings.ServerURL, client.WithHTTPClient(httpClient), client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Bearer "+settings.jwt)
		return nil
	}))
	r, err := c.ShowMyUserSettingsWithResponse(ctx)
	assert.NotNil(t, c)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	assert.Equal(t, 200, r.StatusCode())
	t.Logf("test passed with data: %+v\n", r.JSON200)
	return nil
}
