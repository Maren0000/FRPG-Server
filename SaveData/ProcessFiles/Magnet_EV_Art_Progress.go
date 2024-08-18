package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Art_Progress(UserID string) error {
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_16_Nagi_Badge, 1)
	if err != nil {
		return err
	}

	return nil
}
