package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ripolak/chk/chk"
	"github.com/urfave/cli/v2"
)

// HttpCommand represents a command in the CLI that checks HTTP or HTTPS connection.
var HttpCommand = &cli.Command{
	Name:  "http",
	Usage: "Check if the target address is accessible via HTTP.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   1,
			Usage:   "Time to wait for a response in seconds.",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Value:   false,
			Usage:   "Specifies whether to show verbose information about the http result.",
		},
	},
	Action: httpAction,
}

func httpAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	if err != nil {
		return err
	}
	verbose := c.Bool("verbose")
	timeout := c.Int("timeout")
	res, err := chk.CheckHTTP(address, timeout)
	if err != nil {
		displayHttpErrorResult(address, err)
	} else {
		displayHttpResult(res.Response(), address, verbose)
	}
	return nil
}

func displayHttpResult(res *http.Response, address string, verbose bool) {
	if verbose {
		b, err := json.MarshalIndent(res, "", "	")
		if err != nil {
			fmt.Println("Failed to parse verbose info about the http request result.")
		} else {
			fmt.Println(string(b))
		}
	}
	fmt.Println(fmt.Sprintf("Successful http connection to %s. Response code that was received: %v", address, res.StatusCode))
}

func displayHttpErrorResult(address string, err error) {
	fmt.Println(fmt.Sprintf("HTTP connection to %s failed. The following error was received: '%s'", address, err))
}
