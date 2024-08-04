package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_BT_Frog_Bridge(UserID string) error {
	//Add Quest Progress
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_3_Start_Battle_1, 1)
	if err != nil {
		return err
	}
	return nil
}
