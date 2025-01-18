package gocliselect

import (
	"fmt"
	"github.com/buger/goterm"
	"github.com/pkg/term"
	"io"
	"log"
)

type Menu struct {
	Prompt          string
	CursorPos       int
	MenuItems       []*MenuItem
	ShortcutKeysMap map[KeyCode]int
	selectedColor   SelectedColor
	pageSize        int
	printShortcuts  bool
}

type MenuOption func(menu *Menu)

func WithSelectedColor(color SelectedColor) MenuOption {
	return func(menu *Menu) {
		menu.selectedColor = color
	}
}

func WithPageSize(size int) MenuOption {
	return func(menu *Menu) {
		menu.pageSize = size
	}
}

func WithPrintShortcuts(print bool) MenuOption {
	return func(menu *Menu) {
		menu.printShortcuts = print
	}
}

type MenuItem struct {
	Text        string
	ID          string
	ShortcutKey KeyCode
	//SubMenu *Menu
}

func NewMenu(prompt string, opts ...MenuOption) *Menu {
	menu := &Menu{
		Prompt:          prompt,
		MenuItems:       make([]*MenuItem, 0),
		ShortcutKeysMap: make(map[KeyCode]int),
		selectedColor:   YELLOW,
		pageSize:        3,
		printShortcuts:  true,
	}
	for _, opt := range opts {
		opt(menu)
	}
	return menu
}

// AddItem will add a new menu option to the menu list
func (m *Menu) AddItem(option string, id string) *Menu {
	menuItem := &MenuItem{
		Text:        option,
		ID:          id,
		ShortcutKey: unknownKey,
	}

	m.MenuItems = append(m.MenuItems, menuItem)
	return m
}

// AddItemWithShortcutKey will add a new menu option with shortcut key to the menu list
func (m *Menu) AddItemWithShortcutKey(option string, id string, shortcutKey KeyCode) *Menu {
	menuItem := &MenuItem{
		Text:        option,
		ID:          id,
		ShortcutKey: shortcutKey,
	}
	m.MenuItems = append(m.MenuItems, menuItem)

	if idx, found := m.ShortcutKeysMap[shortcutKey]; found {
		//If shortcut already exists, then override it
		m.MenuItems[idx].ShortcutKey = unknownKey
	}
	m.ShortcutKeysMap[shortcutKey] = len(m.MenuItems) - 1

	return m
}

// renderMenuItems prints the menu item list.
// Setting redraw to true will re-render the options list with updated current selection.
func (m *Menu) renderMenuItems(redraw bool) {
	if redraw {
		// Move the cursor up n lines where n is the number of options, setting the new
		// location to start printing from, effectively redrawing the option list
		//
		// This is done by sending a VT100 escape code to the terminal
		// @see http://www.climagic.org/mirrors/VT100_Escape_Codes.html
		fmt.Printf("\033[%dA", len(m.MenuItems)-1)
	}

	for index, menuItem := range m.MenuItems {
		var newline = "\n"
		if index == len(m.MenuItems)-1 {
			// Adding a new line on the last option will move the cursor position out of range
			// For out redrawing
			newline = ""
		}

		menuItemText := menuItem.Text
		cursor := "  "
		if index == m.CursorPos {
			cursor = goterm.Color("> ", int(m.selectedColor))
			menuItemText = goterm.Color(menuItemText, int(m.selectedColor))
		}

		shortcutKeyText := ""
		if m.printShortcuts && menuItem.ShortcutKey != unknownKey {
			shortcutKeyText = fmt.Sprintf("[%c] ", keyToInputMap1[menuItem.ShortcutKey])
		}

		fmt.Printf("\r%s %s%s%s", cursor, shortcutKeyText, menuItemText, newline)
	}
}

// Display will display the current menu options and awaits user selection
// It returns the users selected choice
func (m *Menu) Display() (string, error) {
	defer func() {
		// Show cursor again.
		fmt.Printf("\033[?25h")
	}()

	fmt.Printf("%s\n", goterm.Color(goterm.Bold(m.Prompt)+":", goterm.CYAN))

	m.renderMenuItems(false)

	// Turn the terminal cursor off
	fmt.Printf("\033[?25l")

	for {
		keyCode := getInput()
		switch keyCode {
		case escKey:
			return "", nil
		case ctrlCKey, ctrlDKey:
			return "", io.EOF
		case enterKey:
			menuItem := m.MenuItems[m.CursorPos]
			fmt.Println("\r")
			return menuItem.ID, nil
		case upKey, shiftTabKey:
			m.CursorPos = (m.CursorPos + len(m.MenuItems) - 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		case downKey, tabKey:
			m.CursorPos = (m.CursorPos + 1) % len(m.MenuItems)
			m.renderMenuItems(true)
		case homeKey:
			m.CursorPos = 0
			m.renderMenuItems(true)
		case endKey:
			m.CursorPos = len(m.MenuItems) - 1
			m.renderMenuItems(true)
		case pageDownKey:
			m.CursorPos = min(m.CursorPos+m.pageSize, len(m.MenuItems)-1)
			m.renderMenuItems(true)
		case pageUpKey:
			m.CursorPos = max(m.CursorPos-m.pageSize, 0)
			m.renderMenuItems(true)
		default:
			//Check for shortcuts
			if idx, found := m.ShortcutKeysMap[keyCode]; found {
				menuItem := m.MenuItems[idx]
				fmt.Println("\r")
				return menuItem.ID, nil
			}
		}
	}
}

// getInput will read raw input from the terminal
// It will map the raw input to KeyCode and return it
func getInput() KeyCode {
	t, _ := term.Open("/dev/tty")

	err := term.RawMode(t)
	if err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)

	t.Restore()
	t.Close()
	// Arrow keys are prefixed with the ANSI escape code which take up the first two bytes.
	// The third byte is the key specific value we are looking for.
	// For example the left arrow key is '<esc>[A' while the right is '<esc>[C'
	// See: https://en.wikipedia.org/wiki/ANSI_escape_code
	switch read {
	case 1:
		return inputToKeyMap1[readBytes[0]]
	case 3:
		if readBytes[0] == byte(27) && readBytes[1] == byte(91) {
			return inputToKeyMap3[readBytes[2]]
		}
	}
	return unknownKey
}
