package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/simonprast/buffalo_entrypoint/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
