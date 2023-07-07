package Func


// data structure from the artists informationss
type Artists struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Member        []string `json:"members"`
	Creation_date int      `json:"creationDate"`
	First_album   string   `json:"firstAlbum"`
}

// data structure from the date informations
type date struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

// data structure from the location informations
type locations struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}

// data structure from the relation informations
type relations struct {
	Index []Relations `json:"index"`
}

type Relations struct {
	Id             int                 `json:"id"`
	Dates_location map[string][]string `json:"datesLocations"`
}

// data structure from the API
type Band struct {
	Art string `json:"artists"`
	Dat string `json:"dates"`
	Loc string `json:"locations"`
	Rel string `json:"relation"`
}
