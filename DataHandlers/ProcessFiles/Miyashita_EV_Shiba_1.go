package ProcessFiles

import (
	Consts_Character "FRPGServer/Consts/Character"
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapPinColor "FRPGServer/Consts/MapPinColor"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Shiba_1(UserID string) error {
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_8)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserCharacterItem(UserID, Consts_Character.Minamimoto, Consts_Item.Item_8)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_20_Last_Death, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_21_Last_Battle, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_1", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_2", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_3", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_4", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_5", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_6", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_7", 1)
	if err != nil {
		return err
	}
	/*err = db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_8", 1)
	if err != nil {
		return err
	}*/

	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_1_Taito_Station", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_2_Tower_Records", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_3_Animate", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_4_Mark_City", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_5_Tokyu_Plaza", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_6_Shibuya_Stream", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Parco", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Shibuya_109", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Shibuya_Tsutaya", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Hachiko", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "GPS_EX_Tokyo_Anime_Center", 1)
	if err != nil {
		return err
	}

	//Remove GPS badge here but the player won't ever see it so *shrug*
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Badge, Consts_LuaHash.Miyashita_EV_Mission_Addition)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Kariya", Consts_MapPin.Kariya, "", Consts_Coords.Miyashita_Kariya_lat, Consts_Coords.Miyashita_Kariya_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Kariya, 1, Consts_LuaHash.Miyashita_EV_Kariya_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Uzuki", Consts_MapPin.Uzuki, "", Consts_Coords.Miyashita_Uzuki_lat, Consts_Coords.Miyashita_Uzuki_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Uzuki, 1, Consts_LuaHash.Miyashita_EV_Uduki_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Kanon", Consts_MapPin.Kanon, "", Consts_Coords.Miyashita_Kanon_lat, Consts_Coords.Miyashita_Kanon_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Kanon, 1, Consts_LuaHash.Miyashita_EV_Kanon_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Pure", Consts_MapPin.MapPin, "", Consts_Coords.Miyashita_Pure_lat, Consts_Coords.Miyashita_Pure_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Pure_1, 1, Consts_LuaHash.Miyashita_EV_Pure1_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Vari_1", Consts_MapPin.MapPin, "", Consts_Coords.Miyashita_Vari_1_lat, Consts_Coords.Miyashita_Vari_1_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Vari_1, 1, Consts_LuaHash.Miyashita_EV_Vari2_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Vari_2", Consts_MapPin.MapPin, "", Consts_Coords.Miyashita_Vari_2_lat, Consts_Coords.Miyashita_Vari_2_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Vari_2, 1, Consts_LuaHash.Miyashita_EV_Vari3_1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Yuusen", Consts_MapPin.MapPin, "", Consts_Coords.Miyashita_Yuusen_lat, Consts_Coords.Miyashita_Yuusen_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Yuusen_1, 1, Consts_LuaHash.Miyashita_EV_Yuusen1_1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_R, 1, Consts_LuaHash.Miyashita_BT_Tyranno_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_R, 1, Consts_LuaHash.Miyashita_BT_Harisenbon_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_R, 1, Consts_LuaHash.Miyashita_BT_Scorpion_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Tyranno_R", Consts_MapPin.Noise_Tyranno, Consts_MapPinColor.Noise_Tyranno_R, Consts_Coords.Miyashita_Tyranno_lat, Consts_Coords.Miyashita_Tyranno_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Harisenbon_R", Consts_MapPin.Noise_Harisenbon, Consts_MapPinColor.Noise_Harisenbon_R, Consts_Coords.Miyashita_Harisenbon_lat, Consts_Coords.Miyashita_Harisenbon_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Scorpion_R", Consts_MapPin.Noise_Scorpion, Consts_MapPinColor.Noise_Scorpion_R, Consts_Coords.Miyashita_Scorpion_lat, Consts_Coords.Miyashita_Scorpion_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Tyranno_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Harisenbon_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_R, "off")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_B, 1, Consts_LuaHash.Miyashita_BT_Tyranno_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_B, 1, Consts_LuaHash.Miyashita_BT_Harisenbon_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_B, 1, Consts_LuaHash.Miyashita_BT_Scorpion_Bridge)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Tyranno_B", Consts_MapPin.Noise_Tyranno, Consts_MapPinColor.Noise_Tyranno_B, Consts_Coords.Miyashita_Tyranno_lat, Consts_Coords.Miyashita_Tyranno_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Harisenbon_B", Consts_MapPin.Noise_Harisenbon, Consts_MapPinColor.Noise_Harisenbon_B, Consts_Coords.Miyashita_Harisenbon_lat, Consts_Coords.Miyashita_Harisenbon_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Scorpion_B", Consts_MapPin.Noise_Scorpion, Consts_MapPinColor.Noise_Scorpion_B, Consts_Coords.Miyashita_Scorpion_lat, Consts_Coords.Miyashita_Scorpion_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Tyranno_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Harisenbon_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_B, "off")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Tyranno_Y, 1, Consts_LuaHash.Miyashita_BT_Tyranno_Bridge)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Harisenbon_Y, 1, Consts_LuaHash.Miyashita_BT_Harisenbon_Bridge)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Scorpion_Y, 1, Consts_LuaHash.Miyashita_BT_Scorpion_Bridge)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Tyranno_Y", Consts_MapPin.Noise_Tyranno, Consts_MapPinColor.Noise_Tyranno_Y, Consts_Coords.Miyashita_Tyranno_lat, Consts_Coords.Miyashita_Tyranno_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "2")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Harisenbon_Y", Consts_MapPin.Noise_Harisenbon, Consts_MapPinColor.Noise_Harisenbon_Y, Consts_Coords.Miyashita_Harisenbon_lat, Consts_Coords.Miyashita_Harisenbon_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Scorpion_Y", Consts_MapPin.Noise_Scorpion, Consts_MapPinColor.Noise_Scorpion_Y, Consts_Coords.Miyashita_Scorpion_lat, Consts_Coords.Miyashita_Scorpion_long, Consts_Quest.Quest_21_Last_Battle, Consts_MapType.MiyashitaParkMap, "3")
		if err != nil {
			return err
		}

		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Tyranno_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Harisenbon_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Scorpion_Y, "off")
		if err != nil {
			return err
		}

	}

	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MiyashitaParkMap, 2)
	if err != nil {
		return err
	}

	return nil
}
