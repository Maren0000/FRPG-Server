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

func Magnet_EV_Mission_Clear(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_19_Nagi_Battle, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_6)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q15_Kubo", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Pure", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Yuusen", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Vari1", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Magnet_Vari2", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Pure, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Yuusen, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Vari1, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Magnet_Vari2, 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Elephant_R, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_R, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q19_Elephant_R", 1)
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Elephant_B, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_B, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q19_Elephant_B", 1)
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Elephant_Y, Consts_LuaHash.Sys_Noise_Error_1)
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_Y, "on")
		if err != nil {
			return err
		}

		err = db_commands.UpdateUserGPSRemove(UserID, "Q19_Elephant_Y", 1)
		if err != nil {
			return err
		}

	}

	err = db_commands.SetUserLocalMap(UserID, "", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_Magnet", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserBuildingStatus(UserID, Consts_Building.Magnet, "1")
	if err != nil {
		return err
	}

	return nil
}
