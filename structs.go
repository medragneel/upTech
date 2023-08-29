package main

type Movie struct {
	Adult           bool     `json:"adult"`
	BackdropPath    string   `json:"backdrop_path"`
	GenreIDs        []int    `json:"genre_ids"`
	ID              int      `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle   string   `json:"original_title"`
	Overview        string   `json:"overview"`
	Popularity      float64  `json:"popularity"`
	PosterPath      string   `json:"poster_path"`
	ReleaseDate     string   `json:"release_date"`
	Title           string   `json:"title"`
	Video           bool     `json:"video"`
	VoteAverage     float64  `json:"vote_average"`
	VoteCount       int      `json:"vote_count"`
}

type MovieList struct {
	Page         int      `json:"page"`
	Results      []Movie  `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

type TVShow struct {
	BackdropPath   string   `json:"backdrop_path"`
	FirstAirDate   string   `json:"first_air_date"`
	GenreIDs       []int    `json:"genre_ids"`
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	OriginCountry  []string `json:"origin_country"`
	OriginalLang   string   `json:"original_language"`
	OriginalName   string   `json:"original_name"`
	Overview       string   `json:"overview"`
	Popularity     float64  `json:"popularity"`
	PosterPath     string   `json:"poster_path"`
	VoteAverage    float64  `json:"vote_average"`
	VoteCount      int      `json:"vote_count"`
}

type TVShowList struct {
	Page           int       `json:"page"`
	Results        []TVShow  `json:"results"`
	TotalPages     int       `json:"total_pages"`
	TotalResults   int       `json:"total_results"`
}


type Result struct {
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	PosterPath       string  `json:"poster_path"`
	MediaType        string  `json:"media_type"`
	GenreIDs         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
	ReleaseDate      string  `json:"release_date"`
	Video            bool    `json:"video"`
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

type Response struct {
	Page         int      `json:"page"`
	Results      []Result `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

// type Anime struct {
//     AnimeID     string `json:"animeId"`
//     EpisodeID   string `json:"episodeId"`
//     AnimeTitle  string `json:"animeTitle"`
//     EpisodeNum  string `json:"episodeNum"`
//     SubOrDub    string `json:"subOrDub"`
//     AnimeImg    string `json:"animeImg"`
//     EpisodeUrl  string `json:"episodeUrl"`
// }



type Anime struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    URL         string `json:"url"`
    Image       string `json:"image"`
    ReleaseDate string `json:"releaseDate"`
    SubOrDub    string `json:"subOrDub"`
}

type AnimeList struct {
    CurrentPage  int     `json:"currentPage"`
    HasNextPage  bool    `json:"hasNextPage"`
    Results      []Anime `json:"results"`
}


type Episode struct {
	ID     string `json:"id"`
	Number int    `json:"number"`
	URL    string `json:"url"`
}

type AnimeInfo struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	URL          string   `json:"url"`
	Genres       []string `json:"genres"`
	TotalEpisodes int      `json:"totalEpisodes"`
	Image        string   `json:"image"`
	ReleaseDate  string   `json:"releaseDate"`
	Description  string   `json:"description"`
	SubOrDub     string   `json:"subOrDub"`
	Type         string   `json:"type"`
	Status       string   `json:"status"`
	OtherName    string   `json:"otherName"`
	Episodes     []Episode `json:"episodes"`
}
type Server struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}


type RequestHeaders struct {
	Referer    string `json:"Referer"`
	Watchsb    string `json:"watchsb"`
	UserAgent  string `json:"User-Agent"`
}

type Source struct {
	URL      string `json:"url"`
	Quality  string `json:"quality"`
	IsM3U8   bool   `json:"isM3U8"`
}

type RequestData struct {
	Headers  RequestHeaders `json:"headers"`
	Sources  []Source       `json:"sources"`
}

type Preview struct {
	Intro string `json:"intro"`
	Full  string `json:"full"`
}

type Feed struct {
	Title      string    `json:"title"`
	ID         string    `json:"id"`
	UploadedAt string    `json:"uploadedAt"`
	Topics     []string  `json:"topics"`
	Preview    Preview   `json:"preview"`
	Thumbnail  string    `json:"thumbnail"`
	URL        string    `json:"url"`
}
