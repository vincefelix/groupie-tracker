package Func

import (
	"encoding/json"
	"io"
	"net/http"
)

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

// get_link returns the links of the api's elements
func get_link(w http.ResponseWriter, r *http.Request) (string, string, string, string) {
	var artist_link, dates_link, location_link, relation_link string

	//fetching datas from the api
	data, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		http.Error(w, "error while fetching data", http.StatusInternalServerError)
	} else {
		//reading the collected datas
		content, ex := io.ReadAll(data.Body)
		if ex != nil {
			http.Error(w, "error while reading the content", http.StatusInternalServerError)
		} else {

			//converting the json file and storing the datas
			var info Band
			err = json.Unmarshal(content, &info)
			if err != nil {
				http.Error(w, "error while reading the content", http.StatusInternalServerError)
			} else {
				artist_link, dates_link, location_link, relation_link = info.Art, info.Dat, info.Loc, info.Rel // getting the links
			}
		}
	}
	return artist_link, dates_link, location_link, relation_link
}

// api_artists fetches and returns datas related to "artists" collected from the groupie json file
// it also returns an error boolean when an problem while fetching, reading or converting the
// datas occurs
func Api_artists(w http.ResponseWriter, r *http.Request) ([]Artists, bool) {

	var affiche_err = true
	// Getting the link heading to datas
	link_artist, _, _, _ := get_link(w, r)

	// Fetching datas
	data, err := http.Get(link_artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return nil, affiche_err
	}

	// Reading the collected datas
	content, err := io.ReadAll(data.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return nil, affiche_err
	}

	// Converting the JSON file and storing the datas
	var artist []Artists
	err = json.Unmarshal(content, &artist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return nil, affiche_err
	}

	return artist, affiche_err
}

// api_dates fetch and returns datas related to "dates" collected from the groupie json file
// it also returns an error boolean when an problem while fetching, reading or converting the
// datas occurs
func Api_dates(w http.ResponseWriter, r *http.Request) (date, bool) {
	affiche_err := true
	//getting the link heading to datas
	_, link_date, _, _ := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_date)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return date{}, affiche_err
	}

	//reading the collected datas
	content, ex := io.ReadAll(data.Body)
	if ex != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return date{}, affiche_err
	}

	//converting the json file and storing the datas
	var thedate date
	err = json.Unmarshal(content, &thedate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return date{}, affiche_err
	}

	return thedate, affiche_err
}

// api_locations fetch and returns datas related to "locations" collected from the groupie json file
// it also returns an error boolean when an problem while fetching, reading or converting the
// datas occurs
func Api_locations(w http.ResponseWriter, r *http.Request) (locations, bool) {

	affiche_err := true
	//getting the link heading to datas
	_, _, link_location, _ := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return locations{}, affiche_err
	}

	//reading the collected datas
	content, ex := io.ReadAll(data.Body)
	if ex != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return locations{}, affiche_err
	}

	//converting the json file and storing the datas
	var zone locations
	err = json.Unmarshal(content, &zone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return locations{}, affiche_err
	}

	return zone, affiche_err
}

// api_relation fetchs and returns datas related to "relations" collected from the groupie json file
// it also returns an error boolean when an problem while fetching, reading or converting the
// datas occurs
func Api_relation(w http.ResponseWriter, r *http.Request) (relations, bool) {
	affiche_err := true
	//getting the link heading to datas
	_, _, _, link_relation := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_relation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return relations{}, affiche_err
	}

	//reading the collected datas
	content, ex := io.ReadAll(data.Body)
	if ex != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return relations{}, affiche_err
	}
	//converting the json file and storing the datas
	var linked relations
	err = json.Unmarshal(content, &linked)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		affiche_err = false
		return relations{}, affiche_err
	}
	return linked, affiche_err
}
