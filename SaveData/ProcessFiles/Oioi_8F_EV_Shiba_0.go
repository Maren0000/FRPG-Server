package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
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

func Oioi_8F_EV_Shiba_0(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_3_Start_Battle_1, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}

	//Update lua files for scans
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Rindo, Consts_LuaHash.Oioi_8F_EV_Rindo_Def)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Fret, Consts_LuaHash.Oioi_8F_EV_Fret_Def)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Nagi, Consts_LuaHash.Oioi_8F_EV_Nagi_Def)
	if err != nil {
		return err
	}

	//Create new scans
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q4_Pure1, 1, Consts_LuaHash.Oioi_8F_EV_Pure1_0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q4_Vari1, 1, Consts_LuaHash.Oioi_8F_EV_Vari1_0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q4_Yuusen1, 1, Consts_LuaHash.Oioi_8F_EV_Yuusen1_0)
	if err != nil {
		return err
	}

	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q2_Rindo", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q2_Fret", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q2_Nagi", 1)
	if err != nil {
		return err
	}

	//Create new GPS for team members
	err = db_commands.CreateUserGPSPin(UserID, "Q3_Pure_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Pure_1_lat, Consts_Coords.Oioi_8F_Pure_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q3_Vari_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Vari_1_lat, Consts_Coords.Oioi_8F_Vari_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q3_Yuusen_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Yuusen_1_lat, Consts_Coords.Oioi_8F_Yuusen_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}

	//Color check
	Color, err := db_commands.GetUserSavaColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R2)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_R_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_R_2, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_R1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_R2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B2)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_B_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_B_2, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_B1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_B2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y2)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_Y_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_Y_2, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_Y1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q3_Frog_Y2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3_Start_Battle_1, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	}
	return nil
}
