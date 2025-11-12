package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/main", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../assets/loginPage.html")
	http.ServeFile(w, r, "../../assets/loginScript.js")
	http.ServeFile(w, r, "../../assets/loginStyle.css")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../assets/mainPage.html")
	http.ServeFile(w, r, "../../assets/mainStyle.css")
}
