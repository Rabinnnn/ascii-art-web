package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	asciiart "ascii-art-web/ascii-art"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/result.html"))
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error aparsing the form data")
			return
		}
		bannerFile := r.FormValue("banner")
		input := r.FormValue("input")
		// fmt.Println(input)
		ascii := asciiart.Art(input, asciiart.BannerArt(asciiart.FileName(bannerFile)))
		// fmt.Print(ascii.String())
		tmpl.Execute(w, ascii.String())
		// w.Write([]byte(ascii))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" || r.URL.Path == "/Home" {
		if r.Method == http.MethodGet {
			_, err := os.Stat("templates/test.html")
			if os.IsNotExist(err) {
				fmt.Fprintf(w, "Not Found")
				return
			}
			tpl := template.Must(template.ParseFiles("templates/test.html"))
			tpl.Execute(w, nil)
			return
		}
	}
	if r.URL.Path == "/about" || r.URL.Path == "/About" {
		if r.Method == http.MethodGet {
			_, err := os.Stat("templates/about.html")
			if os.IsNotExist(err) {
				fmt.Fprintf(w, "Not Found")
				return
			}
			tpl := template.Must(template.ParseFiles("templates/about.html"))
			tpl.Execute(w, nil)
			return
		}
	}
	// 	else if r.Method == http.MethodPost {
	// 		err := r.ParseForm()
	// 		if err != nil {
	// 			fmt.Fprintf(w, "Error aparsing the form data")
	// 			return
	// 		}
	// 		bannerFile := r.FormValue("banner")
	// 		_, errr := os.Stat(bannerFile + ".txt")
	// 		if os.IsNotExist(errr) {
	// 		}
	// 		input := r.FormValue("input")
	// 		fmt.Fprintf(w, "%s\n", input)
	// 		fmt.Fprintf(w, "%s\n", bannerFile)
	// 		tpl := template.Must(template.ParseFiles("templates/test.html"))
	// 		tpl.Execute(w, nil)
	// 	} else {
	// 		fmt.Fprintf(w, "Wrong request method")
	// 	}
	// }
	// fmt.Fprintf(w, "Bad request")
}
