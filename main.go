package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(
            w,
            "%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
            "<html>",
              "<head>",
                "<meta name=viewport content='width=device-width, initial-scale=1.0'>",
                "<style>",
                  ".c{display:flex;flex-flow:row wrap;justify-content:space-around;}",
                  ".i{width:200px;height:150px;background:tomato;margin:10px;font-weight:bold;font-size:3em;text-align:center;line-height:150px;}",
                  "a:visited,a:link{color:white;text-decoration:none;}",
                "</style>",
              "</head>",
              "<body>",
                "<div class=c>",
                  "<a href=/0><div class=i>0/0</div></a>",
                  "<a href=/1><div class=i>0/1</div></a>",
                  "<a href=/2><div class=i>1/0</div></a>",
                  "<a href=/3><div class=i>1/1</div></a>",
                "</div>",
              "</body>",
            "</html>",
        )
    })

    http.HandleFunc("/0", func(w http.ResponseWriter, r *http.Request) {
        go http.Get(fmt.Sprintf("http://%s/1", os.Getenv("HOST_A")))
        go http.Get(fmt.Sprintf("http://%s/1", os.Getenv("HOST_B")))
        http.Redirect(w, r, "/", 307)
    })

    http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
        go http.Get(fmt.Sprintf("http://%s/1", os.Getenv("HOST_A")))
        go http.Get(fmt.Sprintf("http://%s/", os.Getenv("HOST_B")))
        http.Redirect(w, r, "/", 307)
    })

    http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
        go http.Get(fmt.Sprintf("http://%s/", os.Getenv("HOST_A")))
        go http.Get(fmt.Sprintf("http://%s/1", os.Getenv("HOST_B")))
        http.Redirect(w, r, "/", 307)
    })

    http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
        go http.Get(fmt.Sprintf("http://%s/", os.Getenv("HOST_A")))
        go http.Get(fmt.Sprintf("http://%s/", os.Getenv("HOST_B")))
        http.Redirect(w, r, "/", 307)
    })

    log.Fatal(http.ListenAndServe(os.Getenv("LISTEN"), nil))
}

