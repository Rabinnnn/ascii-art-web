package main

import (
	
	"html/template"
	"log"
	"net/http"
	"os"

	asciiart "ascii-art-web/ascii-art"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "405 (Method Not Allowed)", http.StatusMethodNotAllowed)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/result.html"))
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "400 (Bad Request)", http.StatusBadRequest)
			return
		}
		// if err != nil {
		// 	fmt.Fprintf(w, "Error aparsing the form data")
		// 	return
		// }
		bannerFile := r.FormValue("banner")
		input := r.FormValue("input")
		// fmt.Println(input)
		ascii := asciiart.Art(input, asciiart.BannerArt(asciiart.FileName(bannerFile)))
		// fmt.Print(ascii.String())
		tmpl.Execute(w, ascii.String())
		// w.Write([]byte(ascii))
	})

	log.Println("Server listening on http://localhost:8080")
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
				http.NotFound(w, r) // Use built-in handler for 404
                return
			}
			tpl := template.Must(template.ParseFiles("templates/test.html"))
			tpl.Execute(w, nil)
			return
		}
	}else if r.URL.Path == "/about" || r.URL.Path == "/About" {
		if r.Method == http.MethodGet {
			_, err := os.Stat("templates/about.html")
			if os.IsNotExist(err) {
				http.NotFound(w, r) 
                return
			}
			tpl := template.Must(template.ParseFiles("templates/about.html"))
			tpl.Execute(w, nil)
			return
		}
	} else {
        // Handle unexpected methods or paths
        if r.Method == http.MethodGet {
            http.NotFound(w, r) // Use built-in handler for 404
            return
        }
        http.Error(w, "400 (Bad Request)", http.StatusBadRequest) // Clear error message for 400
        return
    }
http.Error(w, "500 (Internal Server Error)", http.StatusInternalServerError)
}
