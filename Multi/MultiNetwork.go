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

/*CS:
mute
syncStart
teamDisband
syncTag?
syncExist?
out?
in?
teamExist
*/

type Room struct {
	Players map[string]*RoomPlayer
}

type RoomPlayer struct {
	Name    string
	ID      string
	PartyID string
	TagHit  bool
}

var ActiveRooms map[string]Room = map[string]Room{}

var IPTrack map[string]*RoomPlayer = map[string]*RoomPlayer{}

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

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)

		s.LeaveAll() //By Default, the library makes a useless room for each client. That's dumb.
		s.Join(data["roomId"].(string))

		room, ok := ActiveRooms[data["roomId"].(string)]
		if !ok {
			newRoom := Room{}
			newRoom.Players = map[string]*RoomPlayer{}
			newPlayer := new(RoomPlayer)
			newPlayer.Name = data["playerName"].(string)
			newPlayer.ID = data["playerId"].(string)
			newPlayer.PartyID = data["roomId"].(string)
			newPlayer.TagHit = false

			newRoom.Players[newPlayer.ID] = newPlayer
			IPTrack[s.RemoteAddr().String()] = newPlayer
			room = newRoom
			ActiveRooms[data["roomId"].(string)] = newRoom
		} else {
			newPlayer := new(RoomPlayer)
			newPlayer.Name = data["playerName"].(string)
			newPlayer.ID = data["playerId"].(string)
			newPlayer.PartyID = data["roomId"].(string)
			newPlayer.TagHit = false

			room.Players[newPlayer.ID] = newPlayer
			IPTrack[s.RemoteAddr().String()] = newPlayer
			ActiveRooms[data["roomId"].(string)] = room
		}
		fmt.Println(room)

		resDataIn := map[string]any{}
		resDataIn["id"] = data["playerId"].(string)
		resDataIn["name"] = data["playerName"].(string)

		resByteIn := Utils.WriteRequest(resDataIn)
		resIn := map[string]string{}
		resIn["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteIn)), 16)
		resEncIn, err := Utils.WSEncrypt(resByteIn, 3723)
		resIn["data"] = resEncIn

		serverWS.BroadcastToRoom("/", data["roomId"].(string), "in", resIn)

		resData := map[string]any{}
		PartyMembers := []any{}

		for _, player := range room.Players {
			member := map[string]any{}
			member["name"] = player.Name
			member["id"] = player.ID
			PartyMembers = append(PartyMembers, member)
		}

		resData["aMember"] = PartyMembers
		resByte := Utils.WriteRequest(resData)
		res := map[string]string{}
		res["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByte)), 16)
		resEnc, err := Utils.WSEncrypt(resByte, 3723)

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

		test := Utils.ReadHashMap(data)
		fmt.Println(test)
		//input map: playerId
		//output: none just emit
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "teamDisband", func(s socketio.Conn, msg map[string]string) {
		log.Println("teamDisband crc:", msg["crc"])
		log.Println("teamDisband data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
		//input map: playerId

		for _, v := range ActiveRooms[s.Rooms()[0]].Players {
			if v.ID == data["playerId"] {
				if len(ActiveRooms[s.Rooms()[0]].Players) == 1 {
					delete(ActiveRooms, s.Rooms()[0])
				} else {
					delete(ActiveRooms[s.Rooms()[0]].Players, v.ID)
				}
			}
		}

		resDataOut := map[string]any{}
		resDataOut["id"] = data["playerId"]

		resByteOut := Utils.WriteRequest(resDataOut)
		resOut := map[string]string{}
		resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
		resEncOut, err := Utils.WSEncrypt(resByteOut, 3723)
		resOut["data"] = resEncOut

		serverWS.BroadcastToRoom("/", s.Rooms()[0], "teamDisband", resOut)
		s.Leave(s.Rooms()[0])
	})

	//Should set PartyID here
	serverWS.OnEvent("/", "teamExist", func(s socketio.Conn, msg map[string]string) {
		log.Println("teamExist crc:", msg["crc"])
		log.Println("teamExist data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
		resDataOut := map[string]any{}
		resDataOut["id"] = data["playerId"]

		resByteOut := Utils.WriteRequest(resDataOut)
		resOut := map[string]string{}
		resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
		resEncOut, err := Utils.WSEncrypt(resByteOut, 3723)
		resOut["data"] = resEncOut

		serverWS.BroadcastToRoom("/", s.Rooms()[0], "teamExist", resOut)
		//input map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncCancel", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncCancel crc:", msg["crc"])
		log.Println("syncCancel data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)

		for _, player := range ActiveRooms[s.Rooms()[0]].Players {
			if player.ID == data["playerId"].(string) {
				player.TagHit = false
			}
		}

		resDataOut := map[string]any{}
		resDataOut["id"] = data["playerId"]
		resByteOut := Utils.WriteRequest(resDataOut)
		resOut := map[string]string{}
		resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
		resEncOut, _ := Utils.WSEncrypt(resByteOut, 3723)
		resOut["data"] = resEncOut
		//input map: playerId
		//Output map: id
		s.Emit("syncCancel", resEncOut)
	})

	serverWS.OnEvent("/", "syncClear", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncClear crc:", msg["crc"])
		log.Println("syncClear data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)

		for _, player := range ActiveRooms[s.Rooms()[0]].Players {
			if player.ID == data["playerId"].(string) {
				player.TagHit = false
			}
		}
		//input map: playerId
	})

	serverWS.OnEvent("/", "teamCreate", func(s socketio.Conn, msg map[string]string) {
		log.Println("teamCreate crc:", msg["crc"])
		log.Println("teamCreate data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
		//data["playerId"]
		//input map: playerId
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncStart", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncStart crc:", msg["crc"])
		log.Println("syncStart data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
		fmt.Println(ActiveRooms[s.Rooms()[0]])
		resDataOut := map[string]any{}
		ExistFlag := true
		PartyMembers := []any{}

		switch data["syncTag"].(string) {
		case "scanTagMode":
			for _, player := range ActiveRooms[s.Rooms()[0]].Players {
				if player.ID == data["playerId"].(string) {
					player.TagHit = true
					member := map[string]any{}
					member["name"] = player.Name
					member["id"] = player.ID
					PartyMembers = append(PartyMembers, member)
				} else if !player.TagHit {
					ExistFlag = false
				} else {
					member := map[string]any{}
					member["name"] = player.Name
					member["id"] = player.ID
					PartyMembers = append(PartyMembers, member)
				}
			}
		}

		resDataOut["res"] = "uh oh"
		resByteOut := Utils.WriteRequest(resDataOut)
		resOut := map[string]string{}
		resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
		resEncOut, _ := Utils.WSEncrypt(resByteOut, 3723)
		resOut["data"] = resEncOut
		//input map: playerId
		//Output map: id
		s.Emit("error", resOut)

		if ExistFlag {
			resDataOut["syncData"] = data["syncData"]
			resDataOut["syncTag"] = data["syncTag"]
			for _, player := range ActiveRooms[s.Rooms()[0]].Players {
				player.TagHit = false
			}
			resByteOut := Utils.WriteRequest(resDataOut)
			resOut := map[string]string{}
			resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
			resEncOut, _ := Utils.WSEncrypt(resByteOut, 3723)
			resOut["data"] = resEncOut
			serverWS.BroadcastToRoom("/", s.Rooms()[0], "syncExist", resOut)
		} else {
			resDataOut["aMember"] = PartyMembers
			resByteOut := Utils.WriteRequest(resDataOut)
			resOut := map[string]string{}
			resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
			resEncOut, _ := Utils.WSEncrypt(resByteOut, 3723)
			resOut["data"] = resEncOut
			serverWS.BroadcastToRoom("/", s.Rooms()[0], "syncTag", resOut)
		}
		//input map: playerId - check - syncData - syncTag
	})

	serverWS.OnEvent("/", "syncTag", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncTag crc:", msg["crc"])
		log.Println("syncTag data:", msg["data"])
		data, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		//output map: aMember
		test := Utils.ReadHashMap(data)
		fmt.Println(test)
	})

	serverWS.OnEvent("/", "mute", func(s socketio.Conn, msg map[string]string) {
		log.Println("mute crc:", msg["crc"])
		log.Println("mute data:", msg["data"])
		enc, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(enc)

		//input map: playerId - bMute
		//output map: playerId - bMute

		resDataIn := map[string]any{}
		resDataIn["playerId"] = data["playerId"].(string)
		resDataIn["bMute"] = data["bMute"].(bool)

		resByteIn := Utils.WriteRequest(resDataIn)
		resIn := map[string]string{}
		resIn["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteIn)), 16)
		resEncIn, err := Utils.WSEncrypt(resByteIn, 3723)
		resIn["data"] = resEncIn

		serverWS.BroadcastToRoom("/", s.Rooms()[0], "mute", resIn)
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "syncExist", func(s socketio.Conn, msg map[string]string) {
		log.Println("syncExist crc:", msg["crc"])
		log.Println("syncExist data:", msg["data"])
		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println((err))
		}

		data := Utils.ReadHashMap(decrypt)

		for _, player := range ActiveRooms[s.Rooms()[0]].Players {
			if player.ID == data["playerId"].(string) {
				player.TagHit = false
			}
		}
		//map: playerId - syncData- syncTag
		//output map: syncdata - synctag
		//s.Emit("reply", "have "+msg)
	})

	serverWS.OnEvent("/", "in", func(s socketio.Conn, msg map[string]string) {
		log.Println("in crc:", msg["crc"])
		log.Println("in data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		//Output map: id - name
		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "out", func(s socketio.Conn, msg map[string]string) {
		log.Println("out crc:", msg["crc"])
		log.Println("out data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		//Output map: id
		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "kick", func(s socketio.Conn, msg map[string]string) {
		log.Println("kick crc:", msg["crc"])
		log.Println("kick data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		//output map: id
		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "message", func(s socketio.Conn, msg map[string]string) {
		log.Println("message crc:", msg["crc"])
		log.Println("message data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		//output map: msg
		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "stamp", func(s socketio.Conn, msg map[string]string) {
		log.Println("stamp crc:", msg["crc"])
		log.Println("stamp data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "script", func(s socketio.Conn, msg map[string]string) {
		log.Println("script crc:", msg["crc"])
		log.Println("script data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		//Output map: id
		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnEvent("/", "syscall", func(s socketio.Conn, msg map[string]string) {
		log.Println("syscall crc:", msg["crc"])
		log.Println("syscall data:", msg["data"])

		decrypt, err := Utils.WSDecrypt(msg["data"], 3723)
		if err != nil {
			fmt.Println(err)
		}

		data := Utils.ReadHashMap(decrypt)
		fmt.Println(data)
	})

	serverWS.OnError("/", func(s socketio.Conn, e error) {
		//output: res
		log.Println("meet error:", e)
	})

	serverWS.OnDisconnect("/", func(s socketio.Conn, reason string) {
		clientIP := s.RemoteAddr().String()
		party := IPTrack[clientIP].PartyID

		room, ok := ActiveRooms[party]
		fmt.Println(room)
		if ok {
			for _, v := range room.Players {
				if v.ID == IPTrack[clientIP].ID {
					if len(room.Players) == 1 {
						delete(ActiveRooms, party)
					} else {
						delete(room.Players, v.ID)
					}
				}
			}
		}

		resDataOut := map[string]any{}
		resDataOut["id"] = IPTrack[clientIP].ID

		resByteOut := Utils.WriteRequest(resDataOut)
		resOut := map[string]string{}
		resOut["crc"] = strconv.FormatUint(uint64(crc32.ChecksumIEEE(resByteOut)), 16)
		resEncOut, _ := Utils.WSEncrypt(resByteOut, 3723)
		resOut["data"] = resEncOut

		serverWS.BroadcastToRoom("/", party, "out", resOut)
		fmt.Println(clientIP)
		log.Println("closed", reason)
	})

	return serverWS
}

func RemoveIndex(s []RoomPlayer, index int) []RoomPlayer {
	ret := make([]RoomPlayer, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
