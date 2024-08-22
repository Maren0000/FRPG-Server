package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Art_3F_Headgear(UserID string) error {
	err := db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_2, "on")
	if err != nil {
		return err
	}

	//to-do: add tflite
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Art_2, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Art_2", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Yuusen", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Yuusen, Consts_LuaHash.Magnet_EV_Yuusen4_Def)
	if err != nil {
		return err
	}

	return nil
}
