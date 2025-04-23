package multi

import (
	"FRPGServer/Utils"
	"encoding/base64"
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
		PartyNames := []any{}
		mem2 := map[string]any{}
		mem2["name"] = "Neku"
		mem2["id"] = "11111"
		PartyNames = append(PartyNames, mem2)
		mem3 := map[string]any{}
		mem3["name"] = "Shiki"
		mem3["id"] = "11112"
		PartyNames = append(PartyNames, mem3)
		mem4 := map[string]any{}
		mem4["name"] = "Beat"
		mem4["id"] = "11113"
		PartyNames = append(PartyNames, mem4)
		mem5 := map[string]any{}
		mem5["name"] = "Rhyme"
		mem5["id"] = "11114"
		PartyNames = append(PartyNames, mem5)
		resData["aMember"] = PartyNames
		resByte := Utils.WriteRequest(resData)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(base64.StdEncoding.EncodeToString(resByte))
		res := map[string]string{}
		res["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByte)), 16)
		resEnc, err := Utils.WSEncrypt(resByte, 3723)

		//map: roomId - playerId - playerName
		//res: name? ArrayList of playerdata?

		res["data"] = resEnc
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

	serverWS.OnEvent("/", "syncTag", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncTag crc:", msg["crc"])
		log.Println("syncTag data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		test, err := Utils.ReadHashMap(data)
		fmt.Println(test)
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

	serverWS.OnEvent("/", "in", func(s socketio.Conn, msg map[string]string) {
		log.Println("in crc:", msg["crc"])
		log.Println("in data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "out", func(s socketio.Conn, msg map[string]string) {
		log.Println("out crc:", msg["crc"])
		log.Println("out data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "kick", func(s socketio.Conn, msg map[string]string) {
		log.Println("kick crc:", msg["crc"])
		log.Println("kick data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "message", func(s socketio.Conn, msg map[string]string) {
		log.Println("message crc:", msg["crc"])
		log.Println("message data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "stamp", func(s socketio.Conn, msg map[string]string) {
		log.Println("stamp crc:", msg["crc"])
		log.Println("stamp data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "script", func(s socketio.Conn, msg map[string]string) {
		log.Println("script crc:", msg["crc"])
		log.Println("script data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "syscall", func(s socketio.Conn, msg map[string]string) {
		log.Println("syscall crc:", msg["crc"])
		log.Println("syscall data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data, err := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
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
