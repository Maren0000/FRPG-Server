package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Oioi_5F_BT_Wolf_Win(UserID string) error {
	//Turn off first battle (if it was not already off)
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Oioi_5F_EV_FirstBattle, 1)
	if err != nil {
		return err
	}

	//Remove Scan Tag
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q6_Pure1, 1)
	if err != nil {
		return err
	}

	//Add Quest Progress
	err = db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_6_Week_Battle, 1)
	if err != nil {
		return err
	}

	//Color check
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		//Update lua files for scans
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q6_Wolf_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		//Update Quest Panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_R, "on")
		if err != nil {
			return err
		}

		//Remove Current GPS markers
		err = db_commands.UpdateUserGPSRemove(UserID, "Q6_Wolf_R", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		//Update lua files for scans
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q6_Wolf_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		//Update Quest Panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_B, "on")
		if err != nil {
			return err
		}

		//Remove Current GPS markers
		err = db_commands.UpdateUserGPSRemove(UserID, "Q6_Wolf_B", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		//Update lua files for scans
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q6_Wolf_Y, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		//Update Quest Panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_Y, "on")
		if err != nil {
			return err
		}

		//Remove Current GPS markers
		err = db_commands.UpdateUserGPSRemove(UserID, "Q6_Wolf_Y", 1)
		if err != nil {
			return err
		}
	}
	return nil
}
