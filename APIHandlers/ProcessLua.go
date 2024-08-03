package APIHandlers

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
	"fmt"
	"strconv"
)

func proccessLua(UserID string, LuaHash uint32) error {
	//Welcome to hell
	switch LuaHash {
	case Consts_LuaHash.Introduction:
		err := db_commands.CreateUserOnceEvent(UserID, Consts_LuaHash.Introduction, 0)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserSaveIntro(UserID, 0)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Deai_0:
		//Clear and start new quest
		err := db_commands.UpdateUserQuest(UserID, Consts_Quest.Quest_1_Start, 1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_2_Start_Twisters)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
		if err != nil {
			fmt.Println(err)
		}

		//Add items for quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Rindo, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Fret, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Nagi, "off")
		if err != nil {
			return err
		}

		//Update Scan Tag
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q1_JoinBadge, Consts_LuaHash.Sys_Error_1)
		if err != nil {
			return err
		}

		//Add new scan tags
		err = db_commands.CreateUserScan(UserID, 2, 3, Consts_ScanTag.QR_Q2_Rindo, 1, Consts_LuaHash.Oioi_8F_EV_Rindo_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, 3, Consts_ScanTag.QR_Q2_Fret, 1, Consts_LuaHash.Oioi_8F_EV_Fret_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 4, 3, Consts_ScanTag.QR_Q2_Nagi, 1, Consts_LuaHash.Oioi_8F_EV_Nagi_0)
		if err != nil {
			return err
		}

		//Update Local Map
		err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MaruiMap, 8)
		if err != nil {
			return err
		}

		//Update GPS pins
		err = db_commands.CreateUserGPSPin(UserID, "2", "Q2_Rindo", Consts_MapPin.Rindo, "", Consts_Coords.Rindo_lat, Consts_Coords.Rindo_long, 2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "3", "Q2_Fret", Consts_MapPin.Fret, "", Consts_Coords.Fret_lat, Consts_Coords.Fret_long, 2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "4", "Q2_Nagi", Consts_MapPin.Nagi, "", Consts_Coords.Nagi_lat, Consts_Coords.Nagi_long, 2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}

		return nil
	default:
		fmt.Println("Unknown LuaHash to process: " + strconv.FormatUint(uint64(LuaHash), 10))
		return nil
	}
}

func fetchSaveData(UserID string) (save PlayerDataModel, bIntro bool, error error) {
	PlayerCharacters, err := db_commands.ListUserCharacters(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerCharacters {
		var temp CharacterModel
		temp.CharacterId = int(data.CharacterId.Int64)
		temp.ItemId = int(data.ItemId.Int64)

		save.ACharacter = append(save.ACharacter, temp)
	}

	PlayerScans, err := db_commands.ListUserScan(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerScans {
		var temp ScanModel
		temp.ID = int(data.ID.Int64)
		temp.Type = int(data.Type.Int64)
		temp.Tag = data.Tag.String
		temp.Height = float32(data.Height.Float64)
		if data.BMulti.Int64 == 0 {
			temp.BMulti = false
		} else {
			temp.BMulti = true
		}

		save.AScan = append(save.AScan, temp)
	}

	PlayerBuildings, err := db_commands.ListUserBuilding(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerBuildings {
		var temp BuildingModel
		temp.Name = data.Name.String
		temp.Prefab = data.Prefab.String
		temp.Status = data.Status.String

		save.ABuildings = append(save.ABuildings, temp)
	}

	PlayerQuest, err := db_commands.GetCurrentUserQuest(UserID)
	if err != nil {
		return save, false, err
	}
	var CurrentQuest QuestModel
	CurrentQuest.ID = int(PlayerQuest.ID.Int64)
	CurrentQuest.Value = int(PlayerQuest.Value.Int64)

	PlayerQuestItems, err := db_commands.ListUserQuestItems(UserID, int(PlayerQuest.ID.Int64))
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerQuestItems {
		var temp QuestItemModel
		temp.Name = data.Name.String
		temp.Type = data.Type.String

		CurrentQuest.ItemList = append(CurrentQuest.ItemList, temp)
	}

	PlayerGPS, err := db_commands.ListUserGPS(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerGPS {
		var temp GPSModel
		temp.ID = data.ID.String
		temp.Name = data.Name.String
		temp.PinType = data.PinType.String
		temp.PinColor = data.PinColor.String
		temp.Latitude = data.Latitude.Float64
		temp.Longitude = data.Longitude.Float64
		temp.BLocationEvent = int(data.BLocationEvent.Int64)
		temp.LuaScript = uint32(data.LuaScript.Int64)
		temp.Quest = &CurrentQuest
		temp.MapType = data.MapType.String
		temp.MapNo = data.MapNo.String

		save.AGPS = append(save.AGPS, temp)
	}

	PlayerOnceEvent, err := db_commands.ListUserOnceEvent(UserID)
	if err != nil {
		return save, false, err
	}
	var OnceEventTemp []uint32
	for _, data := range PlayerOnceEvent {
		OnceEventTemp = append(OnceEventTemp, uint32(data.UInt.Int64))
	}
	save.AOnceEvent = OnceEventTemp

	var RemovedScansTemp []string
	PlayerRemovedScans, err := db_commands.ListUserScanRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedScans {
		RemovedScansTemp = append(RemovedScansTemp, data.Tag.String)
	}
	save.ARemoveScan = RemovedScansTemp

	var RemovedGPSTemp []string
	PlayerRemovedGPS, err := db_commands.ListUserGPSRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedGPS {
		RemovedGPSTemp = append(RemovedGPSTemp, data.Name.String) //Might be ID idk
	}
	save.ARemoveGPS = RemovedGPSTemp

	var RemovedOnceEventTemp []uint32
	PlayerRemovedOnceEvent, err := db_commands.ListUserOnceEventRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedOnceEvent {
		RemovedOnceEventTemp = append(RemovedOnceEventTemp, uint32(data.UInt.Int64))
	}
	save.ARemoveOnceEvent = RemovedOnceEventTemp

	var itemTemp []int
	PlayerItems, err := db_commands.ListUserItems(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerItems {
		itemTemp = append(itemTemp, int(data.Item.Int64))
	}
	save.AItemList = itemTemp

	PlayerCurrentLocalMap, err := db_commands.GetCurrentUserLocalMap(UserID)
	if err != nil {
		return save, false, err
	}
	var LocalmapTemp LocalMapModel
	LocalmapTemp.Name = PlayerCurrentLocalMap.Name.String
	LocalmapTemp.Floor = int(PlayerCurrentLocalMap.Floor.Int64)
	save.LocalMap = &LocalmapTemp

	PlayerSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		return save, false, err
	}

	save.NowHp = int(PlayerSave.NowHP.Int64)
	save.MaxHp = int(PlayerSave.MaxHP.Int64)
	save.ColorId = int(PlayerSave.ColorID.Int64)
	save.Quest = &CurrentQuest
	if PlayerSave.BNewQuest.Int64 == 0 {
		save.BNewQuest = false
	} else {
		save.BNewQuest = true
	}

	if PlayerSave.BIntro.Int64 == 0 {
		return save, false, nil
	} else {
		return save, true, nil
	}
}
