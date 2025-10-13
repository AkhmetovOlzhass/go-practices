package main

import (
    "log"
    "go-practice4/internal"
)

func main() {
    app := internal.NewApp()
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
