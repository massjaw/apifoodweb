package req

type UsersRequestAPI struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}
