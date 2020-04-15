package handlers

import(
    "encoding/json"
    "net/http"
    "time"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
    "github/restfull_jwt_k8s/user"
    "log"
)

type LoginResponse struct {
    Token string `json:"token, omitempty"`
}


type LoginHandler struct {
    //secret string
    users user.Users
}


func LoginConnectHandler(response http.ResponseWriter, request *http.Request) {
  // Connexion data provided by user username => passwd
var (
       connCreds user.ConnectionCreds
       perr error
    )

loginHandler := &LoginHandler{
    users: user.DB,
}

// Get username and password from resquest body
perr = json.NewDecoder(request.Body).Decode(&connCreds)
if perr != nil {
    log.Println("Error in decoder message body")
    response.WriteHeader(http.StatusBadRequest)
    return
}

// Creds provided
username := connCreds.Username
password := connCreds.Password

// Database email
email := loginHandler.users[username].Email

// Compare DB and provided password
perr = bcrypt.CompareHashAndPassword([]byte(loginHandler.users[username].PasswordHash), []byte(password))
if perr != nil {
    log.Println("Error Passwords not matched")
    response.WriteHeader(http.StatusBadRequest)
    return
} 

// Token creation
claims := JWTData{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour).Unix(),
                Issuer: "auth.service",
                IssuedAt: time.Now().Unix(),
                Subject: username,
			},

		   CustomClaims: map[string]string{
           "username": username,
           "email": email,
           },

       }
		
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString, err := token.SignedString([]byte(SECRET))
if err != nil {
			log.Println(err)
			http.Error(response, "failed to get token string", http.StatusUnauthorized)
		}

// Put token in Header Autorization key
// Does not work for the moment do it manually
request.Header.Set("Authorization", tokenString)

// Display token
loginResponse := &LoginResponse{Token: tokenString}
json.NewEncoder(response).Encode(loginResponse)


}

//
//
//  postman 

//   input:
//  {"username": "bastienfr", "password":"myserceretpassord"}
//
//
//   output:  token
//    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODY5NDU1OTksImlhdCI6MTU4Njk0MTk5OSwiaXNzIjoiYXV0aC5zZXJ2aWNlIiwic3ViIjoiYmFzdGllbmZyIiwiY3VzdG9tIjp7ImVtYWlsIjoiYmFzdGllbmZyQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYmFzdGllbmZyIn19.v2JxGeGeeExsAc6tKYyPkpFFFvlrc-niFLKlBZbJWd0"
//
//   
//