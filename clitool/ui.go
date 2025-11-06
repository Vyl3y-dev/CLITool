package clitool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func LaunchUI(tasks []Task) {
	a := app.New()
	w := a.NewWindow("Task Manager")

	// Create a list widget
	taskList := widget.NewList(
		func() int {
			return len(tasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("") // template for each row
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tasks[i].Description)
		},
	)

	addEntry := widget.NewEntry()
	addEntry.SetPlaceHolder("New task...")

	addButton := widget.NewButton("Add", func() {
		desc := addEntry.Text
		if desc != "" {
			tasks = AddTask(tasks, desc)
			SaveTasks(tasks)
			taskList.Refresh()
			addEntry.SetText("")
		}
	})

	w.SetContent(container.NewVBox(
		taskList,
		container.NewHBox(addEntry, addButton),
	))

	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}
