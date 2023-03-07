package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	fonk "groupie-tracker/Fonction"
	lib "groupie-tracker/Lib"
	requestapi "groupie-tracker/RequestAPI"
)

func main() {
	myApp := app.New()
	tab := requestapi.GetData("https://groupietrackers.herokuapp.com/api/artists")
	w := myApp.NewWindow("Grid Wrap Layout")

	grid := container.NewAdaptiveGrid(4)
	for i := range tab {
		img := lib.Img_button(tab[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	w.SetContent(container.NewBorder(fonk.Bar(scroll, w, tab), nil, nil, nil, fonk.Menu(scroll)))
	w.ShowAndRun()
}
