package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Crywoman_Miss(UserID string) error {
	//The SERVER_SCRIPT doesn't have a check for if the user already has the right item for some reason
	//So I will make my own to avoid bugs
	UserItems, err := db_commands.ListUserItems(UserID)
	if err != nil {
		return err
	}

	for _, item := range UserItems {
		if item.Item.Int64 == Consts_Item.Item_11 {
			return nil
		} else {
			continue
		}
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kanon, Consts_LuaHash.Modi_EV_Kanon_2)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q9_Kanon", Consts_MapPin.Kanon, "", Consts_Coords.Modi_Kanon_lat, Consts_Coords.Modi_Kanon_long, Consts_Quest.Quest_9_Rindo_Badge, Consts_MapType.ModiMap, "7")
	if err != nil {
		return err
	}

	return nil
}
