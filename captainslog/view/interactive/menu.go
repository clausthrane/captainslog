package interactive

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func (cli *Cli) setupMenu() error {

	maxX, maxY := cli.gui.Size()

	if v, err := cli.gui.SetView("frame", 0, 0, maxX-2, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Captains Log"
	}

	if _, err := cli.gui.SetView("menu", 1, 1, maxX-3, 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	if menu_item, err := cli.gui.SetView("menu_item", 2, 2, 17, 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(menu_item, "New task: F1")
	}

	return nil
}
