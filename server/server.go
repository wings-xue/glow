package serve

import (
	"fmt"
	"grow/ac"
	"net/http"
)

var grow = ac.New()

func GetUrlArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
	return arg
}

func add(w http.ResponseWriter, req *http.Request) {
	keyword := GetUrlArg(req, "keyword")
	grow.AddWord(keyword)
	fmt.Fprintf(w, keyword)

}

func extract(w http.ResponseWriter, req *http.Request) {
	text := GetUrlArg(req, "text")
	out := grow.ExtractSet(text)
	fmt.Fprintf(w, "%s\n", out)
}

func build(w http.ResponseWriter, req *http.Request) {
	grow.Build()
	fmt.Fprintf(w, "build\n")
}

func Run() {

	http.HandleFunc("/add", add)
	http.HandleFunc("/extract", extract)
	http.HandleFunc("/build", build)

	http.ListenAndServe(":9093", nil)

}
