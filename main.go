package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/gofrs/uuid"
)

var wf *aw.Workflow

func run() {
	query := wf.Args()[0]

	v4, _ := uuid.NewV4()
	addOption(v4.String(), "v4")

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("Unknown format", "")
	wf.SendFeedback()
}

func addOption(value, formatName string) {
	wf.NewItem(formatName).
		Subtitle(value).
		Valid(true).
		Arg(value)
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
