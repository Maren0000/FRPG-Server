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
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NG_Word_Response struct {
	RES int `json:"res"`
	BOK int `json:"bOK"`
}

type Generic_Response struct {
	RES int `json:"res"`
}

type Home_Response struct {
	RES              int              `json:"res"`
	ACharacter       []CharacterModel `json:"aCharacter"`
	AScan            []ScanModel      `json:"aScan"`
	ARemoveScan      []string         `json:"aRemoveScan"`
	AGPS             []GPSModel       `json:"aGPS"`
	ARemoveGPS       []string         `json:"aRemoveGPS"`
	AOnceEvent       []uint32         `json:"aOnceEvent"`
	ARemoveOnceEvent []uint32         `json:"aRemoveOnceEvent"`
	ABuildings       []BuildingModel  `json:"aBuildings"`
	NowHp            int              `json:"nowHp"`
	MaxHp            int              `json:"maxHp"`
	ColorId          int              `json:"colorId"`
	Quest            *QuestModel      `json:"quest"`
	BNewQuest        bool             `json:"bNewQuest"`
	AItemList        []int            `json:"aItemList"`
	LocalMap         *LocalMapModel   `json:"localMap"`
	WebSocketServer  string           `json:"webSockesServer,omitempty"`
	RoomId           string           `json:"roomId,omitempty"`
	TeamId           string           `json:"teamId,omitempty"`
	TeamName         string           `json:"teamName,omitempty"`
	ATeamUser        []UserModel      `json:"aTeamUser"`
	EventInfo        *ResumeDataModel `json:"eventInfo"`
	BIntro           bool             `json:"bIntro"`
}

type Event_Response struct {
	RES              int              `json:"res"`
	ACharacter       []CharacterModel `json:"aCharacter"`
	AScan            []ScanModel      `json:"aScan"`
	ARemoveScan      []string         `json:"aRemoveScan"`
	AGPS             []GPSModel       `json:"aGPS"`
	ARemoveGPS       []string         `json:"aRemoveGPS"`
	AOnceEvent       []uint32         `json:"aOnceEvent"`
	ARemoveOnceEvent []uint32         `json:"aRemoveOnceEvent"`
	ABuildings       []BuildingModel  `json:"aBuildings"`
	NowHp            int              `json:"nowHp"`
	MaxHp            int              `json:"maxHp"`
	ColorId          int              `json:"colorId"`
	Quest            *QuestModel      `json:"quest"`
	BNewQuest        bool             `json:"bNewQuest"`
	AItemList        []int            `json:"aItemList"`
	LocalMap         *LocalMapModel   `json:"localMap"`
}

type ResumeDataModel struct {
	BResume  bool   `json:"bResume,omitempty"`
	Lua      uint32 `json:"lua,omitempty"`
	TagId    int32  `json:"tagId,omitempty"`
	ResumeId int32  `json:"resumeId,omitempty"`
}

type CharacterModel struct {
	CharacterId int `json:"characterId"`
	ItemId      int `json:"itemId"`
}

type ScanModel struct {
	ID     int     `json:"id"`
	Type   int     `json:"type"`
	Tag    string  `json:"tag"`
	Height float32 `json:"height"`
	BMulti bool    `json:"bMulti"`
}

type GPSModel struct {
	ID             int         `json:"int"`
	Name           string      `json:"name"`
	PinType        string      `json:"pinType"`
	PinColor       string      `json:"pinColor"`
	Latitude       float64     `json:"latitude"`
	Longitude      float64     `json:"longitude"`
	LuaScript      uint32      `json:"LuaScript"`
	BLocationEvent int         `json:"bLocationEvent"`
	Quest          *QuestModel `json:"quest"`
	MapType        string      `json:"mapType"`
	MapNo          string      `json:"mapNo"`
}

type BuildingModel struct {
	Prefab string `json:"prefab"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type QuestModel struct {
	ID       int              `json:"id,omitempty"`
	ItemList []QuestItemModel `json:"itemList,omitempty"`
	Value    int              `json:"value,omitempty"`
}

type QuestItemModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type LocalMapModel struct {
	Name  string `json:"name,omitempty"`
	Floor int    `json:"floor,omitempty"`
}

type Party_Start_Response struct {
	RES             int    `json:"res"`
	PartyId         string `json:"partyId"`
	WebSocketServer string `json:"webSockesServer"`
}

type Event_Check_Resume_Response struct {
	RES       int              `json:"res"`
	EventInfo *ResumeDataModel `json:"eventInfo"`
}
