package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	fonk "groupie-tracker/Fonction"
	lib "groupie-tracker/Lib"
	requestapi "groupie-tracker/RequestAPI"
)

func main() {
	myApp := app.New()
	tab := requestapi.GetData("https://groupietrackers.herokuapp.com/api/artists")
	w := myApp.NewWindow("Groupi")

	res, err := fyne.LoadResourceFromPath("./icon.png")
	if err != nil {
		fmt.Println(err)
	}

	w.SetIcon(res)

	grid := container.NewAdaptiveGrid(4)
	for i := range tab {
		img := lib.Img_button(tab[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	w.SetContent(container.NewBorder(fonk.Bar(scroll, w, tab), fonk.HubTri(scroll, w, tab), nil, nil, fonk.Menu(scroll)))
	w.ShowAndRun()
}
