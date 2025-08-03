package rules

import (
	"os"
	"path/filepath"

	"github.com/bad3r/antidot-home/internal/tui"
	"github.com/bad3r/antidot-home/internal/utils"
)

type Migrate struct {
	Source  string
	Dest    string
	Symlink bool
}

func (m Migrate) Apply(actx *ActionContext) error {
	source := utils.ExpandEnv(m.Source)
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		tui.Print("File %s doesn't exist. Skipping action", source)
		return nil
	} else if err != nil {
		return err
	}

	dest := utils.ExpandEnv(m.Dest)

	// NB: No error is thrown if dir already exists
	err = os.MkdirAll(filepath.Dir(dest), os.FileMode(0o755))
	if err != nil {
		return err
	}

	err = utils.MovePath(source, dest)
	if err != nil {
		return err
	}

	if m.Symlink {
		err = os.Symlink(source, dest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m Migrate) Pprint() {
	symlink := ""
	if m.Symlink {
		symlink = " (with symlink)"
	}

	// TODO: move the indentation logic elsewhere
	tui.Print(
		"  %s %s %s %s%s",
		tui.ApplyStyle(tui.Green, "MOVE  "),
		utils.ExpandEnv(m.Source),
		tui.ApplyStyle(tui.Gray, "→"),
		utils.ExpandEnv(m.Dest),
		symlink)
}

func init() {
	registerAction("migrate", Migrate{})
}
