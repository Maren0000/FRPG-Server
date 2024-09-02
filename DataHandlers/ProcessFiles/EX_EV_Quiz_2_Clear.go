package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_Quiz_2_Clear(UserID string) error {
	err := db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Quiz_2_Tsutaya, 1)
	if err != nil {
		return err
	}

	//Remove GPS
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Shibuya_Tsutaya", 1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_100_Secret_Boss, Consts_QuestItem.Quest_100_Hint_3, "on")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_100_Secret_Boss, Consts_QuestItem.Quest_100_Hint_4, "on")
	if err != nil {
		return err
	}

	Q110_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_110_EX_Quiz2)
	if err != nil {
		return err
	}
	if !Q110_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_110_EX_Quiz2, 1)
		if err != nil {
			return err
		}
	}

	//Not in SERVER_SCRIPT. Mainly to avoid issues later on.
	err = db_commands.UpdateUserResumeLuaResume(UserID, 0)
	if err != nil {
		return err
	}

	return nil
}
