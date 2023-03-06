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
	"strings"

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
	DatesLocations map[string][]string `json:"datesLocations"`
	DatesLocation  string
}

type RI struct {
	Index []R `json:"index"`
}

type L struct {
	Locations []string `json:"locations"`
	Location  string
}

type LI struct {
	Index []L `json:"index"`
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

func GetR(url string) (reponse RI) {
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&reponse)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return reponse
}

func GetL(url string) (reponse LI) {
	data, err := http.Get(url)
	if err != nil {
		log.Fatal("Error when opening url: ", err)
	}
	defer data.Body.Close()
	err = json.NewDecoder(data.Body).Decode(&reponse)
	if err != nil {
		log.Fatal("Error during Decode: ", err)
	}
	return reponse
}

func A(id int) string {
	var c int
	var listLocations []string
	li := GetL("https://groupietrackers.herokuapp.com/api/locations")
	ri := GetR("https://groupietrackers.herokuapp.com/api/relation")

	var r R
	contentRelations := ""

	r.DatesLocation = "Relations : "
	for _, i := range li.Index[id].Locations {
		for c = 0; c <= len(ri.Index[id].DatesLocations[i])-1; c++ {
			var isDouble bool = false
			for _, q := range listLocations {
				if q == i {
					isDouble = true
				}
			}
			if !isDouble {
				r.DatesLocation += i + " : " + ri.Index[id].DatesLocations[i][c] + "  "
				listLocations = append(listLocations, i)
			} else {
				r.DatesLocation += ", " + ri.Index[id].DatesLocations[i][c] + " "
			}
			contentRelations = r.DatesLocation
		}
	}
	return contentRelations
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

func SearchBar(s string, tab []Groupe, scroll *container.Scroll, w fyne.Window) *fyne.Container {
	n := 0
	for i := range tab {
		if strings.Compare(s, tab[i].Name) == 0 {
			n++
		}
		if strings.Contains(tab[i].Name, s) == true {
			n++
		}
	}

	var search []Groupe
	v := 0

	for i := range tab {
		if strings.Compare(s, tab[i].Name) == 0 {
			search = append(search, tab[i])
		}
		if strings.Contains(tab[i].Name, s) == true {
			for j := range search {
				if strings.Compare(search[j].Name, tab[i].Name) == 0 {
					v = 1
				}
			}
			if v == 0 {
				fmt.Println(tab[i].Name)
				search = append(search, tab[i])
			} else {
				v = 0
			}
		}
	}

	for i := range search {
		log.Println(search[i].Name)
	}

	grid := container.NewAdaptiveGrid(4)
	for i := range search {
		img := Art_mod(search[i], w, scroll, search)
		grid.Add(img)
	}
	scrol := container.NewHScroll(grid)
	scrol.Direction = container.ScrollBoth

	search = nil

	return container.NewGridWithColumns(1, scrol)
}

func Bar(scroll *container.Scroll, w fyne.Window, tab []Groupe) *fyne.Container {
	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Search", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, SearchBar(entry.Text, tab, scroll, w)))
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
	tmp := container.NewGridWithColumns(5,
		home,
		artist,
		local,
		geo,
		form,
	)

	return tmp
}

func FormatString(s string) string {
	if len(s)%20 != 0 {
		n := 0
		tmp := make([]rune, len(s)+len(s)%20)
		for i := range s {
			if n == 20 {
				tmp[i] = '\n'
				n = 0
			} else {
				tmp[i] = rune(s[i])
				n++
			}
		}
		fmt.Println(string(tmp))
		return string(tmp)
	}
	return s
}

func Desc_art(s Groupe) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(s.Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal

	label := canvas.NewText(s.Name, color.Black)
	label.TextSize = 50

	sublabel := canvas.NewText(fmt.Sprintf("%s : %s", "Menbres", tab_to_string(s.Members)), color.Black)
	sublabel.TextSize = 20

	sublabel1 := canvas.NewText(fmt.Sprintf("%s : %s", "First Album", s.FirstAlbum), color.Black)
	sublabel1.TextSize = 30

	str := FormatString(A(s.ID))

	sublabel3 := canvas.NewText(fmt.Sprintf("%s : %s", "", str), color.Black)
	sublabel3.TextSize = 15

	sublabel4 := canvas.NewText(fmt.Sprintf("%s : %d", "Creation Date", s.CreationDate), color.Black)
	sublabel4.TextSize = 30

	spacer := layout.NewSpacer()

	containers := container.NewVBox(
		container.NewGridWithColumns(1,
			label,
		),
		container.NewGridWithColumns(1,
			img,
		),
		spacer,
		container.NewGridWithColumns(1,
			sublabel,
		),
		container.NewGridWithColumns(2,
			sublabel1,
			sublabel4,
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(1,
			sublabel3,
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
