package interactive

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/clausthrane/captainslog/captainslog/entities"
)

func (cli *Cli) setupGroups() error {
	if groupIds, err := cli.app.Queries.ListGroups(); err != nil {
		return err
	} else {
		for idx, groupId := range groupIds {
			groupErr := cli.setupGroup(groupId, idx)
			if groupErr != nil {
				return groupErr
			}
		}
	}

	return nil
}

func (cli *Cli) setupGroup(groupId entities.TaskGroupID, position int) error {

	category := entities.CatagoryID("foo")
	if tasks, err := cli.app.Queries.ListTasks(groupId, category); err != nil {
		return err
	} else {
		cli.layoutTaskGroup(position, groupId, tasks)
	}
	return nil
}

func (cli *Cli) layoutTaskGroup(position int, groupID entities.TaskGroupID, tasks []*entities.Task) error {

	_, maxY := cli.gui.Size()

	view_id := fmt.Sprintf("g:%s", string(groupID))

	left_start := 1 + (position * 33)


	if v, err := cli.gui.SetView(view_id, left_start, 6, left_start+33, maxY - 3); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
			return err
		}
		v.Title = fmt.Sprintf("Group: %s", string(groupID))
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		for _, t := range tasks {
			fmt.Fprintln(v, t.Description)
		}
	}

	return nil
}
