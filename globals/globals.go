package globals

import "gopkg.in/urfave/cli.v1"

var (
	Search  string
	Interactive string
	CategoryName string
)
var CommonFlags = []cli.Flag {
	cli.StringFlag{
		Name: "interactive, i",
		Value: "false",
		Usage: "Ask confirmation for certain operations",
		Destination: &Interactive,
	},
}
