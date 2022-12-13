package general

import (
	"net/http"
)

func Reflect(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL.String())

	url := "https://misskey.meromeromeiro.top"
	// url := strings.ReplaceAll(r.URL.String(), "test.meromeromeiro.top", "misskey.meromeromeiro.top")

	data := Get(url)

	w.Write(data)
}
