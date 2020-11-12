package types

// User user object
type User struct {
	Handle    string `json:"handle" form:"handle"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Rating    int32  `json:"rating" form:"rating"`
}
