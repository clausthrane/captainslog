package cli

import (
	"github.com/jroimartin/gocui"
	"os"
	"log"
	"fmt"
	"github.com/clausthrane/captainslog/services"
)

var logger = log.New(os.Stdout, " ", log.Ldate | log.Ltime | log.Lshortfile)

type Cli struct {
	taskService *services.TaskService
	gui         *gocui.Gui
}

func NewCli(taskService *services.TaskService) *Cli {

	g, err := gocui.NewGui()
	cli := &Cli{taskService, g}

	if err != nil {
		logger.Panicln(err)
	}

	g.Cursor = true
	g.Mouse = true

	defer g.Close()

	g.SetManagerFunc(cli.setupLayout())

	if err := keybindings(g); err != nil {
		logger.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		logger.Panicln(err)
	}


	return cli
}


func (cli *Cli) setupLayout() func(g *gocui.Gui) error {

	return func(g *gocui.Gui) error {

		cli.setupMenu()

		if taskGroup, err := cli.taskService.Load(); err != nil {
			for groupID, tasks := range taskGroup.Tasks {
				cli.layoutTaskGroup(groupID, tasks)
			}
		} else {
			return err
		}

		return nil
	}

}


func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func showMsg(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX / 2 - 10, maxY / 2, maxX / 2 + 10, maxY / 2 + 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
	}
	return nil
}

func delMsg(g *gocui.Gui, v *gocui.View) error {
	if err := g.DeleteView("msg"); err != nil {
		return err
	}
	return nil
}

