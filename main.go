package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Groupe struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func GetData() (reponse []Groupe) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &reponse)
	if err != nil {
		panic(err)
	}

	return reponse
}

func container_artist(s Groupe, img *canvas.Image) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString("https://upload.wikimedia.org/wikipedia/commons/thumb/f/fa/Apple_logo_black.svg/170px-Apple_logo_black.svg.png")
	logo := canvas.NewImageFromResource(r)
	logo.FillMode = canvas.ImageFillOriginal
	btn := widget.NewButton("", func() {
	})

	containers := container.New(
		layout.NewHBoxLayout(),
		container.New(layout.NewVBoxLayout(), logo, btn),
		container.New(layout.NewCenterLayout(),
			container.New(layout.NewVBoxLayout(), widget.NewLabel(s.Name)),
		))
	return containers
}

func img_button(s Groupe, a fyne.App, myWindow fyne.Window) *fyne.Container { // return type
	//option button
	btn := widget.NewButton("", func() {
	})
	//img button
	r, _ := fyne.LoadResourceFromURLString(s.Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal

	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		widget.NewCard(s.Name, "", img),
		// 2nd btn widget
		btn,
	)
	az := container_artist(s, img)
	btn = widget.NewButton("", func() {
		myWindow.SetContent(az)
	})
	container1.Add(btn)
	container1.Resize(fyne.NewSize(100, 100))
	container1.Refresh()
	// our button is ready
	return container1
}

func tab_to_string(s []string) string {
	var tmp string

	for i := range s {
		tmp = fmt.Sprintf("%s\n%s", tmp, s[i])
	}
	return tmp
}

func Bar() fyne.CanvasObject {
	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Search", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("submit")
		},
	}

	artist := widget.NewButton("Artist", func() {
		canvas.NewText("Artist", color.Black)
	})
	local := widget.NewButton("Localisation", func() {
		canvas.NewText("Localisation", color.Black)
	})
	geo := widget.NewButton("Géolocalisation", func() {
		canvas.NewText("Géolocalisation", color.Black)
	})

	tmp := container.NewGridWithColumns(2,
		artist,
		local,
		geo,
		form,
	)
	return tmp
}

func main() {
	myApp := app.New()
	tab := GetData()
	myWindow := myApp.NewWindow("Grid Wrap Layout")

	myWindow.Resize(fyne.NewSize(200, 200))

	grid := container.NewAdaptiveGrid(4)

	for i := range tab {
		img := img_button(tab[i], myApp, myWindow)
		grid.Add(img)
	}

	label := canvas.NewText("Groupi Tracker", color.Black)
	label.TextSize = 50

	r, _ := fyne.LoadResourceFromURLString("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTQ1CxQj0OZlWftrFpRAs9LiJGL281KBDlMwzlmQ4Q&s")
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillContain
	img.SetMinSize(fyne.Size{Width: 30, Height: 20})

	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	Box := container.NewVBox(
		container.NewGridWithRows(1,
			container.New(
				layout.NewCenterLayout(),
				label,
			),
		),
		container.NewGridWithRows(2,
			img,
		),
		Bar(),
		container.NewGridWithRows(2,
			img,
		),
		container.NewGridWithColumns(2,
			container.NewGridWrap(
				fyne.NewSize(1700, 685),
				scroll,
			),
		),
	)

	myWindow.SetContent(Box)
	// myWindow.Resize(fyne.NewSize(180, 75))*/
	myWindow.ShowAndRun()
}
