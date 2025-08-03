package rules

import (
	"os"

	"github.com/bad3r/antidot-home/internal/tui"
	"github.com/bad3r/antidot-home/internal/utils"
)

type Delete struct {
	Path string
}

func (d Delete) Apply(actx *ActionContext) error {
	path := utils.ExpandEnv(d.Path)
	if !utils.FileExists(path) {
		return nil
	}

	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (d Delete) Pprint() {
	tui.Print("  %s %s", tui.ApplyStyle(tui.Red, "DELETE"), utils.ExpandEnv(d.Path))
}

func init() {
	registerAction("delete", Delete{})
}
