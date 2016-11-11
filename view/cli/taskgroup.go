package cli

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"github.com/clausthrane/captainslog/entities"
)

func (cli *Cli) layoutTaskGroup(groupID entities.TaskGroupID, tasks []*entities.Task) error {

	if v, err := cli.gui.SetView("taskgroup", 2, 2, 22, 7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = string(groupID)
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		for _, t := range tasks {
			fmt.Fprintln(v, t.Description)
		}
	}

	return nil
}

