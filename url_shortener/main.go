package main

import (
    "fmt"
    "net/http"
    "github.com/gophercises/urlshort"
    yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        if dest, ok := pathToUrls[path]; ok {
            http.Redirect(w, r, dest, http.StatusFound)
            return
        }
        fallback.ServeHTTP(w, r)
    }
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
    var pathUrls []pathUrl
    err := yaml.Unmarshal(yamlBytes, &pathUrls)
    if err != nil {
        return nil, err
    }
    pathsToUrls := make(map[string]string)
    for _, pu := range pathUrls {
    pathsToUrls[pu.Path] = pu.URL
    }
    return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
    Path string `yaml:"path"`
    URL string `yaml:"url"`
}

func main() {
    mux := defaultMux()

    pathsToUrls := map[string]string{
        "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
        "/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
    }
    mapHandler := urlshort.MapHandler(pathsToUrls, mux)

    yaml := `
        - path: /urlshort
          url: https://github.com/gophercises/urlshort
        - path: /urlshort-final
          url: https://github.com/gophercises/urlshort/tree/solution
        `

    yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
    if err != nil {
        panic(err)
    }
    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
