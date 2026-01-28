package runner

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/mendersoftware/mender-server/tests/runner/client"
	"github.com/mendersoftware/mender-server/tests/runner/tests"
)

func init() {
	tests.AddTestCase("test_tiers", mainTestTiers)
}

func mainTestTiers(args *cli.Context) error {
	fmt.Printf("test_tiers starting\n")
	ctx := context.Background()
	c, err := client.NewClientWithResponses(args.String("server-url"))
	if err != nil {
		return errors.Wrap(err, "failed to create client")
	}

	basicAuth := b64.StdEncoding.EncodeToString([]byte(args.String("username") + ":" + args.String("password")))
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
	fmt.Printf("login response: jwt len=%d\n", len(jwt))
	return nil
}
