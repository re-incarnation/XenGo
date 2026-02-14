package xftypes

//fix
type User struct {
	Option       UserOption        `json:"option"`
	Profile      UserProfile       `json:"profile"`
	Privacy      UserPrivacy       `json:"privacy"`
	Dob          UserDob           `json:"dob"`
	CustomFields map[string]string `json:"custom_fields"`

	Visible               bool   `json:"visible"`
	ActivityVisible       bool   `json:"activity_visible"`
	Timezone              string `json:"timezone"`
	CustomTitle           string `json:"custom_title"`
	Username              string `json:"username"`
	Email                 string `json:"email"`
	UserGroupID           int    `json:"user_group_id"`
	SecondaryGroupIDs     []int  `json:"secondary_group_ids"`
	UserState             string `json:"user_state"`
	IsStaff               bool   `json:"is_staff"`
	MessageCount          int    `json:"message_count"`
	ReactionScore         int    `json:"reaction_score"`
	TrophyPoints          int    `json:"trophy_points"`
	UsernameChangeVisible bool   `json:"username_change_visible"`
	Password              string `json:"password"`
}

type UserOption struct {
	CreationWatchState    string `json:"creation_watch_state"`
	InteractionWatchState string `json:"interaction_watch_state"`
	ContentShowSignature  bool   `json:"content_show_signature"`
	EmailOnConversation   bool   `json:"email_on_conversation"`
	PushOnConversation    bool   `json:"push_on_conversation"`
	ReceiveAdminEmail     bool   `json:"receive_admin_email"`
	ShowDobYear           bool   `json:"show_dob_year"`
	ShowDobDate           bool   `json:"show_dob_date"`
	IsDiscouraged         bool   `json:"is_discouraged"`
}

type UserProfile struct {
	Location  string `json:"location"`
	Website   string `json:"website"`
	About     string `json:"about"`
	Signature string `json:"signature"`
}

type UserPrivacy struct {
	AllowViewProfile              string `json:"allow_view_profile"`
	AllowPostProfile              string `json:"allow_post_profile"`
	AllowReceiveNewsFeed          string `json:"allow_receive_news_feed"`
	AllowSendPersonalConversation string `json:"allow_send_personal_conversation"`
	AllowViewIdentities           string `json:"allow_view_identities"`
}

type UserDob struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type StatusRequest struct {
	Status string
}

type UserUpdateRequest struct {
	User    User
	Option  UserOption
	Profile UserProfile
	Dob     UserDob
	Privacy UserPrivacy
}
