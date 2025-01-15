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
	APIAddFriend      = API + "/friends"
)
const (
	APIChatRooms = API + "/ws"
)
const (
	APICreateConversation             = API + "/conversation"
	APIListConversation               = API + "/conversation"
	APIAddConversationParticipants    = API + "/conversation/participants"
	APIDeleteConversationParticipants = API + "/conversation/participants"
	APIListConversationParticipants = API + "/conversation/participants/:conversation_id"

)