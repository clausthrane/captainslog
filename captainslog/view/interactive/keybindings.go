package interactive

import "github.com/jroimartin/gocui"

var BIND_EVERYWHERE = ""

func (cli *Cli) keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding(BIND_EVERYWHERE, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding(BIND_EVERYWHERE, gocui.KeyF1, gocui.ModNone, cli.showAddTask); err != nil {
		return err
	}

	if err := g.SetKeybinding("side", gocui.KeyArrowDown, gocui.ModNone, nil); err != nil {
		return err
	}
	if err := g.SetKeybinding("side", gocui.KeyArrowUp, gocui.ModNone, nil); err != nil {
		return err
	}
	if err := g.SetKeybinding("new_task", gocui.KeyF9, gocui.ModNone, cli.closeAddTask); err != nil {
		return err
	}

	if err := g.SetKeybinding("new_task", gocui.KeyCtrlS, gocui.ModNone, cli.saveCloseAddTask); err != nil {
		return err
	}

	//for _, n := range []string{"but1", "but2"} {
	//	if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
	//		return err
	//	}
	//}
	//if err := g.SetKeybinding("msg", gocui.MouseLeft, gocui.ModNone, delMsg); err != nil {
	//	return err
	//}
	return nil
}
