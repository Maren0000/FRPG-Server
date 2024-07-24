package APIHandlers

type Request struct {
	PID        int    `json:"pid"`        //Protocol ID
	Ver        int    `json:"ver"`        //1 is 1.0.0, 2 is 1.1.0, 3 is 1.2.1
	DeviceType int    `json:"deviceType"` //1 is iOS, 2 is Android
	TerminalId string `json:"terminalId"` //Device ID
	CC         int    `json:"CC"`         //Packet Counter
}

type Get_URL_Response struct {
	RES          int    `json:"res"`
	API          string `json:"api"`
	NSC          string `json:"nsc"`
	MainteURL    string `json:"mainteUrl"`
	TermsURL     string `json:"termsUrl"`
	PrivacyURL   string `json:"PrivacyUrl"`
	LicenseURL   string `json:"licenseUrl"`
	CommerceURL  string `json:"commerceUrl"`
	TicketURL    string `json:"ticketUrl"`
	StoreURL     string `json:"storeUrl"`
	EventEndDate int    `json:"eventEndDate"`
}

type Login_Start_Response struct {
	RES         int    `json:"res"`
	NativeToken string `json:"nativeToken"`
}

type Session_Update_Response struct {
	RES     int `json:"res"`
	NewGame int `json:"newGame"`
}

type Login_Game_Response struct {
	RES      int          `json:"res"`
	Result   int          `json:"result"`
	LinkUUID string       `json:"linkUUID"`
	Player   *PlayerModel `json:"player,omitempty"`
}

type PlayerModel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserModel struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type NG_Word_Response struct {
	RES int `json:"res"`
	BOK int `json:"bOK"`
}

type Generic_Response struct {
	RES int `json:"res"`
}

type Home_Response struct {
	RES             int              `json:"res"`
	WebSocketServer string           `json:"webSockesServer,omitempty"`
	RoomId          string           `json:"roomId,omitempty"`
	TeamId          string           `json:"teamId,omitempty"`
	TeamName        string           `json:"teamName,omitempty"`
	ATeamUser       []UserModel      `json:"user"`
	EventInfo       *ResumeDataModel `json:"eventInfo"`
	BIntro          bool             `json:"bIntro"`
}

type ResumeDataModel struct {
	BResume  bool   `json:"bResume"`
	Lua      uint32 `json:"lua"`
	TagId    int32  `json:"tagId"`
	ResumeId int32  `json:"resumeId"`
}

type Party_Start_Response struct {
	RES             int    `json:"res"`
	PartyId         string `json:"partyId"`
	WebSocketServer string `json:"webSockesServer"`
}
