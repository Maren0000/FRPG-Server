package multi

import (
	"FRPGServer/Utils"
	"fmt"
	"hash/crc32"
	"log"
	"net/http"
	"strconv"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func ServerObj() *socketio.Server {
	serverWS := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	serverWS.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	serverWS.OnEvent("/", "join", func(s socketio.Conn, msg map[string]string) {
		log.Println("join crc:", msg["crc"])
		log.Println("join data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)

		resData := map[string]any{}
		resData["name"] = "test"
		resByte, err := Utils.WriteHashMap(resData)
		if err != nil {
			fmt.Println(err)
		}
		resEnc, err := Utils.WSEncrypt(resByte, 3723)

		//map: roomId - playerId - playerName
		//res: name? ArrayList of playerdata?
		res := map[string]string{}
		res["data"] = resEnc
		res["crc"] = strconv.Itoa(int(crc32.Checksum(resByte, crc32.MakeTable(crc32.IEEE))))
		s.Emit("join", res)
	})

	serverWS.OnEvent("/", "battleEscape", func(s socketio.Conn, msg map[string]string) {
		log.Println("battleEscape crc:", msg["crc"])
		log.Println("battleEscape data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "teamDisband", func(s socketio.Conn, msg map[string]string) {
		log.Println("teamDisband crc:", msg["crc"])
		log.Println("teamDisband data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "teamExist", func(s socketio.Conn, msg map[string]string) {
		log.Println("teamExist crc:", msg["crc"])
		log.Println("teamExist data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncCancel", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncCancel crc:", msg["crc"])
		log.Println("syncCancel data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncClear", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncClear crc:", msg["crc"])
		log.Println("syncClear data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncStart", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncStart crc:", msg["crc"])
		log.Println("syncStart data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId - check - syncData - syncTag
		//res: data - tag - type
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "mute", func(s socketio.Conn, msg map[string]string) {
		log.Println("mute crc:", msg["crc"])
		log.Println("mute data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId - bMute
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncExist", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncExist crc:", msg["crc"])
		log.Println("syncExist data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
		//map: playerId - syncData- syncTag
		//map: data - tag
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "in", func(s socketio.Conn, msg any) {
		log.Println("in:", msg)
		//s.SetContext(msg)
		//return "recv " + msg
	})

	serverWS.OnEvent("/", "out", func(s socketio.Conn, msg any) {
		log.Println("out:", msg)
		//last := s.Context().(string)
		//s.Emit("out", last)
		//s.Close()
	})

	serverWS.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	serverWS.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
		s.Namespace()
	})

	return serverWS
}
