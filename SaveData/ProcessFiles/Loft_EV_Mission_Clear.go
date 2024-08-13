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

func Loft_EV_Mission_Clear(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_14_Fret_Battle, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Uzuki, Consts_LuaHash.Loft_EV_Uduki_5)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q11_Uzuki", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Loft_Yuusen", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Loft_Yuusen1, 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Chameleon_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Chameleon_R, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q14_Chameleon_R", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Chameleon_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_7_Chameleon_B, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q14_Chameleon_B", 1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Chameleon_Y, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_7_Chameleon_Y, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q14_Chameleon_Y", 1)
		if err != nil {
			return err
		}
	}

	//Set local map
	err = db_commands.SetUserLocalMap(UserID, "", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_Loft", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserBuildingStatus(UserID, Consts_Building.Loft, "1")
	if err != nil {
		return err
	}

	return nil
}
