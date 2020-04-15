package user 

type User struct {
    Username string
    PasswordHash string
    Email string
}

type Users map[string] User

//  Connexion data provided by user username => passwd
type ConnectionCreds struct {
    Username string `json:"username, omitempty"`
    Password  string `json:"password, omitempty"`
}

var DB = Users {
   "bastienfr": {
       Username: "bastienfr",
       // bcrypt  for password "myserceretpassord"
       PasswordHash: "$2a$04$LSljgUDSOwvnI.kbBsZPf.EfvHZTpPOzGpPsxrsYylbLuC5DX.7mW",
       Email: "bastienfr@gmail.com",
   },
}
