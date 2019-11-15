package homepage

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
	logger *log.Logger
}

type Post struct {
	Num    int     `json:"num"`
	Img    string  `json:"img"`
	Alt    string  `json:"alt"`
	Rating float32 `json:"rating"`
	Title  string  `json:"title"`
}

var posts []Post

func init() {
	posts = append(posts, Post{
		Num:    1,
		Title:  "Where to?",
		Rating: 3,
		Img:    "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
		Alt:    "Don't we all."})
	posts = append(posts, Post{
		Num:    4,
		Title:  "Marvellous Mrs. Maisel",
		Rating: 5,
		Img:    "https://imgs.xkcd.com/comics/island_color.jpg",
		Alt:    "Hello, island"})
	posts = append(posts, Post{
		Num:    2,
		Title:  "Le Petit Prince",
		Rating: 4,
		Img:    "https://imgs.xkcd.com/comics/tree_cropped_(1).jpg",
		Alt:    "'Petit' being a reference to Le Petit Prince, which I only thought about halfway through the sketch"})
	posts = append(posts, Post{
		Num:    3,
		Title:  "The lost Island",
		Rating: 4.5,
		Img:    "https://imgs.xkcd.com/comics/island_color.jpg",
		Alt:    "Hello, island"})
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	json.NewEncoder(w).Encode(posts)
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.logger.Printf("request processed")
		defer fmt.Printf("request processed")
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/getMovies", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http:localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
