package APIHandlers

import (
	Consts_Login "FRPGServer/Consts/Login"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Protocol "FRPGServer/Consts/Protocol"
	Consts_RES "FRPGServer/Consts/Res"
	"FRPGServer/SaveData"
	"FRPGServer/Utils"
	db_commands "FRPGServer/db/commands"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Gw000Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gw000 Hit")
	RequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	DecryptedBody := Utils.DESDecrypt(RequestBody)
	fmt.Println(string(DecryptedBody))
	var JSONRequest Generic_Request
	err = json.Unmarshal(DecryptedBody, &JSONRequest)
	if err != nil {
		fmt.Println(err)
	}

	switch JSONRequest.PID {
	case Consts_Protocol.LOGIN_START:
		NetResultPreLogin(w, r)
	case Consts_Protocol.LOGIN_SESSION:
		NetResultSessionUpdate(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.LOGIN_GAME:
		NetResultLogin(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.NG_WORD:
		NetResultNGWord(w, r)
	case Consts_Protocol.CREATE_USER:
		NetResultCreateUser(w, r, DecryptedBody)
	case Consts_Protocol.PARTY_CREATE:
		NetResultPartyCreate(w, r, DecryptedBody)
	case Consts_Protocol.HOME:
		NetResultHome(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.PARTY_START:
		NetResultPartyStart(w, r)
	case Consts_Protocol.EVENT:
		NetResultEvent(w, r, DecryptedBody)
	case Consts_Protocol.QUEST_CHECKED:
		NetResultNewQuest(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.EVENT_CHECK_RESUME:
		NetResultEventCheck(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.SCAN_TAG:
		NetResultScan(w, r, DecryptedBody)
	case Consts_Protocol.SHOP_BENEFIT:
		NetResultShopBenefit(w, r)
	case Consts_Protocol.SHOP_IDENTIFY_START:
		NetResultShopIdentifyStart(w, r)
	case Consts_Protocol.SHOP_IDENTIFY_END:
		NetResultShopIdentifyEnd(w, r)
	default:
		fmt.Println(strconv.Itoa(JSONRequest.PID) + "Not implemnted!")
	}

}

func NetResultPreLogin(w http.ResponseWriter, r *http.Request) {
	var Response Login_Start_Response
	Response.RES = Consts_RES.SUCCESS
	Response.NativeToken = "FHxveQcSfC+xZ9nc6B5euzQX7sUQ44PVTTNizkQRhvHrVDPnytbXqaXjGLBlfkWpHoFR/g8g0+sgK3YTxznENOMBxJnJ7nKSVumBZCTqDcXN/7j0NBAeARwu0kzjh0FdP9j0rXMBcqG="

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultSessionUpdate(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		fmt.Println(err)
	}

	var Response Session_Update_Response
	Response.RES = Consts_RES.SUCCESS
	Response.NewGame = int(User.NewGame.Int64)

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultLogin(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		fmt.Println(err)
	}
	var Player PlayerModel
	Player.ID = User.ID.String
	Player.Name = User.Name.String

	var Response Login_Game_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Result = int(User.Status.Int64)
	Response.LinkUUID = User.UUID.String
	Response.Player = &Player

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultNGWord(w http.ResponseWriter, r *http.Request) {
	var Response NG_Word_Response
	Response.RES = Consts_RES.SUCCESS
	Response.BOK = 1

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultCreateUser(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Create_User_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}

	err = db_commands.SetUserName(Request.TerminalId, Request.PlayerName)
	if err != nil {
		fmt.Println(err)
	}

	err = db_commands.SetUserStatus(Request.TerminalId, Consts_Login.CREATING_PARTY)
	if err != nil {
		fmt.Println(err)
	}

	err = db_commands.SetUserNewGame(Request.TerminalId, 0)
	if err != nil {
		fmt.Println(err)
	}

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultPartyCreate(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Party_Create_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}

	team, err := db_commands.CreateTeam(Request.PartyName)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range Request.AMember {
		err = db_commands.SetUserTeam(user, team.TeamID.String)
		if err != nil {
			fmt.Println(err)
		}
		err = db_commands.InitSaveData(user)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = db_commands.SetUserStatus(Request.TerminalId, Consts_Login.PLAYING)
	if err != nil {
		fmt.Println(err)
	}

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultHome(w http.ResponseWriter, r *http.Request, Did string) {
	IP := os.Getenv("IP_ADDRESS")
	WSPort := os.Getenv("WS_PORT")

	//Testing data
	/*var UserPointer1 PlayerModel
	UserPointer1.ID = "kPjrXd"
	UserPointer1.Name = "Maren"

	var CharacterPointer CharacterModel
	CharacterPointer.CharacterId = 1
	CharacterPointer.ItemId = 1

	var ScanPointer ScanModel
	ScanPointer.ID = 1
	ScanPointer.Type = 3
	ScanPointer.Tag = "GM_shibuya21_ma8_00004"
	//ScanPointer.Height = I have no clue tbh (Apprently used for Type 1 April Tags)
	ScanPointer.BMulti = true

	var EventInfoPointer ResumeDataModel
	EventInfoPointer.BResume = false
	EventInfoPointer.Lua = Consts_LuaHash.Oioi_8F_EV_Shiba_2
	EventInfoPointer.TagId = 0
	EventInfoPointer.ResumeId = 0

	var ItemPointer QuestItemModel
	ItemPointer.Name = "item0"
	ItemPointer.Type = "off"

	var ItemPointer1 QuestItemModel
	ItemPointer1.Name = "noise10_Y"
	ItemPointer1.Type = "off"

	var ItemPointer2 QuestItemModel
	ItemPointer2.Name = "noise11_Y"
	ItemPointer2.Type = "off"

	var ItemPointer3 QuestItemModel
	ItemPointer3.Name = "noise12_Y"
	ItemPointer3.Type = "off"

	var QuestPointer QuestModel
	QuestPointer.ID = 1
	QuestPointer.ItemList = append(QuestPointer.ItemList, ItemPointer)
	QuestPointer.Value = 0

	var GPSPointer GPSModel
	GPSPointer.ID = "4"
	GPSPointer.Name = "badge"
	GPSPointer.PinType = "badge"
	//GPSPointer.PinColor = "y"
	GPSPointer.Latitude = 300
	GPSPointer.Longitude = 200
	//GPSPointer.LuaScript = Consts_LuaHash.Oioi_8F_EV_Deai_0
	GPSPointer.BLocationEvent = 0
	GPSPointer.Quest = &QuestPointer
	GPSPointer.MapType = "maruiMap"
	GPSPointer.MapNo = "8"

	var GPSPointer3 GPSModel
	GPSPointer3.ID = "2"
	GPSPointer3.Name = "y"
	GPSPointer3.PinType = Consts_MapPin.Noise_Elephant
	GPSPointer3.PinColor = "_Y"
	GPSPointer3.Latitude = 0
	GPSPointer3.Longitude = 0
	//GPSPointer3.LuaScript = Consts_LuaHash.Oioi_8F_EV_Deai_0
	GPSPointer3.BLocationEvent = 0
	GPSPointer3.Quest = &QuestPointer
	GPSPointer3.MapType = Consts_MapType.MaruiMap
	GPSPointer3.MapNo = "8"

	var GPSPointer2 GPSModel
	GPSPointer2.ID = "10"
	GPSPointer2.Name = "marui"
	GPSPointer2.PinType = "mapPin"
	//GPSPointer2.PinColor = "pin_sub"
	GPSPointer2.Latitude = 35.660866
	GPSPointer2.Longitude = 139.701024
	//GPSPointer2.LuaScript = Consts_LuaHash.GPS_1
	//GPSPointer2.BLocationEvent = 0
	GPSPointer2.Quest = &QuestPointer
	//GPSPointer2.MapType = Consts_MapType.GPSMap
	//GPSPointer2.MapNo = "0"

	var BuildingPointer BuildingModel
	BuildingPointer.Name = "マルイ"
	BuildingPointer.Prefab = "marui"
	BuildingPointer.Status = "2" //0 is hide. 1 is off. 2 is on.

	var LocalMapPointer LocalMapModel
	LocalMapPointer.Floor = 8
	LocalMapPointer.Name = "maruiMap"

	var Response Home_Response
	Response.RES = Consts_RES.SUCCESS
	//Response.ACharacter = append(Response.ACharacter, CharacterPointer)
	Response.AScan = append(Response.AScan, ScanPointer)
	Response.ARemoveScan = []string{}
	Response.AGPS = append(Response.AGPS, GPSPointer2, GPSPointer)
	Response.ARemoveGPS = []string{}
	Response.AOnceEvent = []uint32{Consts_LuaHash.Introduction}
	Response.ARemoveOnceEvent = []uint32{}
	Response.ABuildings = append(Response.ABuildings, BuildingPointer)
	Response.NowHp = 100
	Response.MaxHp = 100
	Response.ColorId = 3 //1 is Red, 2 is Blue, 3 is Yellow
	Response.Quest = &QuestPointer
	Response.BNewQuest = false
	Response.AItemList = make([]int, 0)
	//Response.LocalMap = &LocalMapPointer
	Response.WebSocketServer = "ws://" + IP + WSPort //Only used when there is more than one player in a team
	Response.RoomId = "J24YUe"
	Response.TeamId = "J24YUe"
	Response.TeamName = "MarenTeam"
	Response.ATeamUser = append(Response.ATeamUser, UserPointer1)
	Response.EventInfo = &EventInfoPointer
	Response.BIntro = true*/

	User, err := db_commands.GetUser(Did)
	if err != nil {
		fmt.Println(err)
	}
	UserID := User.ID.String

	var Response Home_Response

	var Player PlayerModel
	Player.ID = UserID
	Player.Name = User.Name.String
	Response.ATeamUser = append(Response.ATeamUser, Player)

	saveData, intro, err := fetchSaveData(UserID)
	if err != nil {
		fmt.Println(err)
	}

	resumeData, err := db_commands.GetUserResume(UserID)
	if err != nil {
		fmt.Println(err)
	}
	var EventInfoPointer ResumeDataModel
	if resumeData.BResume.Int64 == 0 {
		EventInfoPointer.BResume = false
	} else {
		EventInfoPointer.BResume = true
	}
	EventInfoPointer.Lua = uint32(resumeData.LuaHash.Int64)
	EventInfoPointer.TagId = int32(resumeData.TagID.Int64)
	EventInfoPointer.ResumeId = int32(resumeData.ResumeID.Int64)

	Response.RES = Consts_RES.SUCCESS
	Response.ACharacter = saveData.ACharacter
	Response.AScan = saveData.AScan
	Response.ARemoveScan = saveData.ARemoveScan
	Response.AGPS = saveData.AGPS
	Response.ARemoveGPS = saveData.ARemoveGPS
	Response.AOnceEvent = saveData.AOnceEvent
	Response.ARemoveOnceEvent = saveData.ARemoveOnceEvent
	Response.ABuildings = saveData.ABuildings
	Response.NowHp = saveData.NowHp
	Response.MaxHp = saveData.MaxHp
	Response.ColorId = saveData.ColorId
	Response.Quest = saveData.Quest
	Response.AItemList = saveData.AItemList
	Response.LocalMap = saveData.LocalMap
	Response.WebSocketServer = "ws://" + IP + WSPort //Only used when there is more than one player in a team
	Response.RoomId = "J24YUe"
	Response.TeamId = "J24YUe"
	Response.TeamName = "MarenTeam"
	Response.EventInfo = &EventInfoPointer
	Response.BIntro = intro

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultPartyStart(w http.ResponseWriter, r *http.Request) {
	//IP := os.Getenv("IP_ADDRESS")
	//WSPort := os.Getenv("WS_PORT")

	var Response Party_Start_Response
	Response.RES = Consts_RES.MAINTENANCE
	//Response.PartyId = "cHifjK"
	//Response.WebSocketServer = "ws://"+IP+WSPort

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEvent(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Event_End_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}

	User, err := db_commands.GetUser(Request.TerminalId)
	if err != nil {
		fmt.Println(err)
	}
	UserID := User.ID.String

	for _, LuaHash := range Request.ALua {
		err = SaveData.ProccessLua(UserID, LuaHash)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = db_commands.UpdateUserResumeBool(UserID, 0)
	if err != nil {
		fmt.Println(err)
	}

	saveData, _, err := fetchSaveData(UserID)
	if err != nil {
		fmt.Println(err)
	}

	var Response Event_Response
	Response.RES = Consts_RES.SUCCESS
	Response.ACharacter = saveData.ACharacter
	Response.AScan = saveData.AScan
	Response.ARemoveScan = saveData.ARemoveScan
	Response.AGPS = saveData.AGPS
	Response.ARemoveGPS = saveData.ARemoveGPS
	Response.AOnceEvent = saveData.AOnceEvent
	Response.ARemoveOnceEvent = saveData.ARemoveOnceEvent
	Response.ABuildings = saveData.ABuildings
	Response.NowHp = saveData.NowHp
	Response.MaxHp = saveData.MaxHp
	Response.ColorId = saveData.ColorId
	Response.Quest = saveData.Quest
	Response.BNewQuest = saveData.BNewQuest
	Response.AItemList = saveData.AItemList
	Response.LocalMap = saveData.LocalMap

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEventCheck(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		fmt.Println(err)
	}

	resumeData, err := db_commands.GetUserResume(User.ID.String)
	var EventInfoPointer ResumeDataModel
	if resumeData.BResume.Int64 == 0 {
		EventInfoPointer.BResume = false
	} else {
		EventInfoPointer.BResume = true
	}
	EventInfoPointer.Lua = uint32(resumeData.LuaHash.Int64)
	EventInfoPointer.TagId = int32(resumeData.TagID.Int64)
	EventInfoPointer.ResumeId = int32(resumeData.ResumeID.Int64)

	var Response Event_Check_Resume_Response
	Response.RES = Consts_RES.SUCCESS
	Response.EventInfo = &EventInfoPointer

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultScan(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Scan_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}

	User, err := db_commands.GetUser(Request.TerminalId)
	if err != nil {
		fmt.Println(err)
	}

	UserScan, err := db_commands.GetUserScan(User.ID.String, Request.TagId)
	if err != nil {
		fmt.Println(err)
	}

	//To-Do: Figure out when ResumeID is needed
	err = db_commands.UpdateUserResume(User.ID.String, 1, uint32(UserScan.LuaHash.Int64), Request.TagId, 0)
	if err != nil {
		fmt.Println(err)
	}

	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = uint32(UserScan.LuaHash.Int64)
	//Response.Result = No Clue

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopBenefit(w http.ResponseWriter, r *http.Request) {

	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = Consts_LuaHash.EV_Novelty_11_TowerRecords
	//Response.Result = No Clue

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopIdentifyStart(w http.ResponseWriter, r *http.Request) {

	var Response Shop_Identify_Start_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Message = "This is a test"
	//Response.Type = 0 //1: Badge, 2: Discount, 3: Bonus Card, Anything Else: Already got
	//Response.Result = No Clue

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopIdentifyEnd(w http.ResponseWriter, r *http.Request) {

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultNewQuest(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		fmt.Println(err)
	}
	err = db_commands.UpdateUserSaveNewQuest(User.ID.String, 0)
	if err != nil {
		fmt.Println(err)
	}

	var Response New_Quest_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Result = 1 //idk tbh

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}
