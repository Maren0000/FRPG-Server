package APIHandlers

type Generic_Request struct {
	PID        int    `json:"pid"`        //Protocol ID
	Ver        int    `json:"ver"`        //1 is 1.0.0, 2 is 1.1.0, 3 is 1.2.1
	DeviceType int    `json:"deviceType"` //1 is iOS, 2 is Android
	TerminalId string `json:"terminalId"` //Device ID
	CC         int    `json:"CC"`         //Packet Counter
}

type Create_User_Request struct {
	Generic_Request
	PlayerName string `json:"playerName"`
}

type Party_Create_Request struct {
	Generic_Request
	PartyName string   `json:"partyName"`
	AMember   []string `json:"aMember"`
}

type Scan_Request struct {
	Generic_Request
	TagId int `json:"tagId"`
}

type Event_End_Request struct {
	Generic_Request
	ALua []uint32 `json:"aLua"`
}

type Battle_In_Request struct {
	Generic_Request
	BattleId int `json:"battleId"`
}

type Battle_Attack_Request struct {
	Generic_Request
	ItemId int `json:"ItemId"`
}

type Battle_Result_Request struct {
	Generic_Request
	Succeeded int `json:"succeeded"`
	NowHp     int `json:"nowHp"`
}

type Generic_Response struct {
	RES int `json:"res"`
}

type Get_URL_Response struct {
	Generic_Response
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
	Generic_Response
	NativeToken string `json:"nativeToken"`
}

type Session_Update_Response struct {
	Generic_Response
	NewGame int `json:"newGame"`
}

type Login_Game_Response struct {
	Generic_Response
	Result   int          `json:"result"`
	LinkUUID string       `json:"linkUUID"`
	Player   *PlayerModel `json:"player"`
}

type NG_Word_Response struct {
	Generic_Response
	BOK int `json:"bOK"`
}

type Home_Response struct {
	Generic_Response
	PlayerDataModel
	WebSocketServer string           `json:"webSockesServer"`
	RoomId          string           `json:"roomId"`
	TeamId          string           `json:"teamId"`
	TeamName        string           `json:"teamName"`
	ATeamUser       []PlayerModel    `json:"aTeamUser"`
	EventInfo       *ResumeDataModel `json:"eventInfo"`
	BIntro          bool             `json:"bIntro"`
}

type Event_Response struct {
	Generic_Response
	PlayerDataModel
}

type Party_Start_Response struct {
	Generic_Response
	PartyId         string `json:"partyId"`
	WebSocketServer string `json:"webSockesServer"`
}

type Event_Check_Resume_Response struct {
	Generic_Response
	EventInfo *ResumeDataModel `json:"eventInfo"`
}

type Scan_Response struct {
	Generic_Response
	Lua    uint32 `json:"lua"`
	Result uint32 `json:"result"`
}

type New_Quest_Response struct {
	Generic_Response
	Result uint32 `json:"result"`
}

type Shop_Identify_Start_Response struct {
	Generic_Response
	Message string `json:"message"`
	Result  uint32 `json:"result"`
	Type    int    `json:"type"`
}

type Battle_In_Response struct {
	Generic_Response
	BGM_ID            int    `json:"bgm_id"`
	NoiseSymbolPath   string `json:"noizeSymbolPath"`
	NoiseID           int    `json:"noizeId"`
	Badge             string `json:"badge"`
	Damage            int    `json:"damage"`
	BIgnoreInputOrder int    `json:"bIgnoreInputOrder"`
	AAttackItemId     []int  `json:"aAttackItemId"`
}

type Battle_Result_Response struct {
	Generic_Response
	Lua uint32 `json:"lua,omitempty"`
}
