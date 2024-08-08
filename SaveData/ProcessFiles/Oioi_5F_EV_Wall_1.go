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

func Oioi_5F_EV_Wall_1(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_5_Week_Death, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_6_Week_Battle)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_1000_Burger)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}

	//Create new scans
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q6_Vari1, 1, Consts_LuaHash.Oioi_5F_EV_Vari3_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q6_Pure1, 1, Consts_LuaHash.Oioi_5F_EV_Pure2_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q6_Pure2, 1, Consts_LuaHash.Oioi_5F_EV_Pure3_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q6_Yuusen1, 1, Consts_LuaHash.Oioi_5F_EV_Yuusen3_1)
	if err != nil {
		return err
	}

	//Update lua files for scans
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q5_Wall, Consts_LuaHash.Oioi_5F_EV_Wall_Def)
	if err != nil {
		return err
	}

	//Create Burger scan tags
	//The Web PDF version cheats by using one tag so we will do that for now.
	//To-Do: Add the other burger tags
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Oioi5F_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
	/*err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Oioi5F_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Oioi5F_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Burger, 1, Consts_LuaHash.Sys_EV_Burger_Open)
	if err != nil {
		return err
	}*/

	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q5_Wall", 1)
	if err != nil {
		return err
	}

	//Create new GPS for team members
	err = db_commands.CreateUserGPSPin(UserID, "Q6_Vari_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_5F_Vari_2_lat, Consts_Coords.Oioi_5F_Vari_2_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q6_Pure_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_5F_Pure_2_lat, Consts_Coords.Oioi_5F_Pure_2_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q6_Pure_2", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_5F_Pure_3_lat, Consts_Coords.Oioi_5F_Pure_3_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q6_Yuusen_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_5F_Yuusen_2_lat, Consts_Coords.Oioi_5F_Yuusen_2_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}

	//Create new GPS for burger shops
	//To-Do: Update these when I get to them
	err = db_commands.CreateUserGPSPin(UserID, "Oioi_5F_Burger", Consts_MapPin.Burger, "", Consts_Coords.Oioi_5F_Burger_lat, Consts_Coords.Oioi_5F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Modi_Burger", Consts_MapPin.Burger, "", Consts_Coords.Modi_F_Burger_lat, Consts_Coords.Modi_F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.ModiMap, "7")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Loft_6F_Burger", Consts_MapPin.Burger, "", Consts_Coords.Loft_6F_Burger_lat, Consts_Coords.Loft_6F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.LoftMap, "6")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_F_Burger", Consts_MapPin.Burger, "", Consts_Coords.Magnet_F_Burger_lat, Consts_Coords.Magnet_F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.MagnetMap, "5")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_2F_Burger", Consts_MapPin.Burger, "", Consts_Coords.Miyashita_2F_Burger_lat, Consts_Coords.Miyashita_2F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.MiyashitaParkMap, "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Miyashita_3F_Burger", Consts_MapPin.Burger, "", Consts_Coords.Miyashita_3F_Burger_lat, Consts_Coords.Miyashita_3F_Burger_long, Consts_Quest.Quest_1000_Burger, Consts_MapType.MiyashitaParkMap, "3")
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
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_R, 1, Consts_LuaHash.Oioi_5F_BT_Crow_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_R, 1, Consts_LuaHash.Oioi_5F_BT_Jellyfish_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_R, 1, Consts_LuaHash.Oioi_5F_BT_Wolf_R)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_R, 1, Consts_LuaHash.Oioi_5F_BT_Bear_R)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Crow_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Jellyfish_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_R, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Bear_R, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Crow_R", Consts_MapPin.Noise_Crow, Consts_MapPinColor.Noise_Crow_R, Consts_Coords.Oioi_5F_Crow_lat, Consts_Coords.Oioi_5F_Crow_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Jellyfish_R", Consts_MapPin.Noise_Jellyfish, Consts_MapPinColor.Noise_Jellyfish_R, Consts_Coords.Oioi_5F_Jellyfish_lat, Consts_Coords.Oioi_5F_Jellyfish_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Wolf_R", Consts_MapPin.Noise_Wolf, Consts_MapPinColor.Noise_Wolf_R, Consts_Coords.Oioi_5F_Wolf_lat, Consts_Coords.Oioi_5F_Wolf_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Bear_R", Consts_MapPin.Noise_Bear, Consts_MapPinColor.Noise_Bear_R, Consts_Coords.Oioi_5F_Bear_lat, Consts_Coords.Oioi_5F_Bear_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_B, 1, Consts_LuaHash.Oioi_5F_BT_Crow_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_B, 1, Consts_LuaHash.Oioi_5F_BT_Jellyfish_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_B, 1, Consts_LuaHash.Oioi_5F_BT_Wolf_B)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_B, 1, Consts_LuaHash.Oioi_5F_BT_Bear_B)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_Y, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Crow_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Jellyfish_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_B, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Bear_B, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Crow_B", Consts_MapPin.Noise_Crow, Consts_MapPinColor.Noise_Crow_B, Consts_Coords.Oioi_5F_Crow_lat, Consts_Coords.Oioi_5F_Crow_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Jellyfish_B", Consts_MapPin.Noise_Jellyfish, Consts_MapPinColor.Noise_Jellyfish_B, Consts_Coords.Oioi_5F_Jellyfish_lat, Consts_Coords.Oioi_5F_Jellyfish_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Wolf_B", Consts_MapPin.Noise_Wolf, Consts_MapPinColor.Noise_Wolf_B, Consts_Coords.Oioi_5F_Wolf_lat, Consts_Coords.Oioi_5F_Wolf_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Bear_B", Consts_MapPin.Noise_Bear, Consts_MapPinColor.Noise_Bear_B, Consts_Coords.Oioi_5F_Bear_lat, Consts_Coords.Oioi_5F_Bear_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		//Create QR scans
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_R, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_B, 1, Consts_LuaHash.Sys_Noise_Error_0)
		if err != nil {
			return err
		}

		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Crow_Y, 1, Consts_LuaHash.Oioi_5F_BT_Crow_Y)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Jellyfish_Y, 1, Consts_LuaHash.Oioi_5F_BT_Jellyfish_Y)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Wolf_Y, 1, Consts_LuaHash.Oioi_5F_BT_Wolf_Y)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, Consts_ScanTag.QR_Q6_Bear_Y, 1, Consts_LuaHash.Oioi_5F_BT_Bear_Y)
		if err != nil {
			return err
		}

		//Add items to quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Crow_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Jellyfish_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Wolf_Y, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Bear_Y, "off")
		if err != nil {
			return err
		}

		//Add GPS markers
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Crow_Y", Consts_MapPin.Noise_Crow, Consts_MapPinColor.Noise_Crow_Y, Consts_Coords.Oioi_5F_Crow_lat, Consts_Coords.Oioi_5F_Crow_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Jellyfish_Y", Consts_MapPin.Noise_Jellyfish, Consts_MapPinColor.Noise_Jellyfish_Y, Consts_Coords.Oioi_5F_Jellyfish_lat, Consts_Coords.Oioi_5F_Jellyfish_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Wolf_Y", Consts_MapPin.Noise_Wolf, Consts_MapPinColor.Noise_Wolf_Y, Consts_Coords.Oioi_5F_Wolf_lat, Consts_Coords.Oioi_5F_Wolf_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "Q6_Bear_Y", Consts_MapPin.Noise_Bear, Consts_MapPinColor.Noise_Bear_Y, Consts_Coords.Oioi_5F_Bear_lat, Consts_Coords.Oioi_5F_Bear_long, Consts_Quest.Quest_6_Week_Battle, Consts_MapType.MaruiMap, "5")
		if err != nil {
			return err
		}
	}
	return nil
}
