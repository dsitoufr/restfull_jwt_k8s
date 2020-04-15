package handlers

import(
   "net/http"
   "encoding/json"
   "github.com/dgrijalva/jwt-go"
   "errors"
   "log"
   "strings"
)


type TokenResponse struct {
    Message string `json:"message, omitempty"`
}

// Standards Claims
type JWTData struct {
     jwt.StandardClaims
     CustomClaims map[string]string `json:"custom,omitempty"`    
}



func JWTTokenHandler(response http.ResponseWriter, request *http.Request) {
  

    // Read token from authorization header
    authToken := request.Header.Get("Authorization")
    authStr := strings.Split(authToken, " ")

    if len(authStr) < 1 {
        log.Println("Authorization header is invalid")
        http.Error(response,"Authorization header is invalid", http.StatusUnauthorized)
    }

    // Parse token 
    jwToken := authStr[0]
    _, err := jwt.ParseWithClaims(jwToken, &JWTData{}, func(token *jwt.Token)(interface{}, error){
       if token.Method != jwt.SigningMethodHS256 {
           return nil, errors.New("Invalid signing method")
        }

        return []byte(SECRET), nil
    })

    if err != nil {
        log.Println(authStr)
        log.Println(err)
        http.Error(response, "failed Parsing token", http.StatusUnauthorized)
        return
    }

    tokenResponse :=  &TokenResponse{Message: jwToken}
    json.NewEncoder(response).Encode(tokenResponse)
}

//
// postman
// 1 - Create Authorization key in Header put the token:
// 2 - Put the token provided by /api/login
//  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODY5NDU1OTksImlhdCI6MTU4Njk0MTk5OSwiaXNzIjoiYXV0aC5zZXJ2aWNlIiwic3ViIjoiYmFzdGllbmZyIiwiY3VzdG9tIjp7ImVtYWlsIjoiYmFzdGllbmZyQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYmFzdGllbmZyIn19.v2JxGeGeeExsAc6tKYyPkpFFFvlrc-niFLKlBZbJWd0"
//
//

