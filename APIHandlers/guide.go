package APIHandlers

import (
	Consts_Protocol "FRPGServer/Consts/Protocol"
	Consts_RES "FRPGServer/Consts/Res"
	"FRPGServer/Utils"
	db_commands "FRPGServer/db/commands"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

func GuideHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Guide Hit")
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
	slog.Debug(string(DecryptedBody))
	var JSONRequest Generic_Request
	err = json.Unmarshal(DecryptedBody, &JSONRequest)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("JSON unmarshal failed",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	if !db_commands.CheckDeviceIdExists(JSONRequest.TerminalId) {
		db_commands.CreateNewUser(JSONRequest.TerminalId)
	}

	switch JSONRequest.PID {
	case Consts_Protocol.GET_URL:
		NetResultGetUrl(w, r)
	}
}

func NetResultGetUrl(w http.ResponseWriter, r *http.Request) {
	Domain := os.Getenv("DOMAIN")
	Port := os.Getenv("PORT")

	var Response Get_URL_Response
	Response.RES = Consts_RES.SUCCESS
	Response.API = "http://" + Domain + ":" + Port + "/gw000.php"
	Response.NSC = "http://" + Domain + ":" + Port + "/nativeBridge/native/session.php"
	Response.MainteURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki"
	Response.TermsURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/terms"
	Response.PrivacyURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/privacy"
	Response.LicenseURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/license"
	Response.CommerceURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/commerce"
	Response.TicketURL = "http://" + Domain + ":" + Port + "/ticket.php"
	Response.StoreURL = "https://play.google.com/store/apps/details?id=com.square_enix.android_googleplay.NTWEWYXFIELDWALKRPG"
	Response.EventEndDate = 1937679599 //Original date: 1637679599

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
			"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
			"Function", ErrorInfo.FunctionName,
			"ErrorDetail", ErrorInfo.ErrorText)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}
