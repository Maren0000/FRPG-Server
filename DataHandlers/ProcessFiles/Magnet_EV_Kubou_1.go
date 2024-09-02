package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Kubou_1(UserID string) error {
	Q15_Exists, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_15_Nagi_Death)
	if err != nil {
		return err
	}

	if !Q15_Exists {
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_15_Nagi_Death, 1)
		if err != nil {
			return err
		}
	}

	Q15_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_15_Nagi_Death)
	if err != nil {
		return err
	}

	if Q15_isCurrent {
		err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_15_Nagi_Death, 1)
		if err != nil {
			return err
		}
	}

	//Create new quest
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_16_Nagi_Badge, 1)
	if err != nil {
		return err
	}

	//(Not in SERVER_SCRIPT)
	/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}*/

	//(Not in SERVER_SCRIPT) set the branch quest to not current in case it isn't already disabled
	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_7_Branch, 0)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_1, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_2, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_3, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_16_Nagi_Badge, Consts_QuestItem.Quest_16_Stamp_4, "hint")
	if err != nil {
		return err
	}

	//Set local map
	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MagnetMap, 7)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Q15_Kubo", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Magnet_EV_Kubou_1, 1)
	if err != nil {
		return err
	}

	//To-Do: Add TFLITE if I ever figure that out
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Art1", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Art_1_lat, Consts_Coords.Magnet_Art_1_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "7")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Art_1, 1, Consts_LuaHash.Magnet_EV_Art_7F_Animal)
	if err != nil {
		return err
	}

	return nil
}
