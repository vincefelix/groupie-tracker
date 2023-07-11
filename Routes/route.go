package Func

import (
	fetch "Func/API"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// home serves the route "/"
func Home(w http.ResponseWriter, r *http.Request) {

	//checking whether the route exists or not
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/info/" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	file, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	tab, error := fetch.Api_artists(w, r)
	if !error {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	one := tab[0:6]

	err = file.Execute(w, struct {
		Top_artist []fetch.Artists
	}{
		Top_artist: one,
	})
	if err != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
}

// artists serves the route "/artists"
func Artists(w http.ResponseWriter, r *http.Request) {
	//parsing the artist page
	file, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	//storing the fetch datas
	res, error := fetch.Api_artists(w, r)

	//if an error occured while fetching datas
	if !error {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	err = file.Execute(w, res)
	if err != nil {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
}

// info serve the route ("/info").
func Info(w http.ResponseWriter, r *http.Request) {

	//retrieving the id from the url
	recup_id := path.Base(r.URL.Path)
	if recup_id == "" {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}

	//storing the api artist datas
	artists_data, error := fetch.Api_artists(w, r)
	//if an error occured while fetching datas
	if !error {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	//converting the id into int and setting a limit
	id, err := strconv.Atoi(recup_id)
	if err != nil || id == 0 || id > len(artists_data) {
		w.WriteHeader(http.StatusNotFound)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "404")
		return
	}
	//retrieving the informations corresponding to the id
	var artists_checked fetch.Artists
	for _, art := range artists_data {
		if art.Id == id {
			artists_checked = art
			break
		}
	}

	//storing the api dates datas
	dates_data, error1 := fetch.Api_dates(w, r)
	//if an error occured while fetching datas
	if !error1 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id
	var dates_checked fetch.Dates
	for _, days := range dates_data.Index {
		if days.Id == id {
			dates_checked = days
			break
		}
	}

	//Removing the "*" in front of dates
	for i := range dates_checked.Date {
		dates_checked.Date[i] = strings.ReplaceAll(dates_checked.Date[i], "*", "")
	}

	//storing the api locations datas
	locations_data, error2 := fetch.Api_locations(w, r)
	//if an error occured while fetching datas
	if !error2 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id
	var locations_checked fetch.Locations
	for _, city := range locations_data.Index {
		if city.Id == id {
			locations_checked = city
			break
		}
	}

	//storing the api relations datas
	relations_data, error3 := fetch.Api_relation(w, r)

	//if an error occured while fetching datas
	if !error3 {
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}
	//retrieving the informations corresponding to the id
	var relations_checked fetch.Relations
	for _, link := range relations_data.Index {
		if link.Id == id {
			relations_checked = link
			break
		}
	}

	// modifying the relation map
	newmap := map[string] []string{}
	for position, day:= range relations_checked.Dates_location {
		splitted := strings.ReplaceAll(position, "-", "\n")
		newmap[splitted] = day
	}

	previd := artists_checked.Id - 1
	nextid := artists_checked.Id + 1
	//struct to excecute
	todisplay := struct {
		The_arts fetch.Artists
		Days     fetch.Dates
		Cities   fetch.Locations
		Links    map[string] []string
		Prev     int
		Next     int
	}{
		The_arts: artists_checked,
		Days:     dates_checked,
		Cities:   locations_checked,
		Links:    newmap,
		Prev:     previd,
		Next:     nextid,
	}

	file, errp := template.ParseFiles("templates/info.html")
	if errp != nil {
		//sending metadata about the error to the servor
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
		return
	}

	err = file.Execute(w, todisplay)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_file := template.Must(template.ParseFiles("templates/error.html"))
		error_file.Execute(w, "500")
	}
}
