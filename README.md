# Golang CLI Select V2
Lightweight interactive CLI selection library 

![](https://media.giphy.com/media/Nmc3muJhaCfPe2LWd9/giphy.gif)


## Menu Navigation
| Key                    | Action                                                   |
|------------------------|----------------------------------------------------------|
| up arrow / shift + tab | Moves the cursor up                                      |
| down arrow / tab       | Moves the cursor down                                    |
| home                   | Moves the cursor to the first item                       |
| end                    | Moves the cursor to the last item                        |
| page up                | Moves the cursor up by `pageSize` steps (configurable)   |
| page down              | Moves the cursor down by `pageSize` steps (configurable) |
| enter                  | Returns the selected item                                |
| escape                 | Returns an empty string without error                    |
| ctrl+c                 | Returns an empty string with `io.EOF` error              |
| ctrl+d                 | Returns an empty string with `io.EOF` error              |

## Menu Options
| Field          | Description                                                                                                  | Method             |
|----------------|--------------------------------------------------------------------------------------------------------------|--------------------|
| selectedColor  | Defines the selected item text color.  Default value is `SelectedColor.YELLOW`                               | WithSelectedColor  |
| pageSize       | Defines the number of steps to move the cursor when clicking page up/down.  Default value is `3`             | WithPageSize       |
| printShortcuts | If false, then the item text won't be printed with the shortcut key prefix `[key]`.  Default value is `true` | WithPrintShortcuts |

## Install the package
```shell
go get github.com/ameenkh/gocliselectv2
```

## Import the package
```go
import "github.com/ameenkh/gocliselectv2"
```

## Usage
Create a new menu, supplying the question as a parameter

```go
menu := gocliselect.NewMenu("Choose an operation")
```
You can change the menu options:
```go
menu := gocliselectv2.NewMenu("Choose an operation", gocliselectv2.WithSelectedColor(gocliselectv2.RED), gocliselectv2.WithPageSize(3))
```


Add any number of options by calling `AddItem()` supplying the display text of the option
as well as the id.

You can add also a shortcut key to any item (available shortcuts are letters and numbers only: [a-z] or [A-Z] or [0-9]):
```go
menu.AddItemWithShortcutKey("List", "list", gocliselectv2.Key_l)
menu.AddItemWithShortcutKey("Get", "get", gocliselectv2.Key_g)
menu.AddItem("Post", "post")
menu.AddItem("Delete", "delete")
```

To display the menu and away the user choice call `Display()`

```go
choice, err := menu.Display()
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

```
