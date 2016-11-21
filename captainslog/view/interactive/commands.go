package interactive

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"github.com/clausthrane/captainslog/captainslog/entities"
)

func (cli *Cli) showAddTask(g *gocui.Gui, v *gocui.View) error {

	width := 80
	height := 50

	maxX, maxY := g.Size()
	x1 := maxX / 2 - (width / 2)
	y1 := maxY / 2 - (height / 2)
	x2 := x1 + (width)
	y2 := y1 + (height)

	if v, err := g.SetView("new_task", x1, y1, x2, y2); err != nil {

		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		fmt.Fprintln(v, "")
		v.Editable = true
		v.Wrap = true
		v.Title = "New task - Save: ctrl+s Exit: F9"
		if _, err := g.SetCurrentView("new_task"); err != nil {
			return err
		}
	}
	return nil
}

func (cli *Cli) saveCloseAddTask(g *gocui.Gui, v *gocui.View) error {
	p := make([]byte, 5)
	n, err := v.Read(p)
	if err != nil {
		return err
	}
	if n > 0 {
		gid := entities.TodaysGroup()
		cli.app.Commands.AddOpenTask(gid, string(p), entities.CatagoryID("foo"))
	}
	cli.closeAddTask(g, v)
	return nil
}

func (cli *Cli) closeAddTask(g *gocui.Gui, v *gocui.View) error {
	g.DeleteView("new_task")
	g.SetCurrentView("menu")
	return nil
}

func (cli *Cli) deleteTask() error {
	return nil
}

func (cli *Cli) deleteGroup() error {
	return nil
}

func (cli *Cli) addGroup() error {
	return nil
}