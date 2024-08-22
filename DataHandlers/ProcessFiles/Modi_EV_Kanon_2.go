package ProcessFiles

import (
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Kanon_2(UserID string) error {
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_11)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kanon, Consts_LuaHash.Modi_EV_Kanon_3)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Kanon", 1)
	if err != nil {
		return err
	}
	return nil
}
