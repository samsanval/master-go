package main

import "net/http"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)

}
func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "./index.html")
}
