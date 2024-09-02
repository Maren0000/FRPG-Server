package APIHandlers

import (
	Consts_Login "FRPGServer/Consts/Login"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Protocol "FRPGServer/Consts/Protocol"
	Consts_RES "FRPGServer/Consts/Res"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	"FRPGServer/DataHandlers"
	"FRPGServer/Models"
	"FRPGServer/Utils"
	db_commands "FRPGServer/db/commands"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func Gw000Handler(w http.ResponseWriter, r *http.Request) {

	RequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Empty Body",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	DecryptedBody, err := Utils.DESDecrypt(RequestBody)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Decrypt Failed",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Request: " + string(DecryptedBody))
	var JSONRequest Generic_Request
	err = json.Unmarshal(DecryptedBody, &JSONRequest)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("JSON unmarshal failed",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Info("GW000 Hit: " + strconv.Itoa(JSONRequest.PID) + " for user " + JSONRequest.TerminalId)

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
	case Consts_Protocol.EVENT_SAVE_RESUME:
		NetResultEventSaveResume(w, r, DecryptedBody)
	case Consts_Protocol.QUEST_CHECKED:
		NetResultNewQuest(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.EVENT_CHECK_RESUME:
		NetResultEventCheck(w, r, JSONRequest.TerminalId)
	case Consts_Protocol.SCAN_TAG:
		NetResultScan(w, r, DecryptedBody)
	case Consts_Protocol.SHOP_BENEFIT:
		NetResultShopBenefit(w, r, DecryptedBody)
	case Consts_Protocol.SHOP_IDENTIFY_START:
		NetResultShopIdentifyStart(w, r, DecryptedBody)
	case Consts_Protocol.SHOP_IDENTIFY_END:
		NetResultShopIdentifyEnd(w, r)
	case Consts_Protocol.BATTLE_ATTACK:
		NetResultBattleAttackSucceeded(w, r, DecryptedBody)
	case Consts_Protocol.BATTLE_IN:
		NetResultBattleIn(w, r, DecryptedBody)
	case Consts_Protocol.BATTLE_RESULT:
		NetResultBattleResult(w, r, DecryptedBody)
	default:
		slog.Error(strconv.Itoa(JSONRequest.PID) + ": Not implemented!")
	}

}

func NetResultPreLogin(w http.ResponseWriter, r *http.Request) {
	var Response Login_Start_Response
	Response.RES = Consts_RES.SUCCESS
	Response.NativeToken = "FHxveQcSfC+xZ9nc6B5euzQX7sUQ44PVTTNizkQRhvHrVDPnytbXqaXjGLBlfkWpHoFR/g8g0+sgK3YTxznENOMBxJnJ7nKSVumBZCTqDcXN/7j0NBAeARwu0kzjh0FdP9j0rXMBcqG="

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultSessionUpdate(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Session_Update_Response
	Response.RES = Consts_RES.SUCCESS
	Response.NewGame = int(User.NewGame.Int64)

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultLogin(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	var Player Models.PlayerModel
	Player.ID = User.ID.String
	Player.Name = User.Name.String

	var Response Login_Game_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Result = int(User.Status.Int64)
	Response.LinkUUID = User.UUID.String
	Response.Player = &Player

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultCreateUser(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Create_User_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	err = db_commands.SetUserName(Request.TerminalId, Request.PlayerName)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to set username",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	err = db_commands.SetUserStatus(Request.TerminalId, Consts_Login.CREATING_PARTY)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to set user status",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	err = db_commands.SetUserNewGame(Request.TerminalId, 0)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to set user new game",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultPartyCreate(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Party_Create_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	team, err := db_commands.CreateTeam(Request.PartyName)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create team",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	for _, user := range Request.AMember {
		err = db_commands.SetUserTeam(user, team.TeamID.String)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to set team",
				"User", Request.TerminalId,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
		err = db_commands.InitSaveData(user)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to create inital save data",
				"User", Request.TerminalId,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
	}

	err = db_commands.SetUserStatus(Request.TerminalId, Consts_Login.PLAYING)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to set user status",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultHome(w http.ResponseWriter, r *http.Request, Did string) {
	IP := os.Getenv("DOMAIN")
	WSPort := os.Getenv("WS_PORT")

	User, err := db_commands.GetUser(Did)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserID := User.ID.String

	var Response Home_Response

	var Player Models.PlayerModel
	Player.ID = UserID
	Player.Name = User.Name.String
	Response.ATeamUser = append(Response.ATeamUser, Player)

	saveData, intro, err := DataHandlers.FetchSaveData(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	resumeData, err := db_commands.GetUserResume(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user resume",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	var EventInfoPointer Models.ResumeDataModel
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEvent(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Event_End_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	User, err := db_commands.GetUser(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserID := User.ID.String

	for _, LuaHash := range Request.ALua {
		err, errorfile := DataHandlers.ProccessLua(UserID, LuaHash)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to update user save",
				"User", Request.TerminalId,
				"LuaFile", errorfile,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
	}

	err = db_commands.UpdateUserResumeBool(UserID, 0)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to update user resume",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	saveData, _, err := DataHandlers.FetchSaveData(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save data",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEventCheck(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	resumeData, err := db_commands.GetUserResume(User.ID.String)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user resume",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	var EventInfoPointer Models.ResumeDataModel
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultScan(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Scan_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	User, err := db_commands.GetUser(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	UserSave, err := db_commands.GetUserSavaData(User.ID.String)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save data",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	UserScan, err := db_commands.GetUserScan(User.ID.String, Request.TagId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user scan data",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	err = db_commands.UpdateUserResume(User.ID.String, 1, uint32(UserScan.LuaHash.Int64), Request.TagId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user resume",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	if UserSave.NowHP.Int64 > 0 {
		Response.Lua = uint32(UserScan.LuaHash.Int64)
	} else if UserSave.NowHP.Int64 == 0 && (UserScan.Tag.String != Consts_ScanTag.QR_Loft6F_Burger && UserScan.Tag.String != Consts_ScanTag.QR_Magnet_Burger && UserScan.Tag.String != Consts_ScanTag.QR_Miyashita_2F_Burger && UserScan.Tag.String != Consts_ScanTag.QR_Miyashita_3F_Burger && UserScan.Tag.String != Consts_ScanTag.QR_Modi_Burger) {
		Response.Lua = Consts_LuaHash.Sys_EV_Burger_Other
	} else if UserSave.NowHP.Int64 == 0 && (UserScan.Tag.String == Consts_ScanTag.QR_Loft6F_Burger || UserScan.Tag.String == Consts_ScanTag.QR_Magnet_Burger || UserScan.Tag.String == Consts_ScanTag.QR_Miyashita_2F_Burger || UserScan.Tag.String == Consts_ScanTag.QR_Miyashita_3F_Burger || UserScan.Tag.String == Consts_ScanTag.QR_Modi_Burger) {
		Response.Lua = uint32(UserScan.LuaHash.Int64)
	}
	//Response.Result =  Used for Invalid QR Code?

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEventSaveResume(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Save_Resume_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	User, err := db_commands.GetUser(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	err = db_commands.UpdateUserResumeLuaResume(User.ID.String, Request.ResumeID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user resume",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopBenefit(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Shop_Benefit_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Scan_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = DataHandlers.FetchShopLua(Request.Code)
	//Response.Result = Used for Invalid QR Code

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopIdentifyStart(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Shop_Benefit_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Shop_Identify_Start_Response
	Response.RES = Consts_RES.SUCCESS
	//Response.Message = "This is a test" //Doesn't do anything I think?
	if Request.Code == "Badge" {
		Response.Type = 1
	} else if Request.Code == "Discount" {
		Response.Type = 2
	} else if Request.Code == "Card" {
		Response.Type = 3
	}
	//Response.Type = 0 //1: Badge, 2: Discount, 3: Bonus Card, Anything Else: Already got
	//Response.Result = Used for Invalid QR Code?

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultShopIdentifyEnd(w http.ResponseWriter, r *http.Request) {

	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultNewQuest(w http.ResponseWriter, r *http.Request, Did string) {
	User, err := db_commands.GetUser(Did)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	err = db_commands.UpdateUserSaveNewQuest(User.ID.String, 0)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to update user save",
			"User", Did,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response New_Quest_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Result = 1 //To-Do: Figure this out

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultBattleIn(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_In_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	err = db_commands.UpdateUserSaveBattleId(UserID, Request.BattleId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to update user save",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	BattleData := DataHandlers.FetchBattleData(Request.BattleId)
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
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultBattleAttackSucceeded(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_Attack_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user data",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	if !UserSave.BattleBadge1.Valid {
		err = db_commands.UpdateUserSaveBattleBadge1(UserID, Request.ItemId)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to get user save",
				"User", Request.TerminalId,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
	} else if !UserSave.BattleBadge2.Valid {
		err = db_commands.UpdateUserSaveBattleBadge2(UserID, Request.ItemId)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to get user save",
				"User", Request.TerminalId,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
	} else if !UserSave.BattleBadge3.Valid {
		err = db_commands.UpdateUserSaveBattleBadge3(UserID, Request.ItemId)
		if err != nil {
			ErrorInfo := Utils.FormatError(err.Error())
			slog.Error("Failed to get user save",
				"User", Request.TerminalId,
				"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
				"Function", ErrorInfo.FunctionName,
				"ErrorDetail", ErrorInfo.ErrorText)
		}
	}
	var Response Generic_Response
	Response.RES = Consts_RES.SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultBattleResult(w http.ResponseWriter, r *http.Request, body []byte) {
	var Request Battle_Result_Request
	err := json.Unmarshal(body, &Request)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to unmarshal JSON Request",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserID, err := db_commands.GetUserID(Request.TerminalId)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user info",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	err = db_commands.UpdateUserSaveNowHP(UserID, Request.NowHp)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	err = db_commands.EmptyUserSaveBadges(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to update user save",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	UserSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to get user save",
			"User", Request.TerminalId,
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	var Response Battle_Result_Response
	Response.RES = Consts_RES.SUCCESS
	Response.Lua = DataHandlers.FetchBattleLuaResult(int(UserSave.BattleID.Int64), Request.Succeeded, Request.NowHp)

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}
	slog.Debug("Response: " + string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}
