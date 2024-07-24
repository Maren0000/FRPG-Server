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
	var UserPointer UserModel
	UserPointer.ID = "kPjrXd"
	UserPointer.Name = "Maren"

	var EventInfoPointer ResumeDataModel
	EventInfoPointer.BResume = true
	EventInfoPointer.Lua = crc32.ChecksumIEEE([]byte("ComicEvent/introduction.lua"))
	EventInfoPointer.TagId = 0
	EventInfoPointer.ResumeId = 0

	var Response Home_Response
	Response.RES = Consts.R_SUCCESS
	Response.WebSocketServer = "ws://192.168.100.141:85"
	Response.RoomId = "J24YUe"
	Response.TeamId = "J24YUe"
	Response.TeamName = "MarenTeam"
	Response.ATeamUser = append(Response.ATeamUser, UserPointer)
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
