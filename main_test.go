package main

import (
	Func "Func/Routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Requete HTTP test de la route "/"
	requete, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enregistreur de reponse HTTP de test
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/"
	Func.Home(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestArtistsHandler(t *testing.T) {
	// Requete HTTP test de la route "/artists"
	requete, err := http.NewRequest("GET", "/artists", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Enregistreur de reponse HTTP de test
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/artists"
	Func.Artists(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestInfoHandler(t *testing.T) {
	// Requete HTTP pour "/info/{id}"
	requete, err := http.NewRequest("GET", "/info/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Enregistreur de reponse HTTP
	enregistre := httptest.NewRecorder()

	// Appelle du handler de la route "/info/{id}"
	Func.Info(enregistre, requete)

	// Vérifier si les codes de statut correspondent
	if status := enregistre.Code; status != http.StatusOK {
		t.Errorf("Wrong code returned. Expected %v, got %v", http.StatusOK, status)
	}
}

func TestHandlers(t *testing.T) {
	// Applle des fonctions de test handler
	t.Run("Home", TestHomeHandler)
	t.Run("Artists", TestArtistsHandler)
	t.Run("Info", TestInfoHandler)
}
