package cashmachine

type UserInfo struct {
	Username string `bson:"username,omitempty"`
	Pin      int    `bson:"pin,omitempty"`
}
