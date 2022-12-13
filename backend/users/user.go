package users

type User struct {
	UserName string `json:"username"`

	// for extra settings
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}
