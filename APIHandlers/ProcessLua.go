package APIHandlers

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func proccessLua(UserID string, LuaHash uint32) {
	switch LuaHash {
	case Consts_LuaHash.Introduction:
		db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Introduction, 1)
		//To-Do: Set bIntro to false in UserSave
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
		//return save, false, err need better error handling
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
