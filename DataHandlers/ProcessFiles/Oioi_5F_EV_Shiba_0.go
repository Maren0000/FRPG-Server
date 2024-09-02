package ProcessFiles

import (
	Consts_Building "FRPGServer/Consts/Building"
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Oioi_5F_EV_Shiba_0(UserID string) error {
	//Create new scans
	err := db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Kariya, 1, Consts_LuaHash.Modi_EV_Kariya_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Kanon, 1, Consts_LuaHash.Modi_EV_Kanon_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Uzuki, 1, Consts_LuaHash.Loft_EV_Uduki_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Yuusen1, 1, Consts_LuaHash.Loft_EV_Yuusen2_Tips)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Kubo, 1, Consts_LuaHash.Magnet_EV_Kubou_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Vari2, 1, Consts_LuaHash.Magnet_EV_Vari1_Tips)
	if err != nil {
		return err
	}

	//GPS Events
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_1", Consts_Coords.GPS_Voice_1_Lat, Consts_Coords.GPS_Voice_1_Long, Consts_LuaHash.GPS_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_2", Consts_Coords.GPS_Voice_2_Lat, Consts_Coords.GPS_Voice_2_Long, Consts_LuaHash.GPS_2)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_3", Consts_Coords.GPS_Voice_3_Lat, Consts_Coords.GPS_Voice_3_Long, Consts_LuaHash.GPS_3)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_4", Consts_Coords.GPS_Voice_4_Lat, Consts_Coords.GPS_Voice_4_Long, Consts_LuaHash.GPS_4)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_5", Consts_Coords.GPS_Voice_5_Lat, Consts_Coords.GPS_Voice_5_Long, Consts_LuaHash.GPS_5)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_6", Consts_Coords.GPS_Voice_6_Lat, Consts_Coords.GPS_Voice_6_Long, Consts_LuaHash.GPS_6)
	if err != nil {
		return err
	}

	//Not in SERVER_SCRIPT. Add QR codes for GPS events
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_1, 1, Consts_LuaHash.GPS_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_2, 1, Consts_LuaHash.GPS_2)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_3, 1, Consts_LuaHash.GPS_3)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_4, 1, Consts_LuaHash.GPS_4)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_5, 1, Consts_LuaHash.GPS_5)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_6, 1, Consts_LuaHash.GPS_6)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_GPS_7, 1, Consts_LuaHash.GPS_7)
	if err != nil {
		return err
	}

	//Color check
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_R, 1, Consts_LuaHash.Modi_BT_Shark_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_R, 1, Consts_LuaHash.Loft_BT_Chameleon_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_R, 1, Consts_LuaHash.Magnet_BT_Elephant_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Chameleon_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_R, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_B, 1, Consts_LuaHash.Modi_BT_Shark_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_B, 1, Consts_LuaHash.Loft_BT_Chameleon_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_B, 1, Consts_LuaHash.Magnet_BT_Elephant_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Chameleon_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_B, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Shark_Y, 1, Consts_LuaHash.Modi_BT_Shark_Y)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Chameleon_Y, 1, Consts_LuaHash.Loft_BT_Chameleon_Y)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Elephant_Y, 1, Consts_LuaHash.Magnet_BT_Elephant_Y)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Shark_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Chameleon_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_7_Branch, Consts_QuestItem.Quest_7_Elephant_Y, "off")
		if err != nil {
			return err
		}
	}

	err = db_commands.SetUserLocalMap(UserID, "", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserBuildingStatus(UserID, Consts_Building.Marui, "1")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserBuilding(UserID, Consts_Building.Modi, "MODI", "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserBuilding(UserID, Consts_Building.Loft, "LOFT", "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserBuilding(UserID, Consts_Building.Magnet, "MAGNET", "2")
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "GPS_Modi", Consts_MapPin.MapPin, "", Consts_Coords.Modi_lat, Consts_Coords.Modi_long, Consts_Quest.Quest_7_Branch, "", "")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "GPS_Loft", Consts_MapPin.MapPin, "", Consts_Coords.Loft_lat, Consts_Coords.Loft_long, Consts_Quest.Quest_7_Branch, "", "")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Loft_Yuusen", Consts_MapPin.MapPin, "", Consts_Coords.Loft_Yuusen_lat, Consts_Coords.Loft_Yuusen_long, Consts_Quest.Quest_7_Branch, Consts_MapType.LoftMap, "6")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "GPS_Magnet", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_lat, Consts_Coords.Magnet_long, Consts_Quest.Quest_7_Branch, "", "")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Vari2", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Vari2_lat, Consts_Coords.Magnet_Vari2_long, Consts_Quest.Quest_7_Branch, Consts_MapType.MagnetMap, "7")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "MaruiPin", 1)
	if err != nil {
		return err
	}

	//THIS ENTIRE CODE IS COPIED IN ANOTHER FILE WTF
	//IsHit checks
	/*IsHit0, err := db_commands.GetUserSaveIsHit0(UserID)
	if err != nil {
		return err
	}
	if IsHit0 == 0 {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Cup, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_WaterBottle, "off")
		if err != nil {
			return err
		}
	}

	IsHit1, err := db_commands.GetUserSaveIsHit1(UserID)
	if err != nil {
		return err
	}
	if IsHit1 == 0 {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Book, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Soap, "off")
		if err != nil {
			return err
		}
	}

	IsHit2, err := db_commands.GetUserSaveIsHit2(UserID)
	if err != nil {
		return err
	}
	if IsHit2 == 0 {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Radio, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Clock, "off")
		if err != nil {
			return err
		}
	}

	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Bag, "off")
	if err != nil {
		return err
	}

	IsHit4, err := db_commands.GetUserSaveIsHit4(UserID)
	if err != nil {
		return err
	}
	if IsHit4 == 0 {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Lantern, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Doll, "off")
		if err != nil {
			return err
		}
	}

	IsHit5, err := db_commands.GetUserSaveIsHit5(UserID)
	if err != nil {
		return err
	}
	if IsHit5 == 0 {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Puzzle, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Balloon, "off")
		if err != nil {
			return err
		}
	}*/

	return nil
}
