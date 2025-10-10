package main

import 
(
	"math/rand"
	"time"
	"net/http"
)

func main() {
	
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	mux := http.NewServeMux();
	mux.HandleFunc(`POST /`, createHandler);
	mux.HandleFunc(`GET /{id}`, getByIdHandler);

	return http.ListenAndServe(`:8080`, mux)
}

var links = make(map[string]string)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(n int) string {
	var b = make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
    return string(b)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    
	var id string
	for {
		id = RandomString(8)
		if _, ok := links[id]; !ok {
			break
		}
	}

	var link = r.Body.ReadAll()
	links[id] = link;

    r.WriteHeader(http.StatusCreated);
	w.Write("http://localhost:8080/" + id)

}

func getByIdHandler(w http.ResponseWriter, r *http.Request) {

	var id = r.PathValue(`id`)
    
	var link, ok = links[id]
	if !ok {
		w.WriteHeader(http.NotFound);
		return;
	}

	w.WriteHeader(http.StatusTemporaryRedirect);
	w.Write(link);
}