package main

import(
    "net/http"
	"html/template"
    "3am_server/getMovies"
    "fmt"
)

type NextPage struct {
    Genre string
    Title[] string
    Path string
}

var title[]string

func indexHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
}

func nextHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    
    movies := getMovies.GetMovies(r.FormValue("genre"))

    for i := 0; i < len(movies); i++ {
        title := movies[i].MTitle
        fmt.Println(title)
    }

    //p := movies
    p := NextPage{Genre: r.FormValue("genre"), Title: title}
    fmt.Println(len(movies))
    fmt.Println(movies)
    t, _ := template.ParseFiles("mvtitles.html")
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/next", nextHandler)

    http.ListenAndServe(":8080", nil)
}