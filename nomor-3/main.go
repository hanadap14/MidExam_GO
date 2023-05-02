package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var people = []Person{
	{Nama: "Dimas", NIM: "250174933", Alamat: "Balikpapan"},
	{Nama: "Diyo", NIM: "250289233", Alamat: "Kutai Kartanegara"},
}

type Person struct {
	Nama   string `json:"NAMA"`
	NIM    string `json:"NIM"`
	Alamat string `json:"ALAMAT"`
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {

	count := 0

	if r.Method == "POST" {
		GetNamePerson := Person{
			Nama: r.FormValue("NAMA"),
		}

		for _, value := range people {
			if strings.EqualFold(GetNamePerson.Nama, value.Nama) {
				response, err := json.Marshal(Person{
					Nama:   value.Nama,
					NIM:    value.NIM,
					Alamat: value.Alamat,
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				count++

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}

		if count == 0 {
			http.Error(w, "NAME NOT FOUND", http.StatusMethodNotAllowed)
			return
		}

	} else {
		http.Error(w, "METHOD NOT FOUND", http.StatusMethodNotAllowed)
	}

}

func GetAllDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		personJson, _ := json.Marshal(people)

		w.Header().Set("Content-Type", "application/json")
		w.Write(personJson)
		return

	} else {
		http.Error(w, "METHOD NOT FOUND", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/nama", PersonHandler)
	mux.HandleFunc("/data", GetAllDataHandler)

	http.ListenAndServe(":8080", mux)
}

