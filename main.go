package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// const baseUrl = "http://localhost:8000/"
const baseUrl = "https://api-for-netlify.vercel.app/"

func getenv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	auth := os.Getenv(key)
	return auth, nil
}
func home(w http.ResponseWriter, r *http.Request) {
	var feeds []Feed
	files := []string{
		"./static/base.html",
		"./static/index.html",
	}
	url := baseUrl + "news/ann/recent-feeds"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &feeds); err != nil {
		log.Println("Error in JSON Unmarshal:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return

	}
	err = tmpl.ExecuteTemplate(w, "base", feeds)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}

func search(w http.ResponseWriter, r *http.Request) {
	a, err := getenv("AUTH")
	if err != nil {
		log.Println("Error loading environment variable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	files := []string{
		"./static/base.html",
		"./static/search.html",
	}
	var response Response
	query := r.FormValue("q")
    encodedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/multi?query=%s&include_adult=false&language=en-US&page=1", encodedQuery)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+a)

	res, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("Error in JSON Unmarshal:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return

	}
	err = tmpl.ExecuteTemplate(w, "base", response)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	defer res.Body.Close()

}
func fetchMovies(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./static/base.html",
		"./static/movies.html",
	}
	a, err := getenv("AUTH")
	if err != nil {
		log.Println("Error loading environment variable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var Movies MovieList
	url := "https://api.themoviedb.org/3/trending/movie/day?language=en-US"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+a)
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &Movies); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)

	}
	tmpl.ExecuteTemplate(w, "base", Movies)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	// fmt.Fprintf(w,string(body))

	defer res.Body.Close()
}

// fetch tv Shows
func fetchTvShows(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./static/base.html",
		"./static/tv.html",
	}
	a, err := getenv("AUTH")
	if err != nil {
		log.Println("Error loading environment variable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var tvshows TVShowList
	url := "https://api.themoviedb.org/3/trending/tv/day?language=en-US"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+a)

	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &tvshows); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)

	}
	tmpl.ExecuteTemplate(w, "base", tvshows)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	// fmt.Fprintf(w,string(body))

	defer res.Body.Close()
}

// fetch Anime
func fetchAnime(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./static/base.html",
		"./static/anime.html",
	}
	var animes AnimeList
	// vars := mux.Vars(r)
	vals := r.URL.Query()
	query := vals.Get("q")
	url := fmt.Sprintf(baseUrl+"anime/gogoanime/%s", query)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &animes); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}
	err = tmpl.ExecuteTemplate(w, "base", animes)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	defer res.Body.Close()

}

// get anime info

func getAnimeInfo(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./static/base.html",
		"./static/animeInfo.html",
	}
	var infos AnimeInfo
	vars := mux.Vars(r)
	id := vars["id"]
	url := fmt.Sprintf(baseUrl+"anime/gogoanime/info/%s", id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &infos); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}
	err = tmpl.ExecuteTemplate(w, "base", infos)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	defer res.Body.Close()

}

func getServers(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./static/base.html",
		"./static/servers.html",
	}
	// var servers []Server
	var response RequestData
	vars := mux.Vars(r)
	id := vars["epid"]
	// url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/servers/%s", id)
	url := fmt.Sprintf(baseUrl+"anime/gogoanime/watch/%s", id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles(files...)

	if err != nil {
		panic(err)

	}
	err = tmpl.ExecuteTemplate(w, "base", response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

}

func main() {
    port,_ := getenv("PORT")

    fmt.Printf("server running on port :%s",port)
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/css/").Handler(
		http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))),
	)
	r.HandleFunc("/", home)
	r.HandleFunc("/search", search)
	r.HandleFunc("/movies", fetchMovies)
	r.HandleFunc("/tv-shows", fetchTvShows)
	r.HandleFunc("/anime", fetchAnime)
	r.HandleFunc("/{id}", getAnimeInfo)
	r.HandleFunc("/watch/{epid}", getServers)
	http.Handle("/", r)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port), r))

}
