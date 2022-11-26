package main

import "net/http"

func main() {
	http.HandleFunc("/", nyaa)
	http.ListenAndServe("127.0.8.1:8080", nil)
}
func nyaa(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("key", "value")
	w.WriteHeader(200)
	w.Header().Add("keykeykeykey", "valuevaluevaluevalue")
	w.Write([]byte("hello\n"))
}
