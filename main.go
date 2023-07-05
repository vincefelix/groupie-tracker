package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

// get link returns the links of the api's elements
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

// api_artists fetches datas related to "/artists" collected from the groupie json file
func api_artists(w http.ResponseWriter, r *http.Request) []Artists {

	// Getting the link heading to datas
	link_artist, _, _, _ := get_link(w, r)

	// Fetching datas
	data, err := http.Get(link_artist)
	if err != nil {
		http.Error(w, "error while fetching datas", http.StatusInternalServerError)
		return nil
	}

	// Reading the collected datas
	content, err := io.ReadAll(data.Body)
	if err != nil {
		http.Error(w, "error while reading the content", http.StatusInternalServerError)
		return nil
	}

	// Converting the JSON file and storing the datas
	var artist []Artists
	err = json.Unmarshal(content, &artist)
	if err != nil {
		http.Error(w, "error while converting json", http.StatusInternalServerError)
		return nil
	}

	return artist
}

func artists(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "500: internal servor error", http.StatusInternalServerError)
		return
	}
	res := api_artists(w, r)
	file.Execute(w, res)
}

// api_dates fetch and displays datas related to "/dates" collected from the groupie json file
func api_dates(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("templates/dates.html")
	if err != nil {
		http.Error(w, "500: internal servor error", http.StatusInternalServerError)
		return
	}

	//getting the link heading to datas
	_, link_date, _, _ := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_date)
	if err != nil {
		http.Error(w, "error while fetching data", http.StatusInternalServerError)
		return
	}

	//reading the collected datas
	content, ex := io.ReadAll(data.Body)
	if ex != nil {
		http.Error(w, "error while reading the content", http.StatusInternalServerError)
		return
	}

	//converting the json file and storing the datas
	var thedate date
	err = json.Unmarshal(content, &thedate)
	if err != nil {
		http.Error(w, "error while converting json", http.StatusInternalServerError)
		return
	}

	err = file.Execute(w, thedate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// api_locations fetch and displays datas related to "/locations" collected from the groupie json file
func api_locations(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//getting the link heading to datas
	_, _, link_location, _ := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_location)
	if err != nil {
		http.Error(w, "error while fetching data", http.StatusInternalServerError)
		return
	}

	//reading the collected datas
	content, ex := io.ReadAll(data.Body)
	if ex != nil {
		http.Error(w, "error while reading the content", http.StatusInternalServerError)
		return
	}

	//converting the json file and storing the datas
	var zone locations
	err = json.Unmarshal(content, &zone)
	if err != nil {
		http.Error(w, "error while converting json", http.StatusInternalServerError)
		return
	}
	err = file.Execute(w, zone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// api_relation fetch and displays datas related to "/relations" collected from the groupie json file
func api_relation(w http.ResponseWriter, r *http.Request) {
	// file, err := template.ParseFiles("templates/locations.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// } else {
	// 	file.Execute(w, r)
	// }

	//getting the link heading to datas
	_, _, _, link_relation := get_link(w, r)

	//fetching datas
	data, err := http.Get(link_relation)
	if err != nil {
		http.Error(w, "error while fetching data", http.StatusInternalServerError)
	} else {

		//reading the collected datas
		content, ex := io.ReadAll(data.Body)
		if ex != nil {
			http.Error(w, "error while reading the content", http.StatusInternalServerError)
		} else {

			//converting the json file and storing the datas
			var linked relations
			err = json.Unmarshal(content, &linked)
			if err != nil {
				http.Error(w, "error while converting json", http.StatusInternalServerError)
			} else {
				for _, v := range linked.Index {
					fmt.Fprintf(w, "Dates and locations:\n%v\n\n\n\n", v.Dates_location)
				}
			}
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	//checking wether the route exists or not
	if r.URL.Path != "/" && r.URL.Path != "/artists" && r.URL.Path != "/locations" && r.URL.Path != "/dates" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}

	file, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tab := api_artists(w, r)
	one := tab[0:5]
	two := tab[5:9]
	file.Execute(w, struct {
		Top_artist []Artists
		Top_album  []Artists
	}{
		Top_artist: one,
		Top_album:  two,
	})

}

// handlers regroups all the routes supported by our servor.
// handlers launches it too.
func handlers() {
	http.HandleFunc("/", home)
	http.HandleFunc("/artists", artists)
	http.HandleFunc("/locations", api_locations)
	http.HandleFunc("/dates", api_dates)
	http.HandleFunc("/relations", api_relation)
	fmt.Println("server has started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handlers()
}
