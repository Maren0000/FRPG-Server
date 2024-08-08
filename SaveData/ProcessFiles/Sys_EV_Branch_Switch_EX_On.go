package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

// To-Do: Add all the EX stuff
func Sys_EV_Branch_Switch_EX_On(UserID string) error {
	//Clear and start new quest
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_6_Week_Battle, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_7_Branch)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}
	return nil
}
