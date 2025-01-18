package main

import (
	"errors"
	"fmt"
	"github.com/ameenkh/gocliselectv2"
	"io"
)

func main() {
	menu := gocliselectv2.NewMenu("Choose an operation", gocliselectv2.WithSelectedColor(gocliselectv2.RED), gocliselectv2.WithPageSize(3))

	menu.AddItemWithShortcutKey("List", "list", gocliselectv2.Key_l)
	menu.AddItemWithShortcutKey("Get", "get", gocliselectv2.Key_g)
	menu.AddItem("Post", "post")
	menu.AddItem("Delete", "delete")

	choice, err := menu.Display()
	if errors.Is(err, io.EOF) {
		fmt.Printf("interrupted...") // in case of ctrl+c or ctrl+d
	} else {
		fmt.Printf("Choice: %s\n", choice)
	}
}
