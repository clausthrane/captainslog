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
	"github.com/fatih/color"
)

var log = config.Logger
var DEFAULT_DATADIR = "mydata"
var DEFAULT_PROJECT = entities.ProjectID("project")
var todoGroupId = entities.TaskGroupID("todo")

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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "date, d",
					Usage: "a specific day",
				},
			},
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
					Usage: "show content for today",
					Action: list_day,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "cat, c",
							Usage: "a catagory",
						},
						cli.StringFlag{
							Name:  "date, d",
							Usage: "a specific day",
						},
					},
				}, {
					Name: "week",
					Usage: "show content for this week",
					Action: list_week,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "cat, c",
							Usage: "a catagory",
						},
					},
				}, {
					Name: "todo",
					Usage: "show all uncompleted tasks",
					Action: list_todo,
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
	var description string
	if c.NArg() > 1 {
		description = strings.Join(c.Args()[1:], " ")
	}

	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	t, e := app.Commands.AddOpenTask(todoGroupId, description, category)

	if e != nil {
		panic(e)
	}
	if t != nil {
		println(t.String())
	}
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
		println(t.StringWthIdx(0))
	}
	return nil
}

func list_week(c *cli.Context) error {
	//category := entities.CatagoryID(c.String("cat"))
	return nil
}

func list_todo(c *cli.Context) error {
	category := entities.CatagoryID(c.String("cat"))
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	title := "TODO"
	var err error
	if tasks, err := app.Queries.ListTasks(todoGroupId, category); err == nil {
		printList(tasks, title)
	}
	return err
}

func list_day(c *cli.Context) error {
	category := entities.CatagoryID(c.String("cat"))

	when := time.Now().UTC()
	if hasStringArg(c, "date") {
		when = utils.ParseDate(c.String("date"))
	}

	groupId := entities.GroupByDate(when)
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	title := entities.DateString(when)
	var err error
	if tasks, err := app.Queries.ListTasks(groupId, category); err == nil {
		printList(tasks, title)
	}
	return err
}

func remove_task(c *cli.Context) error {
	var groupId entities.TaskGroupID

	key := c.Args().Get(0)
	switch {
	case key == "today":
		groupId = entities.TodaysGroup()
	case key == "todo":
		groupId = todoGroupId
	}

	idx, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return err
	}
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)
	return app.Commands.RemoveTask(groupId, idx)
}

func hasStringArg(c *cli.Context, key string) bool {
	return utils.EmptyString(c.String(key))
}

func printList(tasks entities.TaskList, title string) {
	white := color.New(color.FgHiWhite).SprintFunc()
	println(fmt.Sprintf("Showing: %s", white(title)))
	totalTime := tasks.SumTime()
	println(fmt.Sprintf("Time used in total: %s", utils.PrettyPrint(totalTime)))
	println(strings.Repeat("-", 10))
	for idx, task := range tasks {
		println(task.StringWthIdx(idx))
	}
}
