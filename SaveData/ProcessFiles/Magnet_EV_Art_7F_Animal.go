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

func Magnet_EV_Art_7F_Animal(UserID string) error {
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_16_Nagi_Badge, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_16_Nagi_Badge, 0)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_17_Nagi_Badge_2, 0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_17_Nagi_Badge_2, Consts_QuestItem.Quest_17_Kubo, "off")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_4, "on")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_3)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q15_Kubo", Consts_MapPin.Kubo, "", Consts_Coords.Magnet_Kubo_lat, Consts_Coords.Magnet_Kubo_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "7")
	if err != nil {
		return err
	}

	//To-do: Remove TFLIFE if it is figured out
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Art1", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Art_1, 1)
	if err != nil {
		return err
	}

	return nil
}
