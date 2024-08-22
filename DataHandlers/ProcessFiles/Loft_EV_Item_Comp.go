package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Loft_EV_Item_Comp(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Uzuki, Consts_LuaHash.Loft_EV_Uduki_3)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q11_Uzuki", Consts_MapPin.Uzuki, "", Consts_Coords.Loft_Uzuki_lat, Consts_Coords.Loft_Uzuki_long, Consts_Quest.Quest_11_Fret_Death, Consts_MapType.LoftMap, "6")
	if err != nil {
		return err
	}

	//(Not in SERVER_SCRIPT) set the badge quest to not current in case it isn't already disabled
	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_12_Fret_Badge, 0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_13_Fret_Badge_2, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}*/
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_13_Fret_Badge_2, Consts_QuestItem.Quest_13_Uzuki, "off")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Bag, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Balloon, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Book, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Clock, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Cup, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Doll, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Lanturn, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Puzzle, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Radio, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Soap, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_Umbrella, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Search_WaterBottle, Consts_LuaHash.Loft_EV_Shopping_Error_3)
	if err != nil {
		return err
	}

	return nil
}
