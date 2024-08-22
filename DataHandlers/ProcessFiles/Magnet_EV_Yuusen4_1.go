package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Yuusen4_1(UserID string) error {
	//to-do: add tflite
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Art_2, Consts_LuaHash.Magnet_EV_Art_3F_Headgear)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_2, "hint")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Art_2", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Art_2_lat, Consts_Coords.Magnet_Art_2_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Yuusen", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Yuusen, Consts_LuaHash.Magnet_EV_Yuusen4_2)
	if err != nil {
		return err
	}

	return nil
}
