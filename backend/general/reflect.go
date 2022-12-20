package general

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
)

var Client *http.Client

func Proxy(w http.ResponseWriter, r *http.Request) {
	// color.Green(fmt.Sprintln(r))

	req := newRequestFromRequest(r)

	color.Green(fmt.Sprintln(req))

	if resp, err := Client.Do(req); err == nil {
		writeResposeToResponseWriter(w, resp)
	} else {
		color.Red(err.Error())
	}

}

func Reflect(w http.ResponseWriter, r *http.Request) {
	color.Green(fmt.Sprintln(r))

	url := "https://misskey.meromeromeiro.top" + r.URL.String()
	// url := strings.ReplaceAll(r.URL.String(), "test.meromeromeiro.top", "misskey.meromeromeiro.top")
	color.Cyan("Get " + url)

	data := Get(url)

	w.Write(data)
}

func newRequestFromRequest(r *http.Request) *http.Request {
	color.Yellow("https://misskey.meromeromeiro.top" + r.RequestURI)
	text, err := io.ReadAll(r.Body)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	newTxt := strings.ReplaceAll(string(text), "misskey.meromeromeiro.top", "test.meromeromeiro.top")

	color.Yellow(newTxt)

	req, err := http.NewRequest(r.Method, "https://misskey.meromeromeiro.top"+r.RequestURI, bytes.NewReader([]byte(newTxt)))

	if err != nil {
		// handle err
		color.Red(err.Error())
		return req
	}
	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	return req
}

func writeResposeToResponseWriter(w http.ResponseWriter, r *http.Response) {
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.Header().Del("Content-Length")
	// w.WriteHeader(r.StatusCode)
	// io.Copy(w, r.Body)

	body := getPlainTextReader(r.Body, r.Header.Get("Content-Encoding"))
	// w.Header().Set("Content-Encoding", "plain")
	text, err := io.ReadAll(body)
	if err != nil {
		color.Red(err.Error())
		return
	}

	newTxt := strings.ReplaceAll(string(text), "misskey.meromeromeiro.top", "test.meromeromeiro.top")

	// color.White(newTxt)

	w.WriteHeader(r.StatusCode)
	w.Write([]byte(newTxt))
}

func getPlainTextReader(body io.ReadCloser, encoding string) io.ReadCloser {
	switch encoding {
	case "gzip":
		reader, err := gzip.NewReader(body)
		if err != nil {
			log.Fatal("error decoding gzip response", reader)
		}
		return reader
	case "br":
		reader := brotli.NewReader(body)
		if reader == nil {
			log.Fatal("error decoding br response", reader)
		}
		return io.NopCloser(reader)
	default:
		return body
	}
}
