package APIHandlers

import (
	"FRPGServer/Utils"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type NativeSessionResponse struct {
	SharedSecurityKey string `json:"sharedSecurityKey"`
	NativeSessionId   string `json:"nativeSessionId"`
}

// ConnectSessionCreate
func BridgeHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Bridge hit")
	var Response NativeSessionResponse
	Response.SharedSecurityKey = "QetBknGIxXLS4zxfoUQIkFUKChDhvILG"
	Response.NativeSessionId = "73745dbcc387e6b90e04c98b433a1320"

	JSONResponse, err := json.Marshal(Response)
	if err != nil {
		ErrorInfo := Utils.FormatError(err.Error())
		slog.Error("Failed to create json response",
		"File", ErrorInfo.FileName+":"+strconv.Itoa(ErrorInfo.Line),
		"Function", ErrorInfo.FunctionName,
		"ErrorDetail", ErrorInfo.ErrorText)
	}

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	gz.Write(JSONResponse)
	gz.Close()
	sendbyte := buf.Bytes()

	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-type", "text/javascript; charset=utf-8")
	w.Header().Set("Cache-Control", "private, no-store, no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Content-Encoding", "gzip")
	w.Write(sendbyte)
}
