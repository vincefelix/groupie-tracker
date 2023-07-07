package handlers

import (
	Rt "Func/Routes"
	"fmt"
	"log"
	"net/http"
)

// handlers regroups all the routes supported by our servor.
// handlers launches it too.
func Handlers() {
	fs := http.FileServer(http.Dir("testing_templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Rt.Home)
	http.HandleFunc("/artists", Rt.Artists)
	http.HandleFunc("/info/", Rt.Info)
	fmt.Println("server has started at : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
