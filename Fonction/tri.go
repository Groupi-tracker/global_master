package fonk

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	lib "groupie-tracker/Lib"
	requestapi "groupie-tracker/RequestAPI"
	"sort"
	"time"
)

func HubTri(scroll *container.Scroll, w fyne.Window, tab []requestapi.Groupe) *fyne.Container {
	alpha := widget.NewButton("Nom Croissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TrieNomGroupe(w, tab))))
	})
	dealpha := widget.NewButton("Nom Decroissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TrieNomGroupeInverse(w, tab))))
	})
	tC := widget.NewButton("Groupe Croissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(FiltreNombreGroupeCroissant(w, tab))))
	})
	tD := widget.NewButton("Groupe Decroissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(FiltreNombreGroupeDecroissant(w, tab))))
	})
	dC := widget.NewButton("Date Croissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TriDateCroissant(w, tab))))
	})
	dD := widget.NewButton("Date Decroissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TriDateDeroissant(w, tab))))
	})
	aC := widget.NewButton("P.Album Croissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TriPremierAlbumCroissant(w, tab))))
	})
	aD := widget.NewButton("P.Album Decroissant", func() {
		w.SetContent(container.NewBorder(Bar(scroll, w, tab), HubTri(scroll, w, tab), nil, nil, Menu(TriPremierAlbumDecroissant(w, tab))))
	})
	// Création de la toolbar avec les éléments
	toolbar := container.NewGridWithColumns(4,
		container.NewGridWithRows(2,
			alpha,
			dealpha,
		),
		container.NewGridWithRows(2,
			tC,
			tD,
		),
		container.NewGridWithRows(2,
			dC,
			dD,
		),
		container.NewGridWithRows(2,
			aC,
			aD,
		),
	)
	return toolbar
}

//  Filtre Nombre de gens pour le groupe :

func FiltreNombreGroupesParTaille(w fyne.Window, minSize, maxSize int, exactSize int, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	var filteredTable []requestapi.Groupe
	for _, group := range copyTable {
		size := len(group.Members)
		if exactSize > 0 && size != exactSize {
			continue
		}
		if minSize > 0 && size < minSize {
			continue
		}
		if maxSize > 0 && size > maxSize {
			continue
		}
		filteredTable = append(filteredTable, group)
	}

	triParTaille := func(slice1, slice2 []string) bool { return len(slice1) < len(slice2) }
	sort.Slice(filteredTable, func(i, j int) bool { return triParTaille(filteredTable[i].Members, filteredTable[j].Members) })

	grid := container.NewAdaptiveGrid(4)
	for i := range filteredTable {
		img := lib.Img_button(filteredTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

func FiltreNombreGroupeCroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	triParTaille := func(slice1, slice2 []string) bool { return len(slice1) < len(slice2) }
	sort.Slice(copyTable, func(i, j int) bool { return triParTaille(copyTable[i].Members, copyTable[j].Members) })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

func FiltreNombreGroupeDecroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	triParTaille := func(slice1, slice2 []string) bool { return len(slice1) > len(slice2) }
	sort.Slice(copyTable, func(i, j int) bool { return triParTaille(copyTable[i].Members, copyTable[j].Members) })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

//  -------------------------------------------
// Filtre pour les dates de créations :

func FiltreGroupesDate(w fyne.Window, minDate, maxDate, exactDate int, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	var filteredTable []requestapi.Groupe
	for _, group := range copyTable {
		date := group.CreationDate
		if exactDate > 0 && date != exactDate {
			continue
		}
		if minDate > 0 && date < minDate {
			continue
		}
		if maxDate > 0 && date > maxDate {
			continue
		}
		filteredTable = append(filteredTable, group)
	}

	triParTaille := func(slice1, slice2 []string) bool { return len(slice1) < len(slice2) }
	sort.Slice(filteredTable, func(i, j int) bool { return triParTaille(filteredTable[i].Members, filteredTable[j].Members) })

	grid := container.NewAdaptiveGrid(4)
	for i := range filteredTable {
		img := lib.Img_button(filteredTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

func TriDateCroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	triParDate := func(slice1, slice2 int) bool { return slice1 < slice2 }
	sort.Slice(copyTable, func(i, j int) bool { return triParDate(copyTable[i].CreationDate, copyTable[j].CreationDate) })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

func TriDateDeroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	triParDate := func(slice1, slice2 int) bool { return slice1 > slice2 }
	sort.Slice(copyTable, func(i, j int) bool { return triParDate(copyTable[i].CreationDate, copyTable[j].CreationDate) })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

// -------------------------------------------
// Triage Alphabétique normal et inversé

func TrieNomGroupe(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	sort.Slice(copyTable, func(i, j int) bool { return copyTable[i].Name < copyTable[j].Name })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll

}

func TrieNomGroupeInverse(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	sort.Slice(copyTable, func(i, j int) bool { return copyTable[i].Name > copyTable[j].Name })

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll

}

// -------------------------------------------
//  Triage par Date d'album

func TriPremierAlbumCroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	sort.Slice(copyTable, func(i, j int) bool {
		dateI, errI := time.Parse("02-01-2006", copyTable[i].FirstAlbum)
		if errI != nil {
			// Gérer l'erreur
		}
		dateJ, errJ := time.Parse("02-01-2006", copyTable[j].FirstAlbum)
		if errJ != nil {
			// Gérer l'erreur
		}
		return dateI.Before(dateJ)
	})
	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}

func TriPremierAlbumDecroissant(w fyne.Window, tab []requestapi.Groupe) *container.Scroll {
	copyTable := make([]requestapi.Groupe, len(tab))
	copy(copyTable, tab)

	sort.Slice(copyTable, func(i, j int) bool {
		dateI, _ := time.Parse("02-01-2006", copyTable[i].FirstAlbum)
		dateJ, _ := time.Parse("02-01-2006", copyTable[j].FirstAlbum)
		return dateI.After(dateJ) // tri décroissant
	})

	grid := container.NewAdaptiveGrid(4)
	for i := range copyTable {
		img := lib.Img_button(copyTable[i])
		grid.Add(img)
	}
	scroll := container.NewHScroll(grid)
	scroll.Direction = container.ScrollBoth

	return scroll
}
