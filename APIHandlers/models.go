package APIHandlers

type PlayerModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlayerDataModel struct {
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
	BResume  bool   `json:"bResume"`
	Lua      uint32 `json:"lua"`
	TagId    int32  `json:"tagId"`
	ResumeId int32  `json:"resumeId"`
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
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	PinType        string      `json:"pinType"`
	PinColor       string      `json:"pinColor"`
	Latitude       float64     `json:"latitude"`
	Longitude      float64     `json:"longitude"`
	LuaScript      uint32      `json:"luaScript"`
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
	ID       int              `json:"id"`
	ItemList []QuestItemModel `json:"itemList"`
	Value    int              `json:"value"`
}

type QuestItemModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type LocalMapModel struct {
	Name  string `json:"name"`
	Floor int    `json:"floor"`
}
