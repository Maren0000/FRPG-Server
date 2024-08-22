package ProcessFiles

import (
	Consts_Building "FRPGServer/Consts/Building"
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Hint_3(UserID string) error {
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_14)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserBuilding(UserID, Consts_Building.MiyashitaPark, "宮下公園", "2")
	if err != nil {
		return err
	}

	//To-Do: Add GPS 8
	err = db_commands.CreateUserGPSEventPin(UserID, "GPS_VoiceOnly_7", Consts_Coords.GPS_Voice_7_Lat, Consts_Coords.GPS_Voice_7_Long, Consts_LuaHash.GPS_7)
	if err != nil {
		return err
	}

	return nil
}
