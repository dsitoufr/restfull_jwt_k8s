package main

import (
    "github/restfull_jwt_k8s/handlers"
    "flag"
    "log"
    "net/http"
    "github.com/rs/cors"
    "os"
    "os/signal"
    "syscall"
)

var (
        httpAddr = flag.String("http",":7000", "HTTP service address.")
       // healthAddr = flag.String("health", ":7500", "Health service address.")
   ) 

func main() {
    errChan := make(chan error, 10) 
    signalChan := make(chan os.Signal, 1)

    flag.Parse()
    log.Println("Starting server...")
    //log.Printf("Health service listening on %s", *healthAddr)
    log.Printf("HTTP servive listening on %s", *httpAddr)

    // Init Handle Server  
    mux := http.NewServeMux()

    // Handle functions
    mux.HandleFunc("/api", handlers.HelloHandler)
    mux.HandleFunc("/api/version", handlers.VersionHandler)
    mux.HandleFunc("/api/jwt", handlers.JWTTokenHandler)
    mux.HandleFunc("/api/login", handlers.LoginConnectHandler)
    mux.HandleFunc("/api/logging", handlers.LoggingHandler)
    
    handler := cors.Default().Handler(mux)

    // Starting server
    log.Println("starting servers...")
    go func() {
        errChan <- http.ListenAndServe(*httpAddr, handler)
    }()

     // Waiting signal ctr + c
     signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
     
    // Check errors
    for {
                    select {
                            case err := <- errChan:
                                if err != nil {
                                     log.Fatal(err)
                                 }   

                           case  sig := <- signalChan:  
                           // Exit when cath signal
                           log.Printf("Get Os signal : %v. Exiting program...!", sig)
                           os.Exit(0)
                             
                     }
    }
    
  

}