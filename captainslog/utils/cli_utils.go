package utils

import "github.com/urfave/cli"

func HasStringArg(c *cli.Context, key string) bool {
	return EmptyString(c.String(key))
}
