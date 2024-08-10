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
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_EV_Frog3ren_1(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_3_Start_Battle_1, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_4_Start_Battle_2, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}

	//Update GPS pointers
	err = db_commands.UpdateUserGPSQuest(UserID, "Q3_Pure_1", Consts_Quest.Quest_4_Start_Battle_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSQuest(UserID, "Q3_Vari_1", Consts_Quest.Quest_4_Start_Battle_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSQuest(UserID, "Q3_Yuusen_1", Consts_Quest.Quest_4_Start_Battle_2)
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
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R3)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R4)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_R_3, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R5)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_R_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_R_2, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_R_3, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_R1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_3_lat, Consts_Coords.Oioi_8F_Frog_3_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_R2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_4_lat, Consts_Coords.Oioi_8F_Frog_4_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_R3", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_5_lat, Consts_Coords.Oioi_8F_Frog_5_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B3)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B4)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_B_3, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B5)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_B_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_B_2, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_B_3, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_B1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_3_lat, Consts_Coords.Oioi_8F_Frog_3_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_B2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_4_lat, Consts_Coords.Oioi_8F_Frog_4_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_B3", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_5_lat, Consts_Coords.Oioi_8F_Frog_5_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y3)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y4)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q4_Frog_Y_3, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y5)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_Y_1, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_Y_2, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_Y_3, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_Y1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_3_lat, Consts_Coords.Oioi_8F_Frog_3_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_Y2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_4_lat, Consts_Coords.Oioi_8F_Frog_4_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q4_Frog_Y3", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_5_lat, Consts_Coords.Oioi_8F_Frog_5_long, Consts_Quest.Quest_4_Start_Battle_2, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
	}

	return nil
}
