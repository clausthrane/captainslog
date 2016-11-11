package cli

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	for _, n := range []string{"but1", "but2"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}
	if err := g.SetKeybinding("msg", gocui.MouseLeft, gocui.ModNone, delMsg); err != nil {
		return err
	}
	return nil
}

