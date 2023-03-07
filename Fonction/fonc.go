package fonk

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	fynex "fyne.io/x/fyne/widget"
	lib "groupie-tracker/Lib"
	requestapi "groupie-tracker/RequestAPI"
	"image/color"
	"net/http"
	"strings"
)

func SearchBar(s string, tab []requestapi.Groupe, scroll *container.Scroll, w fyne.Window) *fyne.Container {
	n := 0
	for i := range tab {
		if strings.Compare(s, tab[i].Name) == 0 {
			n++
		}
		if strings.Contains(tab[i].Name, s) == true {
			n++
		}
	}

	var search []requestapi.Groupe
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

	grid := container.NewAdaptiveGrid(4)
	for i := range search {
		img := Art_mod(search[i], w, scroll, search)
		grid.Add(img)
	}
	scrol := container.NewHScroll(grid)
	scrol.Direction = container.ScrollBoth

	return container.NewGridWithColumns(1, scrol)
}

func Bar(scroll *container.Scroll, w fyne.Window, tab []requestapi.Groupe) *fyne.Container {
	entry := fynex.NewCompletionEntry([]string{})

	// When the use typed text, complete the list.
	entry.OnChanged = func(s string) {
		// completion start for text length >= 3
		if len(s) < 3 {
			entry.HideCompletion()
			return
		}

		// Make a search on wikipedia
		resp, err := http.Get(
			"https://en.wikipedia.org/w/api.php?action=opensearch&search=" + entry.Text,
		)
		if err != nil {
			entry.HideCompletion()
			return
		}

		// Get the list of possible completion
		var results [][]string
		json.NewDecoder(resp.Body).Decode(&results)

		// no results
		if len(results) == 0 {
			entry.HideCompletion()
			return
		}

		// then show them
		entry.SetOptions(results[1])
		entry.ShowCompletion()
	}

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Search", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			if entry.Text == "" {
				w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Artist(w, tab, scroll)))
			}
			w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, SearchBar(entry.Text, tab, scroll, w)))
		},
	}

	artist := widget.NewButton("Artist", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Artist(w, tab, scroll)))
	})
	geo := widget.NewButton("Géolocalisation", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Geo(w, tab, scroll)))
	})
	home := widget.NewButton("home", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(scroll)))
	})
	tmp := container.NewGridWithColumns(4,
		home,
		artist,
		geo,
		form,
	)

	return tmp
}

func Desc_art(s requestapi.Groupe) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(s.Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal

	label := canvas.NewText(s.Name, color.Black)
	label.TextSize = 50

	sublabel := canvas.NewText(fmt.Sprintf("%s : %s", "Menbres", lib.Tab_to_string(s.Members)), color.Black)
	sublabel.TextSize = 20

	sublabel1 := canvas.NewText(fmt.Sprintf("%s : %s", "First Album", s.FirstAlbum), color.Black)
	sublabel1.TextSize = 20

	sep := canvas.NewText("Relation :", color.Black)
	sep.TextSize = 20

	maps, date := lib.FormatString(requestapi.A(s.ID))

	sublabel4 := canvas.NewText(fmt.Sprintf("%s : %d", "Creation Date", s.CreationDate), color.Black)
	sublabel4.TextSize = 20

	containers := container.New(layout.NewVBoxLayout(),
		container.NewGridWithColumns(1,
			label,
		),
		container.NewGridWithColumns(1,
			img,
		),
		container.NewGridWithColumns(1,
			sublabel,
		),
		canvas.NewText("------------------------------------------------------", color.Black),
		container.NewGridWithColumns(1,
			sublabel1,
		),
		canvas.NewText("------------------------------------------------------", color.Black),
		container.NewGridWithColumns(1,
			sublabel4,
		),
		canvas.NewText("------------------------------------------------------", color.Black),
		container.NewGridWithColumns(1,
			sep,
		),
	)
	a := 0
	if len(maps) > len(date) {
		a = len(date)
	} else {
		a = len(maps)
	}
	for i := 0; i < a; i++ {
		m := canvas.NewText(fmt.Sprintf("%s => %s", maps[i], date[i]), color.Black)
		m.TextSize = 15
		p := container.NewGridWithColumns(1,
			m,
		)
		containers.Add(p)
	}
	return containers
}

func Art_mod(s requestapi.Groupe, w fyne.Window, scroll *container.Scroll, tab []requestapi.Groupe) *fyne.Container {
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

func Artist(w fyne.Window, tab []requestapi.Groupe, scroll *container.Scroll) *fyne.Container {
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
	label := canvas.NewText("Groupie Tracker", color.Black)
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
