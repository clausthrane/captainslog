package main

import (
	"fmt"
	"github.com/clausthrane/captainslog/captainslog"
	"github.com/clausthrane/captainslog/captainslog/config"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"github.com/clausthrane/captainslog/captainslog/view/interactive"
	"github.com/urfave/cli"
	"os"
	"time"
	"github.com/clausthrane/captainslog/captainslog/utils"
	"strconv"
	"strings"
)

var log = config.Logger
var DEFAULT_DATADIR = "mydata"
var DEFAULT_PROJECT = entities.ProjectID("project")

func main() {

	commandline := cli.NewApp()
	commandline.Name = fmt.Sprintf("Captains Log: %s", config.GetString("captain"))
	commandline.Usage = "record your stuff"
	commandline.HelpName = "Space: the final frontier. These are the voyages of the starship Enterprise"
	commandline.EnableBashCompletion = true

	commandline.Commands = []cli.Command{
		{
			Name:    "done",
			Aliases: []string{"a"},
			Usage:   "create a past record",
			Action:  complete_task,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "time, t",
					Usage: "minutes used",
				},
				cli.StringFlag{
					Name:  "cat, c",
					Usage: "a catagory",
				},
				cli.StringFlag{
					Name:  "message, m",
					Usage: "a message",
				},
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"t"},
			Usage:   "remove note of a voyage",
			Action:  remove_task,
		}, {
			Name:    "todo",
			Aliases: []string{"t"},
			Usage:   "make note of a future voyage",
			Action:  todo_task,
		}, {
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list records",
			Subcommands: []cli.Command{
				{
					Name: "today",
					Usage: "show content for a given day",
					Action: list_day,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "cat, c",
							Usage: "a catagory",
						},
					},
				}, {
					Name: "week",
					Usage: "show content for a given week",
					Action: list_week,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "cat, c",
							Usage: "a catagory",
						},
					},
				}, {
					Name: "all",
					Usage: "show content for a given week",
					Action: list_all,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "cat, c",
							Usage: "a catagory",
						},
					},
				},
			},
		},
		{
			Name:    "view",
			Action:  main_view,
		},
	}

	commandline.Run(os.Args)
}

func main_view(c *cli.Context) error {
	interactive.NewInteractiveCli(captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT))
	return nil
}

func todo_task(c *cli.Context) error {
	category := entities.CatagoryID(c.Args().Get(0))
	description := strings.Join(c.Args()[1:], " ")

	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	gid := entities.TodaysGroup()
	t, e := app.Commands.AddOpenTask(gid, description, category)

	if e != nil {
		panic(e)
	}
	if t != nil {
		println(t.String())
	}
	return nil
}

func remove_task(c *cli.Context) error {
	return nil
}

func complete_task(c *cli.Context) error {

	numberOfMinutes, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		numberOfMinutes = 15
	}

	timeUsed := time.Duration(numberOfMinutes) * time.Minute
	category := entities.CatagoryID(c.Args().Get(1))
	description := strings.Join(c.Args()[2:], " ")

	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	gid := entities.TodaysGroup()
	t, e := app.Commands.AddDoneTask(gid, description, category, timeUsed)

	if e != nil {
		panic(e)
	}
	if t != nil {
		println(t.String())
	}
	return nil
}

func list_all(c *cli.Context) error {
	//category := entities.CatagoryID(c.String("cat"))
	return nil
}

func list_week(c *cli.Context) error {
	//category := entities.CatagoryID(c.String("cat"))
	return nil
}

func list_day(c *cli.Context) error {
	category := entities.CatagoryID(c.String("cat"))

	when := time.Now().UTC()
	if hasStringArg(c, "date") {
		when = utils.ParseDate(c.String("date"))
	}

	groupId := entities.GroupByDate(when)
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	if tasks, err := app.Queries.ListTasks(groupId, category); err != nil {
		panic(err)
	} else {
		totalTime := tasks.SumTime()
		println(fmt.Sprintf("Time used in total: %s", utils.PrettyPrint(totalTime)))
		println(strings.Repeat("-", 10))
		for _, task := range tasks {
			println(task.String())
		}
	}
	return nil
}

func hasStringArg(c *cli.Context, key string) bool {
	return utils.EmptyString(c.String(key))
}

