package ProcessFiles

import (
	Consts_Character "FRPGServer/Consts/Character"
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_EV_Shiba_2(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_4_Start_Battle_2, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_5_Week_Death, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}

	//Add item to quest panel
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_5_Week_Death, Consts_QuestItem.Quest_5_Wall, "off")
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
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q4_Pure1, Consts_LuaHash.Oioi_8F_EV_Pure1_Def)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q4_Vari1, Consts_LuaHash.Oioi_8F_EV_Vari1_Def)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q4_Yuusen1, Consts_LuaHash.Oioi_8F_EV_Yuusen1_Def)
	if err != nil {
		return err
	}

	//Create new scans
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Q5_Wall, 1, Consts_LuaHash.Oioi_5F_EV_Wall_1)
	if err != nil {
		return err
	}

	//Add character (Not in SERVER_SCRIPT) and item
	_, err = db_commands.CreateUserCharacter(UserID, Consts_Character.Minamimoto, Consts_Item.Item_7)
	if err != nil {
		return err
	}
	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_7)
	if err != nil {
		return err
	}

	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q3_Pure_1", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q3_Vari_1", 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q3_Yuusen_1", 1)
	if err != nil {
		return err
	}

	//Create new GPS for team members
	err = db_commands.CreateUserGPSPin(UserID, "Q5_Wall", Consts_MapPin.Shinigami, "", Consts_Coords.Oioi_5F_Wall_lat, Consts_Coords.Oioi_5F_Wall_long, Consts_Quest.Quest_5_Week_Death, Consts_MapType.MaruiMap, "5")
	if err != nil {
		return err
	}

	//Update Local Map
	err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MaruiMap, 5)
	if err != nil {
		return err
	}

	return nil
}
