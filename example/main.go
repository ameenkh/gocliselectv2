package main

import (
	"errors"
	"fmt"
	"github.com/ameenkh/gocliselectv2"
	"io"
)

func main() {
	menu := gocliselectv2.NewMenu("Chose a colour")

	menu.AddItemWithShortcutKey("Red", "red", gocliselectv2.Key_r)
	menu.AddItemWithShortcutKey("Blue", "blue", gocliselectv2.Key_b)
	menu.AddItem("Green", "green")
	menu.AddItem("Yellow", "yellow")
	menu.AddItem("Cyan", "cyan")

	choice, err := menu.Display()
	if errors.Is(err, io.EOF) {
		fmt.Printf("interrupted...") // in case of (ctrl + c/d)
	} else {
		fmt.Printf("Choice: %s\n", choice)
	}
}
