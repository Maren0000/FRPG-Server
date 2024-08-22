package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	"FRPGServer/Utils"
	db_commands "FRPGServer/db/commands"
)

func Loft_EV_Uduki_1(UserID string) error {
	Q11_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_11_Fret_Death)
	if err != nil {
		return err
	}

	if Q11_isCurrent {
		err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_11_Fret_Death, 1)
		if err != nil {
			return err
		}
	}

	//Create new quest
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_12_Fret_Badge, 1)
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

	//Set local map
	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.LoftMap, 6)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Loft_EV_Uduki_1, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Uzuki, Consts_LuaHash.Loft_EV_Uduki_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q11_Uzuki", 1)
	if err != nil {
		return err
	}

	//GPS markers are commented out
	//IsHit checks
	if Utils.IsHit(0) {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Cup, 1, Consts_LuaHash.Loft_EV_3F_Cup)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_WaterBottle, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Cup, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Cup, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_WaterBottle, 1, Consts_LuaHash.Loft_EV_3F_Bottle)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_WaterBottle, "off")
		if err != nil {
			return err
		}
	}

	if Utils.IsHit(1) {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Book, 1, Consts_LuaHash.Loft_EV_3F_Book)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Soap, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Book, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Book, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Soap, 1, Consts_LuaHash.Loft_EV_4F_Holder)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Soap, "off")
		if err != nil {
			return err
		}
	}

	if Utils.IsHit(2) {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Radio, 1, Consts_LuaHash.Loft_EV_4F_Radio)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Clock, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Radio, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Radio, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Clock, 1, Consts_LuaHash.Loft_EV_4F_Clock)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Clock, "off")
		if err != nil {
			return err
		}
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Bag, 1, Consts_LuaHash.Loft_EV_5F_Bag)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Umbrella, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Bag, "off")
	if err != nil {
		return err
	}

	if Utils.IsHit(4) {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Lanturn, 1, Consts_LuaHash.Loft_EV_5F_Lanthanum)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Doll, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Lantern, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Lanturn, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Doll, 1, Consts_LuaHash.Loft_EV_6F_Stuffed)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Doll, "off")
		if err != nil {
			return err
		}
	}

	if Utils.IsHit(5) {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Puzzle, 1, Consts_LuaHash.Loft_EV_6F_Puzzle)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Balloon, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Puzzle, "off")
		if err != nil {
			return err
		}
	} else {
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Puzzle, 1, Consts_LuaHash.Loft_EV_Shopping_Error_1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Loft_Search_Balloon, 1, Consts_LuaHash.Loft_EV_6F_Balloon)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_12_Fret_Badge, Consts_QuestItem.Quest_12_Balloon, "off")
		if err != nil {
			return err
		}
	}

	return nil
}
