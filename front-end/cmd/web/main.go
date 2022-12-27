package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	HOST := "localhost"
	PORT := 8080
	ADDRESS := fmt.Sprintf("%s:%d", HOST, PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Printf("Starting front end service on %s:%d\n", HOST, PORT)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./front-end/cmd/web/templates/base.layout.gohtml",
		"./front-end/cmd/web/templates/header.partial.gohtml",
		"./front-end/cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./front-end/cmd/web/templates/%s", t))

	// for _, x := range partials {
	// 	templateSlice = append(templateSlice, x)
	// }
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
