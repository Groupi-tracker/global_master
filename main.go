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

type R struct {
	DatesLocations string `json:"datesLocations"`
}

func GetData(url string) (reponse []Groupe) {
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

func GetDataR(url string) (reponse []R) {
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

func img_button(s Groupe) *fyne.Container { // return type
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
	)
	// our button is ready
	return container1
}

func tab_to_string(s []string) string {
	var tmp string

	for i := range s {
		tmp = fmt.Sprintf("%s %s", tmp, s[i])
	}
	return tmp
}

func Bar(scroll *container.Scroll, w fyne.Window, tab []Groupe) *fyne.Container {
	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Search", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			log.Println("submit")
		},
	}

	artist := widget.NewButton("Artist", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Artist(w, tab, scroll)))
	})
	local := widget.NewButton("Localisation", func() {
		canvas.NewText("Localisation", color.Black)
	})
	geo := widget.NewButton("Géolocalisation", func() {
		canvas.NewText("Géolocalisation", color.Black)
	})
	home := widget.NewButton("home", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Menu(scroll)))
	})
	tmp := container.NewGridWithColumns(4,
		home,
		artist,
		local,
		geo,
		form,
	)

	return tmp
}

func Desc_art(s Groupe) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(s.Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal

	label := canvas.NewText(s.Name, color.Black)
	label.TextSize = 50

	sublabel := canvas.NewText(fmt.Sprintf("%s : %s", "Menbres", tab_to_string(s.Members)), color.Black)
	sublabel.TextSize = 30

	sublabel1 := canvas.NewText(fmt.Sprintf("%s : %s", "First Album", s.FirstAlbum), color.Black)
	sublabel1.TextSize = 30

	sublabel2 := canvas.NewText(fmt.Sprintf("%s : %s", "Relation", s.Relations), color.Black)
	sublabel2.TextSize = 30

	sublabel3 := canvas.NewText(fmt.Sprintf("%s : %s", "Concert Dates", s.ConcertDates), color.Black)
	sublabel3.TextSize = 30

	sublabel4 := canvas.NewText(fmt.Sprintf("%s : %s", "Creation Date", s.CreationDate), color.Black)
	sublabel4.TextSize = 30

	containers := container.NewVBox(
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				label,
			),
		),
		container.NewGridWithColumns(1,
			img,
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				sublabel,
			),
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				sublabel1,
			),
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				sublabel2,
			),
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				sublabel3,
			),
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			container.New(
				layout.NewCenterLayout(),
				sublabel4,
			),
		),
	)
	return containers
}

func Art_mod(s Groupe, w fyne.Window, scroll *container.Scroll, tab []Groupe) *fyne.Container {
	btn := widget.NewButton("", func() {
	})

	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		widget.NewCard(s.Name, "", nil),
		// 2nd btn widget
		btn,
	)
	btn = widget.NewButton("", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Desc_art(s)))
	})
	container1.Add(btn)
	return container1
}

func Artist(w fyne.Window, tab []Groupe, scroll *container.Scroll) *fyne.Container {
	grid := container.NewAdaptiveGrid(4)

	for i := range tab {
		img := Art_mod(tab[i], w, scroll, tab)
		grid.Add(img)
	}
	scrol := container.NewHScroll(grid)
	scrol.Direction = container.ScrollBoth

	return container.NewGridWithColumns(1, scrol)
}

func Menu(scroll *container.Scroll) *fyne.Container {
	label := canvas.NewText("Groupi Tracker", color.Black)
	label.TextSize = 50

	r, _ := fyne.LoadResourceFromURLString("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTQ1CxQj0OZlWftrFpRAs9LiJGL281KBDlMwzlmQ4Q&s")
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillContain

	tmp := container.NewVBox(
		container.NewGridWithRows(1,
			container.New(
				layout.NewCenterLayout(),
				label,
			),
		),
		container.NewGridWithRows(1,
			img,
		),
		container.NewGridWithColumns(1,
			container.NewGridWrap(
				fyne.NewSize(1700, 685),
				scroll,
			),
		),
	)
	return tmp
}

func main() {
	myApp := app.New()
	tab := GetData("https://groupietrackers.herokuapp.com/api/artists")
	w := myApp.NewWindow("Grid Wrap Layout")

	grid := container.NewAdaptiveGrid(4)
	for i := range tab {
		img := img_button(tab[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Menu(scroll)))
	w.ShowAndRun()
}
