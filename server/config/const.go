package config

const (
	API = "/api"
)

const (
	APILogin    = API + "/login"
	APIRegister = API + "/register"
)

const (
	APIUpdateUser = API + "/users/:id"
	APIGetUser    = API + "/users/:id"
	APIListUsers  = API + "/users"
	APIDeleteUser = API + "/users/:id"
)
const (
	APIGetUserFriends = API + "/users/:id/friends"
	APISearchFriends  = API + "/friends/:username/friends"
	APIDeleteFriend   = API + "/friends"
	APIAddFriend   = API + "/friends"
)
