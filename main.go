package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
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
func api_artists(w http.ResponseWriter, r *http.Request) ([]Artists, bool) {

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

// artists serves the route "/artists"
func artists(w http.ResponseWriter, r *http.Request) {
	//parsing the artist page
	file, err := template.ParseFiles("testing_templates/artist.html")
	if err != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	//storing the fetch datas
	res, error := api_artists(w, r)

	//if an error occured while fetching datas
	if !error {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	err=file.Execute(w, res)
	if err != nil  {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
	}
}

// api_dates fetch and returns datas related to "dates" collected from the groupie json file
// it also returns an error boolean when an problem while fetching, reading or converting the
// datas occurs
func api_dates(w http.ResponseWriter, r *http.Request) (date, bool) {
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
func api_locations(w http.ResponseWriter, r *http.Request) (locations, bool) {

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
func api_relation(w http.ResponseWriter, r *http.Request) (relations, bool) {
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

func home(w http.ResponseWriter, r *http.Request) {

	//checking whether the route exists or not
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/info/" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	file, err := template.ParseFiles("testing_templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	tab, error := api_artists(w, r)
	if !error {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	one := tab[0:6]

	err=file.Execute(w, struct {
		Top_artist []Artists
	}{
		Top_artist: one,
	})
	if err != nil  {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
	}
}

//info serve the route ("/info"). 
func info(w http.ResponseWriter, r *http.Request) {

	//retrieving the id from the url
	recup_id := path.Base(r.URL.Path)
	if recup_id == "" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

//storing the api artist datas
	artists_data, error := api_artists(w, r)
	//if an error occured while fetching datas
	if !error {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	
	//converting the id into int and setting a limit
	id, err := strconv.Atoi(recup_id)
	if err != nil || id==0 || id > len(artists_data) {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "404")
		return
	}
//retrieving the informations corresponding to the id 
	var artists_checked Artists
	for _, art := range artists_data {
		if art.Id == id {
			artists_checked = art
			break
		}
	}

	//storing the api dates datas
	dates_data, error1 := api_dates(w, r)
	//if an error occured while fetching datas
	if !error1 {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id 
	var dates_checked Dates
	for _, days := range dates_data.Index {
		if days.Id == id {
			dates_checked = days
			break
		}
	}

	//storing the api locations datas
	locations_data, error2 := api_locations(w, r)
	//if an error occured while fetching datas
	if !error2 {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id 
	var locations_checked Locations
	for _, city := range locations_data.Index {
		if city.Id == id {
			locations_checked = city
			break
		}
	}

	//storing the api relations datas
	relations_data, error3 := api_relation(w, r)

	//if an error occured while fetching datas
	if !error3 {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id 
	var relations_checked Relations
	for _, link := range relations_data.Index {
		if link.Id == id {
			relations_checked = link
			break
		}
	}

	//struct to excecute
	todisplay := struct {
		The_arts Artists
		Days     Dates
		Cities   Locations
		Links    Relations
	}{
		The_arts: artists_checked,
		Days:     dates_checked,
		Cities:   locations_checked,
		Links:    relations_checked,
	}

	file, errp := template.ParseFiles("testing_templates/info.html")
	if errp != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	err = file.Execute(w, todisplay)
	if err != nil  {
		error_file := template.Must(template.ParseFiles("testing_templates/error.html"))
		error_file.Execute(w, "500")
	}
}

// handlers regroups all the routes supported by our servor.
// handlers launches it too.
func handlers() {
	fs := http.FileServer(http.Dir("testing_templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)
	http.HandleFunc("/artists", artists)
	http.HandleFunc("/info/", info)
	fmt.Println("server has started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handlers()
}
