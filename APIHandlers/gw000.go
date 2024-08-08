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
	case Consts_Protocol.BATTLE_ATTACK:
		NetResultBattleAttackSucceeded(w, r, DecryptedBody)
	case Consts_Protocol.BATTLE_IN:
		NetResultBattleIn(w, r, DecryptedBody)
	case Consts_Protocol.BATTLE_RESULT:
		NetResultBattleResult(w, r, DecryptedBody)
	default:
		fmt.Println(strconv.Itoa(JSONRequest.PID) + ": Not implemnted!")
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
	//To-Do: Maybe stop all lua events when player is dead?
	err = db_commands.UpdateUserResume(User.ID.String, 1, uint32(UserScan.LuaHash.Int64), Request.TagId, 0)
	if err != nil {
		fmt.Println(err)
	}

	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = uint32(UserScan.LuaHash.Int64)
	//Response.Result = To-Do: Figure this out

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

	//To-Do: Add scan QR codes for this
	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = Consts_LuaHash.Novelty_EV_11_TowerRecords
	//Response.Result = To-Do: Figure this out

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

	//To-Do: Add scan QRs for this
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
	Response.Result = 1 //To-Do: Figure this out

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultBattleIn(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_In_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}
	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		fmt.Println(err)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		fmt.Println(err)
	}
	err = db_commands.UpdateUserSaveBattleId(UserID, Request.BattleId)
	if err != nil {
		fmt.Println(err)
	}

	BattleData := FetchBattleData(Request.BattleId)
	var Response Battle_In_Response
	Response.RES = Consts_RES.SUCCESS
	Response.BGM_ID = BattleData.BGM_ID
	Response.NoiseSymbolPath = BattleData.NoiseSymbolPath
	Response.NoiseID = BattleData.NoiseID
	Response.Badge = BattleData.Badge
	Response.Damage = BattleData.Damage
	Response.BIgnoreInputOrder = BattleData.BIgnoreInputOrder
	if UserSave.BattleBadge1.Valid {
		Response.AAttackItemId = append(Response.AAttackItemId, int(UserSave.BattleBadge1.Int64))
		if UserSave.BattleBadge2.Valid {
			Response.AAttackItemId = append(Response.AAttackItemId, int(UserSave.BattleBadge2.Int64))
			if UserSave.BattleBadge3.Valid {
				Response.AAttackItemId = append(Response.AAttackItemId, int(UserSave.BattleBadge3.Int64))
			}
		}
	} else {
		Response.AAttackItemId = []int{}
	}

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultBattleAttackSucceeded(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_Attack_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}

	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		fmt.Println(err)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		fmt.Println(err)
	}

	if !UserSave.BattleBadge1.Valid {
		err = db_commands.UpdateUserSaveBattleBadge1(UserID, Request.ItemId)
		if err != nil {
			fmt.Println(err)
		}
	} else if !UserSave.BattleBadge2.Valid {
		err = db_commands.UpdateUserSaveBattleBadge2(UserID, Request.ItemId)
		if err != nil {
			fmt.Println(err)
		}
	} else if !UserSave.BattleBadge3.Valid {
		err = db_commands.UpdateUserSaveBattleBadge3(UserID, Request.ItemId)
		if err != nil {
			fmt.Println(err)
		}
	}
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

func NetResultBattleResult(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_Result_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		fmt.Println(err)
	}
	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		fmt.Println(err)
	}
	err = db_commands.UpdateUserSaveNowHP(UserID, Request.NowHp)
	if err != nil {
		fmt.Println(err)
	}
	err = db_commands.EmptyUserSaveBadges(UserID)
	if err != nil {
		fmt.Println(err)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		fmt.Println(err)
	}

	var Response Battle_Result_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = FetchBattleLuaResult(int(UserSave.BattleID.Int64), Request.Succeeded, Request.NowHp)

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}
