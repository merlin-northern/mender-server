package tests

import (
	"context"
	b64 "encoding/base64"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/mendersoftware/mender-server/tests/runner/client"
)

func init() {
	AddTestCase("test_login", mainTest)
}

func mainTest(t *testing.T, settings TestSettings) error {
	t.Logf("login test starting\n")
	ctx := context.Background()
	c, err := client.NewClientWithResponses(settings.ServerURL)
	if err != nil {
		return errors.Wrap(err, "failed to create client")
	}

	basicAuth := b64.StdEncoding.EncodeToString([]byte(settings.Username + ":" + settings.Password))
	l, err := c.LoginWithResponse(ctx, client.LoginJSONRequestBody{}, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", "Basic "+basicAuth)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "failed LoginJSONRequestBody")
	}

	if l.HTTPResponse == nil {
		return errors.Errorf("l.HTTPResponse is nil")
	}

	if l.HTTPResponse != nil && l.HTTPResponse.StatusCode != 200 {
		return errors.Errorf("failed LoginJSONRequestBody: %d", l.HTTPResponse.StatusCode)
	}

	jwt := string(l.Body)
	assert.True(t, len(jwt) > 0)
	t.Logf("test passed with login response: jwt len=%d\n", len(jwt))
	return nil
}
