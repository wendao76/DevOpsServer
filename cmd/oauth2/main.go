package main

import (
    "flag"
    "go_web/internal/oauth2/http"
)

func main() {
    flag.Parse()
    http.New()
}
