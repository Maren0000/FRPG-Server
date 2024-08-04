package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func ProcessIntroductionLua(UserID string) error {
	err := db_commands.CreateUserOnceEvent(UserID, Consts_LuaHash.Introduction, 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserSaveIntro(UserID, 0)
	if err != nil {
		return err
	}
	return nil
}
