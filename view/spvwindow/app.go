package spvwindow

import "github.com/andlabs/ui"

func App() {
	input := ui.NewEntry()
	button := ui.NewButton("Greet")
	greeting := ui.NewLabel("128768")
	box := ui.NewVerticalBox()
	box.Append(ui.NewLabel("Enter your name:"), false)
	box.Append(input, false)
	box.Append(button, false)
	box.Append(greeting, false)
	window := ui.NewWindow("Hello", 800, 600, false)
	window.SetMargined(true)
	window.SetChild(box)
	button.OnClicked(func(*ui.Button) {
		greeting.SetText("Hello, " + input.Text() + "!")
	})
	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	window.Show()
}
