package fonk

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	lib "groupie-tracker/Lib"
	requestapi "groupie-tracker/RequestAPI"
	"image/color"
	"io/ioutil"
	"net/http"
	"net/url"
)

func loc(s string) (float64, float64) {
	// Ville dont on veut récupérer les coordonnées
	city := s

	// Construire l'URL de l'API de géocodage de Google Maps
	apiURL := "https://maps.googleapis.com/maps/api/geocode/json"
	params := url.Values{}
	params.Set("address", city)
	params.Set("key", "AIzaSyBH6ssvOYwwrir9IekBjNXksswDTwn5Wy0")
	fullURL := apiURL + "?" + params.Encode()

	// Envoyer une requête GET à l'API de géocodage
	resp, err := http.Get(fullURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lire la réponse de l'API de géocodage
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Analyser la réponse JSON de l'API de géocodage
	var data struct {
		Results []struct {
			Geometry struct {
				Location struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"location"`
			} `json:"geometry"`
		} `json:"results"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// Récupérer les coordonnées de la ville
	if len(data.Results) > 0 {
		lat := data.Results[0].Geometry.Location.Lat
		lng := data.Results[0].Geometry.Location.Lng
		return lat, lng
	} else {
		fmt.Printf("Impossible de récupérer les coordonnées de %s\n", city)
	}
	return 0, 0
}

func mapImg(o int, lat float64, lon float64) *canvas.Image {
	// URL de la carte avec les coordonnées
	mapURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/staticmap?center=%f,%f&zoom=%d&size=640x480&key=AIzaSyBH6ssvOYwwrir9IekBjNXksswDTwn5Wy0", lat, lon, o)

	// Créer un objet Image avec l'URL de la carte
	r, _ := fyne.LoadResourceFromURLString(mapURL)
	mapImage := canvas.NewImageFromResource(r)
	mapImage.FillMode = canvas.ImageFillOriginal
	return mapImage
}

func refresh(o int, lat float64, lon float64) *fyne.Container {
	return container.New(layout.NewCenterLayout(), mapImg(o, lat, lon))
}

func zoom(o int, lat float64, lon float64, w fyne.Window, scroll *container.Scroll, tab []requestapi.Groupe) *fyne.Container {

	plus := widget.NewButton("+", func() {
		o++
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, zoom(o, lat, lon, w, scroll, tab), nil, refresh(o, lat, lon)))
	})

	moin := widget.NewButton("-", func() {
		o--
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, zoom(o, lat, lon, w, scroll, tab), nil, refresh(o, lat, lon)))
	})

	return container.NewGridWithRows(2, plus, moin)
}

func Desc_geo(s requestapi.Groupe, w fyne.Window, scroll *container.Scroll, tab []requestapi.Groupe) *fyne.Container {
	r, _ := fyne.LoadResourceFromURLString(s.Image)
	img := canvas.NewImageFromResource(r)
	img.FillMode = canvas.ImageFillOriginal

	label := canvas.NewText(s.Name, color.Black)
	label.TextSize = 50

	sublabel := canvas.NewText("Localisation Des Concert : ", color.Black)
	sublabel.TextSize = 20

	maps, _ := lib.FormatStringGeo(requestapi.A(s.ID))

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
	)
	m := container.NewAdaptiveGrid(len(maps))
	for i := 0; i < len(maps); i++ {
		lat, lon := loc(maps[i])
		o := 12
		p := widget.NewButton(maps[i], func() {
			w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, zoom(o, lat, lon, w, scroll, tab), nil, refresh(o, lat, lon)))
		})
		m.Add(p)
	}
	containers.Add(m)
	return containers
}

func Geo_mod(s requestapi.Groupe, w fyne.Window, scroll *container.Scroll, tab []requestapi.Groupe) *fyne.Container {
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
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), nil, nil, nil, Desc_geo(s, w, scroll, tab)))
	})
	container1.Add(btn)
	return container1
}

func Geo(w fyne.Window, tab []requestapi.Groupe, scroll *container.Scroll) *fyne.Container {
	grid := container.NewAdaptiveGrid(4)

	for i := range tab {
		img := Geo_mod(tab[i], w, scroll, tab)
		grid.Add(img)
	}
	scrol := container.NewHScroll(grid)
	scrol.Direction = container.ScrollBoth

	return container.NewGridWithColumns(1, scrol)
}
