package main

import (
	"fmt"
	"github.com/dutchakdev/clif/category"
	"github.com/dutchakdev/clif/command"
	"github.com/dutchakdev/clif/db"
	"github.com/dutchakdev/clif/globals"
	"github.com/dutchakdev/clif/helpers"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strings"
)

func main() {

	app := cli.NewApp()
	app.Name = "clif"
	app.Usage = "CLI Favorites"
	app.Version = Version()
	app.EnableBashCompletion = true
	app.Flags = globals.CommonFlags

	app.Action = func(c *cli.Context) error {
		if (len(os.Args) >= 2) {
			command.RunCmdByPath(os.Args[1], os.Args[2:])
		} else {
			command.FindCmdInCategories("")
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "categories",
			Aliases: []string{"c"},
			Usage:   "Categories of commands",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new category",
					Action: func(c *cli.Context) error {
						category.CreateCategory()
						return nil
					},
					Flags: []cli.Flag {
						cli.StringFlag{
							Name: "category, c",
							Value: "",
							Usage: "Category name",
							Destination: &globals.CategoryName,
						},
						cli.StringFlag{
							Name: "interactive, i",
							Value: "false",
							Usage: "Ask confirmation for certain operations",
							Destination: &globals.Interactive,
						},
					},
				},
				{
					Name:  "rm",
					Usage: "remove category",
					Action: func(c *cli.Context) error {
						category.RemoveCategory(strings.Join(os.Args[3:], " "))
						return nil
					},
					Flags: globals.CommonFlags,
				},
				{
					Name:  "ls",
					Usage: "list category",
					Action: func(c *cli.Context) error {
						fmt.Println(strings.Join(category.GetCategories()[1:], "\n"))
						return nil
					},
				},
			},
		},
		{
			Name:    "command",
			Aliases: []string{"cmd"},
			Usage:   "Commands",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new command",
					Action: command.ActionAddCmd,
				},
				{
					Name:  "rm",
					Usage: "remove command",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				//{
				//	Name:  "ls",
				//	Usage: "list command",
				//	Action: func(c *cli.Context) error {
				//		fmt.Println("new task template: ", c.Args().First())
				//		return nil
				//	},
				//},
			},
		},
	}

	err := app.Run(os.Args)
	helpers.CheckErr(err)
}

func init() {
	db.PrepareDatabase()
}

