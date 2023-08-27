package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getenv(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	auth := os.Getenv(key)
	return auth, nil
}
func home(w http.ResponseWriter, r *http.Request) {
	a, err := getenv("AUTH")
	print(fmt.Sprintf("Bearer %s", a))
	if err != nil {
		log.Println("Error loading environment variable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var response Response
	query := r.FormValue("q")
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/multi?query=%s&include_adult=false&language=en-US&page=1", query)

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

	fmt.Println("Response Body:", string(body))
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, response)
	defer res.Body.Close()

}
func fetchMovies(w http.ResponseWriter, r *http.Request) {
	a, err := getenv("AUTH")
	print(fmt.Sprintf("Bearer %s", a))
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
	tmpl, err := template.ParseFiles("./static/movies.html")
	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, Movies)
	// fmt.Fprintf(w,string(body))

	defer res.Body.Close()
}

// fetch tv Shows
func fetchTvShows(w http.ResponseWriter, r *http.Request) {
	a, err := getenv("AUTH")
	print(fmt.Sprintf("Bearer %s", a))
	if err != nil {
		log.Println("Error loading environment variable:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var tvshows TVShowList
	url := "https://api.themoviedb.org/3/trending/tv/day?language=en-US"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + a)

	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &tvshows); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./static/tv.html")
	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, tvshows)
	// fmt.Fprintf(w,string(body))

	defer res.Body.Close()
}

// fetch Anime
func fetchAnime(w http.ResponseWriter, r *http.Request) {
	var animes AnimeList
	// vars := mux.Vars(r)
	vals := r.URL.Query()
	query := vals.Get("q")
	url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/%s", query)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &animes); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./static/anime.html")

	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, animes)
	defer res.Body.Close()

}

// get anime info

func getAnimeInfo(w http.ResponseWriter, r *http.Request) {
	var infos AnimeInfo
	vars := mux.Vars(r)
	id := vars["id"]
	url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/info/%s", id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &infos); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./static/animeInfo.html")

	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, infos)
	defer res.Body.Close()

}

func getServers(w http.ResponseWriter, r *http.Request) {
	// var servers []Server
	var response RequestData
	vars := mux.Vars(r)
	id := vars["epid"]
	// url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/servers/%s", id)
	url := fmt.Sprintf("http://localhost:3000/anime/gogoanime/watch/%s", id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("Error in Json Unmarshal")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("./static/servers.html")

	if err != nil {
		panic(err)

	}
	tmpl.Execute(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

}

func main() {

	fmt.Println("server running on port :5000")
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/css/").Handler(
		http.StripPrefix("/css/", http.FileServer(http.Dir("static/css/"))),
	)
	r.HandleFunc("/", home)
	r.HandleFunc("/movies", fetchMovies)
	r.HandleFunc("/tv-shows", fetchTvShows)
	r.HandleFunc("/anime", fetchAnime)
	r.HandleFunc("/{id}", getAnimeInfo)
	r.HandleFunc("/watch/{epid}", getServers)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":5000", r))

}
