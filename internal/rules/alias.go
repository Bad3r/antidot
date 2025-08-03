package rules

import (
	"github.com/bad3r/antidot-home/internal/tui"
)

type Alias struct {
	Alias   string
	Command string
}

// TODO: remove code duplication with export.go
func (a Alias) Apply(actx *ActionContext) error {
	err := actx.KeyValueStore.AddAlias(a.Alias, a.Command)
	if err != nil {
		return err
	}

	return nil
}

func (a Alias) Pprint() {
	tui.Print(
		"  %s %s%s\"%s\"",
		tui.ApplyStyle(tui.Magenta, "ALIAS"),
		a.Alias,
		tui.ApplyStyle(tui.Gray, "="),
		a.Command,
	)
}

func init() {
	registerAction("alias", Alias{})
}
