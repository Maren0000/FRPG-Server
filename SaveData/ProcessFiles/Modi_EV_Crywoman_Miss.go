package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Crywoman_Miss(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kanon, Consts_LuaHash.Modi_EV_Kanon_2)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Q9_Kanon", Consts_MapPin.Kanon, "", Consts_Coords.Modi_Kanon_lat, Consts_Coords.Modi_Kanon_long, Consts_Quest.Quest_9_Rindo_Badge, Consts_MapType.ModiMap, "7")
	if err != nil {
		return err
	}

	return nil
}
