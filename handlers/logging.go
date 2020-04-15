package handlers 

import (
   "fmt"
   "net/http"
   "time"
)

func LoggingHandler(response http.ResponseWriter, request * http.Request) {
  format := "%s -- [%s] \"%s %s %s\" %s\n"
  fmt.Printf(format, 
   request.RemoteAddr,
   time.Now().Format(time.RFC1123),
   request.Method,
   request.URL.Path,
   request.Proto,
   request.UserAgent())
}
