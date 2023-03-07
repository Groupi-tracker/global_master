package requestapi

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
