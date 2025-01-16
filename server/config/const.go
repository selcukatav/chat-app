package config


//define endpoints

//base
const (
	API = ""
)

//Auth
const (
	APILogin    = API + "/login"
	APIRegister = API + "/register"
)

//User
const (
	APIUpdateUser = API + "/users/:id"
	APIGetUser    = API + "/users/:id"
	APIListUsers  = API + "/users"
	APIDeleteUser = API + "/users/:id"
)

//Friends
const (
	APIGetUserFriends = API + "/users/:id/friends"
	APISearchFriends  = API + "/friends/:username/friends"
	APIDeleteFriend   = API + "/friends"
	APIAddFriend      = API + "/friends"
)

//Websocket
const (
	APIConversationRoom = API + "/ws/conversation"
)

//Conversation
const (
	APICreateConversation             = API + "/conversation"
	APIListConversation               = API + "/conversation"
	APIAddConversationParticipants    = API + "/conversation/participants"
	APIDeleteConversationParticipants = API + "/conversation/participants"
	APIListConversationParticipants = API + "/conversation/participants/:conversation_id"
	APIListUserConversations = API + "/user/conversations/:user_id"

)