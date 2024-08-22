package ProcessFiles

import (
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Girl_2(UserID string) error {
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_9)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Girl, Consts_LuaHash.Modi_EV_Girl_Def)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Gal", 1)
	if err != nil {
		return err
	}
	return nil
}
