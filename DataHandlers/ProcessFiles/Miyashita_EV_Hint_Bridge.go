package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Hint_Bridge(UserID string) error {
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_7_Branch, 1)
	if err != nil {
		return err
	}
	//Not in SERVER_SCRIPT
	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_7_Branch, 1)
	if err != nil {
		return err
	}

	return nil
}
