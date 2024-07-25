package APIHandlers

import (
	"FRPGServer/Consts"
	"FRPGServer/Utils"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
)

func Gw000Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gw000 Hit")
	RequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	DecryptedBody := Utils.DESDecrypt(RequestBody)
	fmt.Println(string(DecryptedBody))
	var JSONRequest Request
	err = json.Unmarshal(DecryptedBody, &JSONRequest)
	if err != nil {
		fmt.Println(err)
	}

	switch JSONRequest.PID {
	case Consts.P_LOGIN_START:
		NetResultPreLogin(w, r)
	case Consts.P_LOGIN_SESSION:
		NetResultSessionUpdate(w, r)
	case Consts.P_LOGIN_GAME:
		NetResultLogin(w, r)
	case Consts.P_NG_WORD:
		NetResultNGWord(w, r)
	case Consts.P_CREATE_USER:
		NetResultCreateUser(w, r)
	case Consts.P_PARTY_CREATE:
		NetResultPartyCreate(w, r)
	case Consts.P_HOME:
		NetResultHome(w, r)
	case Consts.P_PARTY_START:
		NetResultPartyStart(w, r)
	case Consts.P_EVENT:
		NetResultEvent(w, r)
	case Consts.P_EVENT_CHECK_RESUME:
		NetResultEventCheck(w, r)
	}

}

func NetResultPreLogin(w http.ResponseWriter, r *http.Request) {
	var Response Login_Start_Response
	Response.RES = Consts.R_SUCCESS
	Response.NativeToken = "FHxveQcSfC+xZ9nc6B5euzQX7sUQ44PVTTNizkQRhvHrVDPnytbXqaXjGLBlfkWpHoFR/g8g0+sgK3YTxznENOMBxJnJ7nKSVumBZCTqDcXN/7j0NBAeARwu0kzjh0FdP9j0rXMBcqG="

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultSessionUpdate(w http.ResponseWriter, r *http.Request) {
	var Response Session_Update_Response
	Response.RES = Consts.R_SUCCESS
	Response.NewGame = 1

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultLogin(w http.ResponseWriter, r *http.Request) {
	var Player PlayerModel
	Player.ID = "kPjrXd"
	Player.Name = "Maren"

	var Response Login_Game_Response
	Response.RES = Consts.R_SUCCESS
	Response.Result = Consts.L_PLAYING
	Response.LinkUUID = "00000000-0000-0000-0000-000000000001"
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
	Response.RES = Consts.R_SUCCESS
	Response.BOK = 1

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultCreateUser(w http.ResponseWriter, r *http.Request) {
	var Response Generic_Response
	Response.RES = Consts.R_SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultPartyCreate(w http.ResponseWriter, r *http.Request) {
	var Response Generic_Response
	Response.RES = Consts.R_SUCCESS

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultHome(w http.ResponseWriter, r *http.Request) {
	var UserPointer1 UserModel
	UserPointer1.ID = "kPjrXd"
	UserPointer1.Name = "Maren"

	/*var UserPointer2 UserModel
	var UserPointer3 UserModel
	var UserPointer4 UserModel*/

	var EventInfoPointer ResumeDataModel
	EventInfoPointer.BResume = false
	EventInfoPointer.Lua = crc32.ChecksumIEEE([]byte("ComicEvent/introduction.lua"))
	EventInfoPointer.TagId = 0
	EventInfoPointer.ResumeId = 0

	var QuestPointer QuestModel
	QuestPointer.ID = 1
	QuestPointer.Value = 1

	var GPSPointer GPSModel
	GPSPointer.ID = 1
	GPSPointer.Name = "marui"
	GPSPointer.PinType = "main"
	GPSPointer.PinColor = "y"
	GPSPointer.Latitude = 35.646530
	GPSPointer.Longitude = 139.708430
	GPSPointer.LuaScript = crc32.ChecksumIEEE([]byte("ComicEvent/oioi8f/oioi8f_ev_deai_000.lua"))
	GPSPointer.BLocationEvent = 0
	GPSPointer.Quest = &QuestPointer
	GPSPointer.MapType = "GPSMap"

	var Response Home_Response
	Response.RES = Consts.R_SUCCESS
	Response.ACharacter = []CharacterModel{}
	Response.AScan = []ScanModel{}
	Response.ARemoveScan = []string{}
	Response.AGPS = append(Response.AGPS, GPSPointer)
	Response.ARemoveGPS = []string{}
	Response.AOnceEvent = []uint32{}
	Response.ARemoveOnceEvent = []uint32{}
	Response.ABuildings = []BuildingModel{}
	Response.NowHp = 100
	Response.MaxHp = 500
	Response.ColorId = 1
	Response.Quest = &QuestPointer
	Response.BNewQuest = false
	Response.AItemList = make([]int, 0)
	Response.LocalMap = &LocalMapModel{}
	Response.WebSocketServer = "ws://192.168.100.141:85"
	Response.RoomId = "J24YUe"
	Response.TeamId = "J24YUe"
	Response.TeamName = "MarenTeam"
	Response.ATeamUser = append(Response.ATeamUser, UserPointer1)
	Response.EventInfo = &EventInfoPointer
	Response.BIntro = true

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
	var Response Party_Start_Response
	Response.RES = Consts.R_SUCCESS
	Response.PartyId = "cHifjK"
	Response.WebSocketServer = "ws://192.168.100.141:85"

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEvent(w http.ResponseWriter, r *http.Request) {

	var QuestPointer QuestModel
	QuestPointer.ID = 1
	QuestPointer.Value = 1

	var GPSPointer GPSModel
	GPSPointer.ID = 1
	GPSPointer.Name = "marui"
	GPSPointer.PinType = "main"
	GPSPointer.PinColor = "y"
	GPSPointer.Latitude = 35.646530
	GPSPointer.Longitude = 139.708430
	GPSPointer.LuaScript = crc32.ChecksumIEEE([]byte("ComicEvent/oioi8f/oioi8f_ev_deai_000.lua"))
	GPSPointer.BLocationEvent = 0
	GPSPointer.Quest = &QuestPointer
	GPSPointer.MapType = "GPSMap"

	var Response Home_Response
	Response.RES = Consts.R_SUCCESS
	Response.ACharacter = []CharacterModel{}
	Response.AScan = []ScanModel{}
	Response.ARemoveScan = []string{}
	Response.AGPS = append(Response.AGPS, GPSPointer)
	Response.ARemoveGPS = []string{}
	Response.AOnceEvent = []uint32{}
	Response.ARemoveOnceEvent = []uint32{}
	Response.ABuildings = []BuildingModel{}
	Response.NowHp = 100
	Response.MaxHp = 500
	Response.ColorId = 1
	Response.Quest = &QuestPointer
	Response.BNewQuest = false
	Response.AItemList = make([]int, 0)
	Response.LocalMap = &LocalMapModel{}

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(JSONResponse))
	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}

func NetResultEventCheck(w http.ResponseWriter, r *http.Request) {

	var EventInfoPointer ResumeDataModel
	EventInfoPointer.BResume = false
	EventInfoPointer.Lua = crc32.ChecksumIEEE([]byte("ComicEvent/introduction.lua"))
	EventInfoPointer.TagId = 0
	EventInfoPointer.ResumeId = 0

	var Response Event_Check_Resume_Response
	Response.RES = Consts.R_SUCCESS
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
