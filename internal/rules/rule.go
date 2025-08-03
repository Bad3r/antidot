package rules

import (
	"github.com/bad3r/antidot-home/internal/dotfile"
	"github.com/bad3r/antidot-home/internal/tui"
)

type Rule struct {
	Name        string
	Description string
	Dotfile     *dotfile.Dotfile
	Ignore      bool
	Actions     []Action
	Notes       []string
}

func (r Rule) Pprint() {
	tui.Print(tui.ApplyStylef(tui.Cyan, "Rule %s:", r.Name))
	if len(r.Notes) != 0 {
		for _, note := range r.Notes {
			tui.Print("  %s %s", tui.ApplyStyle(tui.Cyan, "NOTICE"), note)
		}
	}

	for _, action := range r.Actions {
		action.Pprint()
	}

	if r.Ignore {
		tui.Print(tui.ApplyStyle(tui.Gray, "  IGNORED"))
	}
}

// TODO: return the error
func (r Rule) Apply(actx *ActionContext) {
	if !r.Ignore {
		for _, action := range r.Actions {
			err := action.Apply(actx)
			if err != nil {
				tui.Warn("Failed to run rule %s: %v", r.Name, err)
				break
			}
		}
	}
}
