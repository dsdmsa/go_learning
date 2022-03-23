package main

import (
	"encoding/json"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {

	playerChoiche, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoiche)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println("error parsing,", err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println("error executing,", err)
		return
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)
	log.Println("server started ")
	http.ListenAndServe(":8080", nil)
}
