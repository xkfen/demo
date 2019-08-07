package main
import (
    "fmt"
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func main() {
    http.HandleFunc("/", sayHello)
    log.Println("server start...")
    if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
        log.Fatal("server start err:" + err.Error())
    }
}

