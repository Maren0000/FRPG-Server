package ProcessFiles

import (
	db_commands "FRPGServer/db/commands"
)

func Sys_EV_Burger_Refresh(UserID string) error {
	err := db_commands.UpdateUserSaveNowHP(UserID, 5)
	if err != nil {
		return err
	}
	return nil
}
