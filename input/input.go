package input

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/golangast/sugargen/colors"
	"github.com/golangast/sugargen/loggers"
	"github.com/golangast/sugargen/printer"
)

func InputScanDirections(directions string) string {
	fmt.Println(directions)

	scannerdesc := bufio.NewScanner(os.Stdin)
	tr := scannerdesc.Scan()
	if tr {
		dir := scannerdesc.Text()
		stripdir := strings.TrimSpace(dir)
		return stripdir
	} else {
		return ""
	}

}

func MenuInstuctions(list []string, cols int, background, foreground, instructions string) string {
	logger := loggers.CreateLogger()

	var (
		atline int
		chosen []string
		ans    string
	)

	fmt.Print("\033[H\033[2J")
	colors.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")
	colors.ColorizeCol("purple", "purple", instructions)
	fmt.Println("\n")
	printer.PrintColumns(cols, atline, list, chosen, background, foreground)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		//press arrows to change index to highlight selected item
		switch key.String() {
		case "up": //up arrow
			atlines, run, err := printer.UP(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "down": //down arrow
			atlines, run, err := printer.Down(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "right": //left arrow
			atlines, run, err := printer.Right(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "left": //left arrow
			atlines, run, err := printer.Left(atline, cols, background, foreground, list, chosen)
			atline = atlines
			return run, err
		case "enter": //enter
			ans = list[atline]
			return true, nil
		case "q", "esc", "c", "ctrl+c": //to quit
			return true, nil
		default:
			fmt.Println(key.String())
			return false, nil // Return false to continue listening
		}
	})

	if err != nil {
		logger.Error(
			"pressing keys",
			slog.String("error: ", err.Error()),
		)
	}
	return ans

}
