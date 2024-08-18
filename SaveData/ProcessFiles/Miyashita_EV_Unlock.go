package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Unlock(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_7_Branch, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_20_Last_Death, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_20_Last_Death, Consts_QuestItem.Quest_20_Badge, "off")
	if err != nil {
		return err
	}

	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Badge, 1, Consts_LuaHash.Miyashita_EV_Shiba_1)
	if err != nil {
		return err
	}
	//Add badge GPS here but the player would never see it so *shrug*
	err = db_commands.CreateUserGPSPin(UserID, "GPS_Miyashita", Consts_MapPin.MapPin, "", Consts_Coords.Miyashita_lat, Consts_Coords.Miyashita_long, Consts_Quest.Quest_20_Last_Death, "", "")
	if err != nil {
		return err
	}

	//Set back to GPS map
	err = db_commands.SetUserLocalMap(UserID, "", 0)
	if err != nil {
		return err
	}
	return nil
}
