package APIHandlers

import (
	Consts_Protocol "FRPGServer/Consts/Protocol"
	Consts_RES "FRPGServer/Consts/Res"
	"FRPGServer/Utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GuideHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Guide Hit")
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
	case Consts_Protocol.GET_URL:
		NetResultGetUrl(w, r)
	}
}

func NetResultGetUrl(w http.ResponseWriter, r *http.Request) {
	var Response Get_URL_Response
	Response.RES = Consts_RES.SUCCESS
	Response.API = "http://192.168.100.141:80/gw000.php"
	Response.NSC = "http://192.168.100.141:80/nativeBridge/native/session.php"
	Response.MainteURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki"
	Response.TermsURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/terms/"
	Response.PrivacyURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/privacy/"
	Response.LicenseURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/license"
	Response.CommerceURL = "https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/commerce"
	Response.TicketURL = "http://192.168.100.141:80/ticket.php"
	Response.StoreURL = "https://play.google.com/store/apps/details?id=com.square_enix.android_googleplay.NTWEWYXFIELDWALKRPG"
	Response.EventEndDate = 1937679599 //Original date: 1637679599

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		fmt.Println(err)
	}

	sendbyte := Utils.DESEncrypt(JSONResponse)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sendbyte)
}
