package printer

import (
	"fmt"
	"slices"

	"github.com/golangast/sugargen/colors"
)

func ClearDirections() {
	fmt.Print("\033[H\033[2J")
	colors.ColorizeCol("purple", "purple", "(q-quit) - (c-multiselection) - (r-remove) - (enter-select/execute) - (u-update tag) - down/up/left/right")
	fmt.Println("\n")

}

func Directions() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("")
}

func PrintColumns(cols, atline int, list, chosen []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {

				colors.ColorizeCol(background, foreground, list[atline])

			} else {
				if slices.Contains(chosen, list[i]) {
					colors.ColorizeCol("purple", foreground, list[i])

				} else {
					fmt.Printf("%-11s%s", list[i], " ")
				}
			}
		}
		fmt.Println() //yes this needs to be here for padding
	}
}

func PrintColumnsWChosen(cols, atline int, list []string, background, foreground string) {
	rows := (len(list) + cols - 1) / cols

	for row := 0; row < rows; row++ {

		for col := 0; col < cols; col++ {
			i := col*rows + row
			if i >= len(list) {
				break // This means the last column is not "full"
			}

			if i == atline {
				colors.ColorizeCol(background, foreground, list[atline])

			} else {
				fmt.Printf("%-11s%s", list[i], " ")

			}

		}
		fmt.Println() //yes this needs to be here for padding

	}
}

func UP(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	if atline >= 1 {
		atline--
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Down(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	if atline <= linecount-2 {
		atline++
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil

}
func Right(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	linecount := len(list)
	rows := (len(list) + cols - 1) / cols
	if atline <= linecount-rows {
		atline = atline + rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)
	return atline, false, nil
}

func Left(atline, cols int, background, foreground string, list, chosen []string) (int, bool, error) {
	ClearDirections()

	rows := (len(list) + cols - 1) / cols
	if atline >= rows {
		atline = atline - rows
	}
	PrintColumns(cols, atline, list, chosen, background, foreground)

	return atline, false, nil
}
