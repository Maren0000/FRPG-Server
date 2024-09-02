package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapPinColor "FRPGServer/Consts/MapPinColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_Quiz_3_Clear(UserID string) error {
	err := db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_EX_Quiz_3_Hachiko, 1)
	if err != nil {
		return err
	}

	//Remove GPS
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Hachiko", 1)
	if err != nil {
		return err
	}

	Q10_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_10_Rindo_Battle)
	if err != nil {
		return err
	}
	if Q10_Cleared {
		Q100_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_100_Secret_Boss)
		if err != nil {
			return err
		}
		if !Q100_Cleared {
			err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Tokyo_Anime_Center", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Modi_lat, Consts_Coords.Modi_long, Consts_Quest.Quest_100_Secret_Boss, "", "")
			if err != nil {
				return err
			}
		}
	}

	Q111_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_111_EX_Quiz3)
	if err != nil {
		return err
	}
	if !Q111_Cleared {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_111_EX_Quiz3, 1)
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
