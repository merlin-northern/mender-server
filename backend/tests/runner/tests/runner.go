package tests

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type MainTestFunc func(context *cli.Context) error

var TestCases = make(map[string]MainTestFunc)

func AddTestCase(name string, main MainTestFunc) {
	TestCases[name] = main
}

func Main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "integration",
				Usage:  "Run the integration tests",
				Action: IntegrationRun,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "username",
						Usage:   "user name to use in case of running against existing accounts",
						EnvVars: []string{"MENDER_USERNAME"},
					},
					&cli.StringFlag{
						Name:    "password",
						Usage:   "password to use in case of running against existing accounts",
						EnvVars: []string{"MENDER_PASSWORD"},
					},
					&cli.StringFlag{
						Name:     "server-url",
						Usage:    "full URL for the instance for testing, example: https://staging.hosted.mender.io",
						Required: true,
						EnvVars:  []string{"MENDER_URL"},
					},
				},
			},
		},
	}
	app.Usage = "Backend Integration Tests Runner"
	app.Version = "1.0"
	app.Action = IntegrationRun
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func IntegrationRun(context *cli.Context) error {
	for name, test := range TestCases {
		fmt.Printf("Running test %s\n", name)
		err := test(context)
		if err != nil {
			return err
		}
	}

	return nil
}
