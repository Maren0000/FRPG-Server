package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_4(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_4, Consts_LuaHash.Miyashita_EV_Mission_Cleared)
	if err != nil {
		return err
	}

	//Remove GPS
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_4_Mark_City", 1)
	if err != nil {
		return err
	}

	Q104_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_104_EX_4)
	if err != nil {
		return err
	}
	if !Q104_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_104_EX_4, 1)
		if err != nil {
			return err
		}
	}

	return nil
}
