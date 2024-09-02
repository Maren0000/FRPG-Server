package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_Rhino_Win(UserID string) error {
	err := db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Minamimoto, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Tokyo_Anime_Center", 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_Rhino_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_Rhino_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_Rhino_Y, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}
	}

	Q100_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_100_Secret_Boss)
	if err != nil {
		return err
	}
	if !Q100_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_100_Secret_Boss, 1)
		if err != nil {
			return err
		}
	}
	Q107_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_107_EX_Minamimoto)
	if err != nil {
		return err
	}
	if !Q107_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_107_EX_Minamimoto, 1)
		if err != nil {
			return err
		}
	}

	//Change mission once completed back to branch or miyashita

	return nil
}
