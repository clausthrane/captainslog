package interactive

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/clausthrane/captainslog/captainslog/entities"
	"sort"
)

func (cli *Cli) setupGroups() error {
	var err error
	if groupIds, err := cli.app.Queries.ListGroups(); err == nil {
		sort.Sort(groupIds)
		for idx, groupId := range groupIds {
			err := cli.setupGroup(groupId, idx)
			if err != nil {

				return err
			}
		}
	}

	return err
}

func (cli *Cli) setupGroup(groupId entities.TaskGroupID, position int) error {
	var err error
	if tasks, err := cli.app.Queries.ListTasks(groupId, entities.BlankCategory()); err == nil {
		cli.layoutTaskGroup(position, groupId, tasks)
	}
	return err
}

func (cli *Cli) layoutTaskGroup(position int, groupID entities.TaskGroupID, tasks []*entities.Task) error {

	_, maxY := cli.gui.Size()

	view_id := string(groupID) + string(len(tasks))

	left_start := 1 + (position * 33)

	if v, err := cli.gui.SetView(view_id, left_start, 6, left_start + 33, maxY - 3); err != nil {
		if err != gocui.ErrUnknownView {

			return err
		}
		v.Title = string(groupID)
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		for idx, t := range tasks {
			fmt.Fprintln(v, t.StringWthIdx(idx))
		}
		fmt.Fprintln(v, len(tasks))
	}

	return nil
}
