package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Scorpion_Win(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Miyashita_EV_Noise_Firstbattle, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_21_Last_Battle, 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Scorpion_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Scorpion_R", 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_R, "on")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Scorpion_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Scorpion_B", 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_B, "on")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Scorpion_Y, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Scorpion_Y", 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_Y, "on")
		if err != nil {
			return err
		}

	}

	return nil
}
