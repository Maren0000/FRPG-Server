package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Wakare(UserID string) error {

	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Gorilla, Consts_LuaHash.Sys_Noise_Error_1)
	if err != nil {
		return err
	}

	//ClearGame?
	//To-Do: Maybe clear user data instead of locking out
	err = db_commands.SetUserStatusWithUserID(UserID, 10)
	if err != nil {
		return err
	}

	return nil
}
