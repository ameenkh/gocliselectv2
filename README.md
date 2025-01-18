# Golang CLI Select V2
Lightweight interactive CLI selection library 

![](https://media.giphy.com/media/Nmc3muJhaCfPe2LWd9/giphy.gif)


## Import the package
```go
import "github.com/ameenkh/gocliselectv2"
```

## Usage
Create a new menu, supplying the question as a parameter

```go
menu := gocliselect.NewMenu("Chose a colour")
```

Add any number of options by calling `AddItem()` supplying the display text of the option
as well as the id
```go
menu.AddItem("Red", "red")
menu.AddItem("Blue", "blue")
menu.AddItem("Green", "green")
menu.AddItem("Yellow", "yellow")
menu.AddItem("Cyan", "cyan")
```

To display the menu and away the user choice call `Display()`

```go
choice := menu.Display()
```

## Example
```go
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
		fmt.Printf("interrupted...")
	} else {
		fmt.Printf("Choice: %s\n", choice)
	}
}

```
