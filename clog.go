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
	"errors"
)

var log = config.Logger
var DEFAULT_DATADIR = "mydata"
var DEFAULT_PROJECT = entities.ProjectID("project")
var todoGroupId = entities.TaskGroupID("todo")
var red = color.New(color.FgHiRed).SprintFunc()

func main() {

	commandline := cli.NewApp()
	commandline.Name = fmt.Sprintf("Captains %s's Log", config.GetString("captain"))
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
			//Action:  remove_task,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "date, d",
					Usage: "a specific day",
				},
			},
			Subcommands: []cli.Command{
				{
					Name: "today",
					Action: remove_task_from_current_day,
				},
				{
					Name: "todo",
					Action: remove_todo_task,
				},
			},
		}, {
			Name:    "todo",
			Aliases: []string{"t"},
			Usage:   "this will make note of a future voyage",
			Action:  todo_task,
			Subcommands: []cli.Command{
				{
					Name: "done",
					Action: complete_todo,
				},
			},
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

func foo(c *cli.Context) error {
	print("foo")
	return nil
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

	var err error
	if t, err := app.Commands.AddOpenTask(todoGroupId, description, category); err == nil {
		println(t.String())
	}
	return err
}

func complete_task(c *cli.Context) error {
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	timeUsed := utils.StringAsDuration(c.Args().Get(0), time.Minute)
	category := entities.CatagoryID(c.Args().Get(1))
	description := strings.Join(c.Args()[2:], " ")
	gid := entities.TodaysGroup()

	var err error
	if t, err := app.Commands.AddDoneTask(gid, description, category, timeUsed); err == nil {
		println(t.StringWthIdx(0))
	}
	return err
}

func complete_todo(c *cli.Context) error {
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	todoIdx, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		println(red("Which todo do you want to complete?"))
		return err
	}
	timeUsed := utils.StringAsDuration(c.Args().Get(1), time.Minute)

	task, err := app.Queries.GetTask(todoGroupId, todoIdx)
	if err == nil {
		err = app.Commands.RemoveTask(todoGroupId, todoIdx)
		if err == nil {
			groupId := entities.TodaysGroup()
			task, err = app.Commands.AddDoneTask(groupId, task.Description, task.Category, timeUsed)
			print(task.FormatDone(todoIdx))
		}
	}
	return err

}

func list_week(c *cli.Context) error {

	var err error
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)

	allWeek := make(map[entities.CatagoryID][]*entities.Task)

	for _, d := range utils.DatesInThisWeek() {
		groupId := entities.GroupByDate(d)
		if tasks, err := app.Queries.ListTasks(groupId, entities.BlankCategory()); err == nil {
			for _, task := range tasks {
				items := allWeek[task.Category]
				if items == nil {
					items = make([]*entities.Task, 0)
				}
				items = append(items, task)
				allWeek[task.Category] = items
			}
		}
	}

	for key, value := range allWeek {
		println(key.String())
		for _, task := range value {
			println(fmt.Sprintf("- %s", task.Description))
		}
	}

	return err
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
	if utils.HasStringArg(c, "date") {
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

func remove_task_from_current_day(c *cli.Context) error {
	groupId := entities.TodaysGroup()

	idx, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		return errors.New("Make sure to provide an idx of the task to remove")
	}
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)
	return app.Commands.RemoveTask(groupId, idx)
}


func remove_todo_task(c *cli.Context) error {
	groupId := todoGroupId

	idx, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		return errors.New("Make sure to provide an idx of the task to remove")
	}
	app := captainslog.NewCaptainsLog(c, DEFAULT_DATADIR, DEFAULT_PROJECT)
	return app.Commands.RemoveTask(groupId, idx)
}


func printList(tasks entities.TaskList, title string) {
	white := color.New(color.FgHiWhite).SprintFunc()
	totalTime := tasks.SumTime()

	line := strings.Repeat(white("-"), 80)
	println(line)
	println(fmt.Sprintf("Showing: %s", white(title)))
	if ! (totalTime == 0) {
		println(fmt.Sprintf("Time used in total: %s", utils.PrettyPrint(totalTime)))
	}
	println(line)
	for idx, task := range tasks {
		println(task.StringWthIdx(idx))
	}
	println(line)
}
