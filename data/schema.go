package data

//go:generate reform

//reform:users
type User struct {
	ID string `reform:"id,pk"`
}
