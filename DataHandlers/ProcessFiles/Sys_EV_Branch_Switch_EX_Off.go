package ProcessFiles

import (
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Sys_EV_Branch_Switch_EX_Off(UserID string) error {
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}

	Q100_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_100_Secret_Boss)
	if err != nil {
		return err
	}
	if !Q100_Cleared {
		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Rhino_R, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Rhino_B, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Rhino_Y, 1)
			if err != nil {
				return err
			}
		}
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_1_Taito_Station", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_1, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_2_Tower_Records", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_2, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_3_Animate", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_3, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_4_Mark_City", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_4, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_5_Tokyu_Plaza", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_5, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_6_Shibuya_Stream", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_6, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Parco", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Nazo, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Nazo_Goal, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Shibuya_109", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Quiz_1_109, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Shibuya_Tsutaya", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Quiz_2_Tsutaya, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Hachiko", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Quiz_3_Hachiko, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Tokyo_Anime_Center", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Minamimoto, 1)
	if err != nil {
		return err
	}

	return nil
}
