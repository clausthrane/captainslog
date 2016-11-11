package cli

import (
	"github.com/jroimartin/gocui"
	"fmt"
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

	if v, err := cli.gui.SetView("menu_item", 2, 2, 7, 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "foo")
	}

	return nil
}
