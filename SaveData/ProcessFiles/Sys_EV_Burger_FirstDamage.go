package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

// To-Do seems to trigger when low health (1 or 2)
func Sys_EV_Burger_FirstDamage(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Sys_EV_Burger_FirstDamage, 1)
	if err != nil {
		return err
	}
	return nil
}
