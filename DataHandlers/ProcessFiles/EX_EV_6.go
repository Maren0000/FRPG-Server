package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_6(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_6, Consts_LuaHash.Miyashita_EV_Mission_Cleared)
	if err != nil {
		return err
	}

	//Remove GPS
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_6_Shibuya_Stream", 1)
	if err != nil {
		return err
	}

	Q106_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_106_EX_6)
	if err != nil {
		return err
	}
	if !Q106_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_106_EX_6, 1)
		if err != nil {
			return err
		}
	}

	return nil
}
