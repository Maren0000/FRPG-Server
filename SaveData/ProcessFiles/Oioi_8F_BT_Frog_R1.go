package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_BT_Frog_R1(UserID string) error {
	//Update lua file for scan
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q3_Frog_R_1, Consts_LuaHash.Sys_Noise_Error_1)
	if err != nil {
		return err
	}

	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q3_Frog_R1", 1)
	if err != nil {
		return err
	}

	return nil
}
