package tools

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
)

func NewProgressBar(max int) *progressbar.ProgressBar {

	if max == -1 {
		return progressbar.NewOptions(-1, progressbar.OptionShowDescriptionAtLineEnd(), progressbar.OptionEnableColorCodes(true), progressbar.OptionClearOnFinish())
	} else {
		return progressbar.NewOptions(max, progressbar.OptionShowDescriptionAtLineEnd(), progressbar.OptionEnableColorCodes(true), progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}), progressbar.OptionOnCompletion(func() {
			fmt.Println()
		}))
	}
}
