package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func Sys_BT_Lose(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Oioi_5F_EV_FirstBattle, 1)
	if err != nil {
		return err
	}
	return nil
}
