package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_BT_Gorilla_Lose(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Gorilla, Consts_LuaHash.Miyashita_EV_Gorilla_Rematch)
	if err != nil {
		return err
	}

	return nil
}
