package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Loft_EV_Remaining(UserID string) error {
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_12_Fret_Badge, 1)
	if err != nil {
		return err
	}

	return nil
}
