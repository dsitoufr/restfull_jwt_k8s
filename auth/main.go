package main 

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    //"github.com/braintree/manners"
   // "github/restfull_jwt_k8s/handlers"
    //"github/restfull_jwt_k8s/user"
    "github/restfull_jwt_k8s/health"
    "github.com/rs/cors"
)

var (
     httpAddr = flag.String("http", ":7500", "Health Service Address")
)

func main() {

    flag.Parse()
    log.Println("starting Auth service...")
    log.Printf("health service listening on %s", *httpAddr)

    errChan := make(chan error)
    signalChan := make(chan os.Signal, 1)

    // Init handle server
    mux := http.NewServeMux()

    // Handle functions
    mux.HandleFunc("/api/healthz", health.HealthzHandler)
    mux.HandleFunc("/api/readiness", health.ReadinessHandler)
    mux.HandleFunc("/api/healthz/status", health.HealthzStatusHandler)
    mux.HandleFunc("/api/readiness/status", health.ReadinessStatusHandler)
    handler := cors.Default().Handler(mux)

    // starting server
    go func() {
       errChan <- http.ListenAndServe(*httpAddr, handler)
    }()

    // Wait signal (ctrl+c) notification
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)


    // check errors
    for {
        select {
            case err := <- errChan :
                 if err != nil {
                     log.Fatal(err)
                 }
            case sig := <- signalChan:
                 log.Println(fmt.Sprintf("captured signal %v  I am exiting...!", sig))
                 health.SetReadinessStatus(http.StatusServiceUnavailable)
                 os.Exit(0)
                 
        }
    }


}