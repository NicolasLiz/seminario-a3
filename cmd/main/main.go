package main

import (
	"a3/internal/keys"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", login)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/main", mainPage)

	http.HandleFunc("/api/publicKey", keys.GetPublicKey)
	http.HandleFunc("/submit", keys.DecryptMessage)

	//servir com https
	//log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/loginPage.html")
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/html/mainPage.html")
}
