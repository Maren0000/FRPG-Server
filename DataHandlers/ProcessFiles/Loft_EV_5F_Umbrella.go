package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

//Not actually used afaik

func Loft_EV_5F_Umbrella(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Umbrella, Consts_LuaHash.Loft_EV_Shopping_Error_2)
	if err != nil {
		return err
	}

	//Disable pin?

	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Umbrella, "on")
	if err != nil {
		return err
	}

	return nil
}
