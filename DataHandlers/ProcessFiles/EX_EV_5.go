package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_5(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_5, Consts_LuaHash.Miyashita_EV_Mission_Cleared)
	if err != nil {
		return err
	}

	//Remove GPS
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_5_Tokyu_Plaza", 1)
	if err != nil {
		return err
	}

	Q105_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_105_EX_5)
	if err != nil {
		return err
	}
	if !Q105_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_105_EX_5, 1)
		if err != nil {
			return err
		}
	}

	return nil
}
