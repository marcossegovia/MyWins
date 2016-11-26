package main

import (
        "fmt"
        "log"
        "net/http"
        "os"

        "github.com/MarcosSegovia/MyWins/src/wins/config"
)

func main() {

        appEnvironment := os.Getenv("app_env")
        if appEnvironment == "prod" {
                fmt.Println("Executing Prod Configuration.")
                config.SetProdEnvironment()
        } else {
                fmt.Println("Executing Dev Configuration.")
                config.SetDevEnvironment()
        }

        BootstrapClient()

        serverRouter := NewServerRouter()
        clientRouter := NewClientRouter()

        go http.ListenAndServe(":8081", clientRouter)
        e := http.ListenAndServe(":8080", serverRouter)

        if e != nil {

                log.Fatal(e)
        }
}
