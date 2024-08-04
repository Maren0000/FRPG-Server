package ProcessFiles

import (
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_EV_Shiba_1(UserID string) error {
	//Add badge items
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_1)
	if err != nil {
		return err
	}
	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_3)
	if err != nil {
		return err
	}
	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_5)
	if err != nil {
		return err
	}

	//Update lua files for scans
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Rindo, Consts_LuaHash.Oioi_8F_EV_Rindo_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Fret, Consts_LuaHash.Oioi_8F_EV_Fret_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Nagi, Consts_LuaHash.Oioi_8F_EV_Nagi_2)
	if err != nil {
		return err
	}

	return nil
}
