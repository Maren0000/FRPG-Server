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

func Oioi_8F_EV_Deai_0(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_1_Start, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
	if err != nil {
		return err
	}
	/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}*/

	//Add items for quest panel
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Rindo, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Fret, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Nagi, "off")
	if err != nil {
		return err
	}

	//Update Scan Tag
	//To-Do: This should probably be remove not update
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q1_JoinBadge, Consts_LuaHash.Sys_Error_1)
	if err != nil {
		return err
	}

	//Add new scan tags
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q2_Rindo, 1, Consts_LuaHash.Oioi_8F_EV_Rindo_0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q2_Fret, 1, Consts_LuaHash.Oioi_8F_EV_Fret_0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q2_Nagi, 1, Consts_LuaHash.Oioi_8F_EV_Nagi_0)
	if err != nil {
		return err
	}

	//Update Local Map
	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MaruiMap, 8)
	if err != nil {
		return err
	}

	//Update GPS pins
	err = db_commands.CreateUserGPSPin(UserID, "Q2_Rindo", Consts_MapPin.Rindo, "", Consts_Coords.Oioi_8F_Rindo_lat, Consts_Coords.Oioi_8F_Rindo_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q2_Fret", Consts_MapPin.Fret, "", Consts_Coords.Oioi_8F_Fret_lat, Consts_Coords.Oioi_8F_Fret_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q2_Nagi", Consts_MapPin.Nagi, "", Consts_Coords.Oioi_8F_Nagi_lat, Consts_Coords.Oioi_8F_Nagi_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
	if err != nil {
		return err
	}

	return nil
}
