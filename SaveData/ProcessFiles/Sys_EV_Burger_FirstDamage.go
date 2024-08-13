package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

// To-Do see if this is triggered after 2 or 3 damage
func Sys_EV_Burger_FirstDamage(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Sys_EV_Burger_FirstDamage, 1)
	if err != nil {
		return err
	}
	return nil
}
