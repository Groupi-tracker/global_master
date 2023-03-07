package lib

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	requestapi "groupie-tracker/RequestAPI"
)

func Img_button(s requestapi.Groupe) *fyne.Container { // return type
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

func Tab_to_string(s []string) string {
	var tmp string

	for i := range s {
		tmp = fmt.Sprintf("%s %s", tmp, s[i])
	}
	return tmp
}

func FormatString(s string) ([]string, []string) {
	var res []string
	var tmp string

	j := 0
	n := 0

	for _, i := range s {
		if ((rune(i) >= 'a' && rune(i) <= 'z') || (rune(i) >= 'A' && rune(i) <= 'Z')) && rune(i) != ' ' && rune(i) != '-' {
			j++
		}
		if rune(i) == ':' {
			n++
		}
		if n == 2 {
			break
		}
	}
	for i := j; i != len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && s[i] != ':' {
			tmp += string(s[i])
		} else {
			res = append(res, tmp)
			tmp = ""
		}
	}
	res = append(res, tmp)
	var date []string
	for _, i := range res {
		if i != "" {
			if len(i) > 2 {
				date = append(date, i)
			}
		}
	}

	res = nil
	tmp = ""
	var tmpMaps []string
	var maps []string
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i == '-' || i == '_') {
			tmp += string(i)
		} else {
			if tmp != "" {
				if tmp[0] != '-' {
					res = append(res, tmp)
				}
			}
			tmp = ""
		}
	}
	for i := range res {
		if res[i] != "" {
			tmpMaps = append(tmpMaps, res[i])
		}
	}
	for i := 1; i < len(tmpMaps); i++ {
		s1, s2 := SepString(tmpMaps[i])
		maps = append(maps, fmt.Sprintf("%s [%s]", s1, s2))
	}

	return maps, date
}

func FormatStringGeo(s string) ([]string, []string) {
	var res []string
	var tmp string

	j := 0
	n := 0

	for _, i := range s {
		if ((rune(i) >= 'a' && rune(i) <= 'z') || (rune(i) >= 'A' && rune(i) <= 'Z')) && rune(i) != ' ' && rune(i) != '-' {
			j++
		}
		if rune(i) == ':' {
			n++
		}
		if n == 2 {
			break
		}
	}
	for i := j; i != len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && s[i] != ':' {
			tmp += string(s[i])
		} else {
			res = append(res, tmp)
			tmp = ""
		}
	}
	res = append(res, tmp)
	var date []string
	for _, i := range res {
		if i != "" {
			if len(i) > 2 {
				date = append(date, i)
			}
		}
	}

	res = nil
	tmp = ""
	var tmpMaps []string
	var maps []string
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i == '-' || i == '_') {
			tmp += string(i)
		} else {
			if tmp != "" {
				if tmp[0] != '-' {
					res = append(res, tmp)
				}
			}
			tmp = ""
		}
	}
	for i := range res {
		if res[i] != "" {
			tmpMaps = append(tmpMaps, res[i])
		}
	}
	for i := 1; i < len(tmpMaps); i++ {
		s1, _ := SepString(tmpMaps[i])
		maps = append(maps, s1)
	}

	return maps, date
}

func SepString(s string) (string, string) {
	var tmp1 string
	var tmp2 string
	n := 0

	for _, i := range s {
		if i == '-' {
			break
		}
		tmp1 += string(i)
		n++
	}
	for i := n + 1; i < len(s); i++ {
		tmp2 += string(s[i])
	}
	return tmp1, tmp2
}
