package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapPinColor "FRPGServer/Consts/MapPinColor"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

// Note: isOrderableQuest should be flipped (true becomes false)
func Sys_EV_Branch_Switch_EX_On(UserID string) error {
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}

	Q100_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_100_Secret_Boss)
	if err != nil {
		return err
	}
	if !Q100_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_100_Secret_Boss, 0)
		if err != nil {
			return err
		}

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_100_Secret_Boss, Consts_QuestItem.Quest_100_Rhino_R, "off")
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_100_Secret_Boss, Consts_QuestItem.Quest_100_Rhino_B, "off")
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_100_Secret_Boss, Consts_QuestItem.Quest_100_Rhino_Y, "off")
			if err != nil {
				return err
			}
		}
	}

	Q100_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_100_Secret_Boss)
	if err != nil {
		return err
	}
	if !Q100_Cleared {
		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Rhino_R, 1, Consts_LuaHash.EX_BT_Rhino_Bridge)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Rhino_B, 1, Consts_LuaHash.EX_BT_Rhino_Bridge)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Rhino_Y, 1, Consts_LuaHash.EX_BT_Rhino_Bridge)
			if err != nil {
				return err
			}
		}

		Q107_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_107_EX_Minamimoto)
		if err != nil {
			return err
		}
		if !Q107_Exists {
			err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_107_EX_Minamimoto, 0)
			if err != nil {
				return err
			}
		}

		Q10_Clear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_10_Rindo_Battle)
		if err != nil {
			return err
		}
		if Q10_Clear {
			Q111_Clear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_111_EX_Quiz3)
			if err != nil {
				return err
			}
			if Q111_Clear {
				err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Tokyo_Anime_Center", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Modi_lat, Consts_Coords.Modi_long, Consts_Quest.Quest_100_Secret_Boss, "", "")
				if err != nil {
					return err
				}
			}
		}

		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Minamimoto, 0, Consts_LuaHash.EX_EV_Minamimoto_1)
		if err != nil {
			return err
		}
	}

	Q101_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_101_EX_1)
	if err != nil {
		return err
	}
	if !Q101_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_101_EX_1, 0)
		if err != nil {
			return err
		}
	}
	Q101_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_101_EX_1)
	if err != nil {
		return err
	}
	if !Q101_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_1_Taito_Station", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Tatio_Station_lat, Consts_Coords.Tatio_Station_long, Consts_Quest.Quest_101_EX_1, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_1, 1, Consts_LuaHash.EX_EV_1)
		if err != nil {
			return err
		}
	}

	Q102_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_102_EX_2)
	if err != nil {
		return err
	}
	if !Q102_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_102_EX_2, 0)
		if err != nil {
			return err
		}
	}
	Q102_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_102_EX_2)
	if err != nil {
		return err
	}
	if !Q102_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_2_Tower_Records", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Tower_Records_lat, Consts_Coords.Tower_Records_long, Consts_Quest.Quest_102_EX_2, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_2, 1, Consts_LuaHash.EX_EV_2)
		if err != nil {
			return err
		}
	}

	Q103_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_103_EX_3)
	if err != nil {
		return err
	}
	if !Q103_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_103_EX_3, 0)
		if err != nil {
			return err
		}
	}
	Q103_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_103_EX_3)
	if err != nil {
		return err
	}
	if !Q103_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_3_Animate", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Animate_lat, Consts_Coords.Animate_long, Consts_Quest.Quest_103_EX_3, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_3, 1, Consts_LuaHash.EX_EV_3)
		if err != nil {
			return err
		}
	}

	Q104_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_104_EX_4)
	if err != nil {
		return err
	}
	if !Q104_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_104_EX_4, 0)
		if err != nil {
			return err
		}
	}
	Q104_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_104_EX_4)
	if err != nil {
		return err
	}
	if !Q104_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_4_Mark_City", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Mark_City_lat, Consts_Coords.Mark_City_long, Consts_Quest.Quest_104_EX_4, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_4, 1, Consts_LuaHash.EX_EV_4)
		if err != nil {
			return err
		}
	}

	Q105_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_105_EX_5)
	if err != nil {
		return err
	}
	if !Q105_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_105_EX_5, 0)
		if err != nil {
			return err
		}
	}
	Q105_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_105_EX_5)
	if err != nil {
		return err
	}
	if !Q105_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_5_Tokyu_Plaza", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Tokyu_Plaza_lat, Consts_Coords.Tokyu_Plaza_long, Consts_Quest.Quest_105_EX_5, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_5, 1, Consts_LuaHash.EX_EV_5)
		if err != nil {
			return err
		}
	}

	Q106_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_106_EX_6)
	if err != nil {
		return err
	}
	if !Q106_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_106_EX_6, 0)
		if err != nil {
			return err
		}
	}
	Q106_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_106_EX_6)
	if err != nil {
		return err
	}
	if !Q106_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_6_Shibuya_Stream", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Shibuya_Stream_lat, Consts_Coords.Shibuya_Stream_long, Consts_Quest.Quest_106_EX_6, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_6, 1, Consts_LuaHash.EX_EV_6)
		if err != nil {
			return err
		}
	}

	Q108_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_108_EX_Nazo)
	if err != nil {
		return err
	}
	if !Q108_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_108_EX_Nazo, 0)
		if err != nil {
			return err
		}
	}
	Q108_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_108_EX_Nazo)
	if err != nil {
		return err
	}
	if !Q108_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Parco", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Parco_lat, Consts_Coords.Parco_long, Consts_Quest.Quest_108_EX_Nazo, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Nazo, 1, Consts_LuaHash.EX_EV_Nazo_Start)
		if err != nil {
			return err
		}
	}
	Q120_Current, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_120_EX_Nazo_Playing)
	if err != nil {
		return err
	}
	if Q120_Current {
		Q120_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_120_EX_Nazo_Playing)
		if err != nil {
			return err
		}
		if !Q120_Cleared {
			err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Parco", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Parco_lat, Consts_Coords.Parco_long, Consts_Quest.Quest_108_EX_Nazo, "", "")
			if err != nil {
				return err
			}
			//To-Do: Add Tflite scan
			err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Nazo_Goal, 1, Consts_LuaHash.EX_EV_Nazo_Goal)
			if err != nil {
				return err
			}
		}
	}

	Q109_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_109_EX_Quiz1)
	if err != nil {
		return err
	}
	if !Q109_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_109_EX_Quiz1, 0)
		if err != nil {
			return err
		}
	}
	Q109_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_109_EX_Quiz1)
	if err != nil {
		return err
	}
	if !Q109_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Shibuya_109", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Shibuya_109_lat, Consts_Coords.Shibuya_109_long, Consts_Quest.Quest_109_EX_Quiz1, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Quiz_1_109, 1, Consts_LuaHash.EX_EV_Quiz_1)
		if err != nil {
			return err
		}
	}

	Q110_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_110_EX_Quiz2)
	if err != nil {
		return err
	}
	if !Q110_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_110_EX_Quiz2, 0)
		if err != nil {
			return err
		}
	}
	Q110_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_110_EX_Quiz2)
	if err != nil {
		return err
	}
	if !Q110_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Shibuya_Tsutaya", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Shibuya_Tsutaya_lat, Consts_Coords.Shibuya_Tsutaya_long, Consts_Quest.Quest_110_EX_Quiz2, "", "")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Quiz_2_Tsutaya, 1, Consts_LuaHash.EX_EV_Quiz_2)
		if err != nil {
			return err
		}
	}

	Q111_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_111_EX_Quiz3)
	if err != nil {
		return err
	}
	if !Q111_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_111_EX_Quiz3, 0)
		if err != nil {
			return err
		}
	}
	Q111_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_111_EX_Quiz3)
	if err != nil {
		return err
	}
	if !Q111_Cleared {
		err = db_commands.CreateUserGPSPin(UserID, "GPS_EX_Hachiko", Consts_MapPin.MapPin, Consts_MapPinColor.MapPin_Color_Sub, Consts_Coords.Hachiko_lat, Consts_Coords.Hachiko_long, Consts_Quest.Quest_111_EX_Quiz3, "", "")
		if err != nil {
			return err
		}
		//To-Do: Add Tflite scan
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Quiz_3_Hachiko, 1, Consts_LuaHash.EX_EV_Quiz_3)
		if err != nil {
			return err
		}
	}

	Q7_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_7_Branch)
	if err != nil {
		return err
	}
	if !Q7_Exists {
		//Clear and start new quest
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_6_Week_Battle, 1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_7_Branch, 1)
		if err != nil {
			return err
		}
		//(Not in SERVER_SCRIPT)
		/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
		if err != nil {
			return err
		}*/
	}
	return nil
}
