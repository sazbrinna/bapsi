package auth

type inputResources struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Nama     string `json:"nama"`
}
