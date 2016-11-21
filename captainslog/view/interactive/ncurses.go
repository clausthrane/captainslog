package interactive

import (
	"github.com/jroimartin/gocui"
	"log"
	"os"
	"github.com/clausthrane/captainslog/captainslog"
)

var logger = log.New(os.Stdout, " ", log.Ldate | log.Ltime | log.Lshortfile)

type Cli struct {
	app *captainslog.CaptainsLog
	gui *gocui.Gui
}

func NewInteractiveCli(captainslog *captainslog.CaptainsLog) *Cli {

	g, err := gocui.NewGui()
	cli := &Cli{captainslog, g}

	if err != nil {
		logger.Panicln(err)
	}

	g.Cursor = true
	g.Mouse = true

	defer g.Close()

	g.SetManagerFunc(cli.setupLayout())

	if err := cli.keybindings(g); err != nil {
		logger.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		logger.Panicln(err)
	}

	return cli
}

func (cli *Cli) setupLayout() func(g *gocui.Gui) error {

	return func(g *gocui.Gui) error {
		e := cli.setupMenu()
		e = cli.setupGroups()
		return e
	}

}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
//
//func showMsg(g *gocui.Gui, v *gocui.View) error {
//	var l string
//	var err error
//
//	if _, err := g.SetCurrentView(v.Name()); err != nil {
//		return err
//	}
//
//	_, cy := v.Cursor()
//	if l, err = v.Line(cy); err != nil {
//		l = ""
//	}
//
//	maxX, maxY := g.Size()
//	if v, err := g.SetView("msg", maxX / 2 - 10, maxY / 2, maxX / 2 + 10, maxY / 2 + 2); err != nil {
//		if err != gocui.ErrUnknownView {
//			return err
//		}
//		fmt.Fprintln(v, l)
//	}
//	return nil
//}
//
//func delMsg(g *gocui.Gui, v *gocui.View) error {
//	if err := g.DeleteView("msg"); err != nil {
//		return err
//	}
//	return nil
//}
