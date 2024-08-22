package ProcessFiles

import (
	Consts_Building "FRPGServer/Consts/Building"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Mission_Clear(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Modi_EV_FirstBattle, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_10_Rindo_Battle, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kariya, Consts_LuaHash.Modi_EV_Kariya_4)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q8_Kariya", 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Shark_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_R, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Shark_R", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Shark_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_B, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Shark_B", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Shark_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_Y, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Shark_B", 1)
		if err != nil {
			return err
		}
	}

	err = db_commands.SetUserLocalMap(UserID, "", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_Modi", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserBuildingStatus(UserID, Consts_Building.Modi, "1")
	if err != nil {
		return err
	}

	return nil
}
