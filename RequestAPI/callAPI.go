package requestapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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
