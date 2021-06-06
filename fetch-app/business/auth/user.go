package auth

//User domain user who access system
type User struct {
	ID    int
	Phone string
	Name  string
	Role  int
}
