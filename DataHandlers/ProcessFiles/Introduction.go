package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func ProcessIntroductionLua(UserID string) error {
	//Remove OnceEvent
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Introduction, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveIntro(UserID, 0)
	if err != nil {
		return err
	}
	return nil
}
