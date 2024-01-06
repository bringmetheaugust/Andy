package components

import (
	"time"

	"github.com/briandowns/spinner"
)

var Spinner *spinner.Spinner

func init() {
	Spinner = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
}
