package handlers

import(
    "net/http"
    "encoding/json"
    "log"
)


type HelloResponse struct {
    Message string `json:"message"`
}

func HelloHandler(response http.ResponseWriter, request *http.Request ) {
   var helloResp HelloResponse

   // Read message from request body
   err := json.NewDecoder(request.Body).Decode(&helloResp)
   if err != nil {
       log.Println("Error In decoder message body")
       response.WriteHeader(http.StatusBadRequest)
   }

   // write back message to response stream

   json.NewEncoder(response).Encode(helloResp)
   return
   
}