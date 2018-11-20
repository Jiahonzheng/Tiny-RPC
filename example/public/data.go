package public

// User struct
type User struct {
	Name string
	Age  int
}

// ResponseQueryUser struct
type ResponseQueryUser struct {
	User
	Msg string
}
