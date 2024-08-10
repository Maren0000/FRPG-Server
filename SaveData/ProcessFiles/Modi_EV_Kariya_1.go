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

func Modi_EV_Kariya_1(UserID string) error {
	Q8_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_8_Rindo_Death)
	if err != nil {
		return err
	}

	if Q8_isCurrent {
		err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_8_Rindo_Death, 1)
		if err != nil {
			return err
		}
	}

	//(Not in SERVER_SCRIPT) set the branch quest to not current in case it isn't already disabled
	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_7_Branch, 0)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_9_Rindo_Badge, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_1, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_2, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_3, "off")
	if err != nil {
		return err
	}

	//Set local map
	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.ModiMap, 7)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Modi_EV_Kariya_1, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kariya, Consts_LuaHash.Modi_EV_Kariya_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q8_Kariya", 1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_CryWoman, 1, Consts_LuaHash.Modi_EV_Crywoman_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q9_Crywoman", Consts_MapPin.Cry_Woman, "", Consts_Coords.Modi_Crywoman_lat, Consts_Coords.Modi_Crywoman_long, Consts_Quest.Quest_9_Rindo_Badge, Consts_MapType.ModiMap, "7")
	if err != nil {
		return err
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Modi_Girl, 1, Consts_LuaHash.Modi_EV_Girl_1)
	if err != nil {
		return err
	}

	//Color check
	Color, err := db_commands.GetUserSavaColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserGPSPin(UserID, "Q9_Shark_R", Consts_MapPin.Noise_Shark, Consts_MapPinColor.Noise_Shark_R, Consts_Coords.Modi_Shark_lat, Consts_Coords.Modi_Shark_long, Consts_Quest.Quest_7_Branch, Consts_MapType.ModiMap, "7")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserGPSPin(UserID, "Q9_Shark_B", Consts_MapPin.Noise_Shark, Consts_MapPinColor.Noise_Shark_B, Consts_Coords.Modi_Shark_lat, Consts_Coords.Modi_Shark_long, Consts_Quest.Quest_7_Branch, Consts_MapType.ModiMap, "7")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserGPSPin(UserID, "Q9_Shark_Y", Consts_MapPin.Noise_Shark, Consts_MapPinColor.Noise_Shark_Y, Consts_Coords.Modi_Shark_lat, Consts_Coords.Modi_Shark_long, Consts_Quest.Quest_7_Branch, Consts_MapType.ModiMap, "7")
		if err != nil {
			return err
		}
	}

	return nil
}
