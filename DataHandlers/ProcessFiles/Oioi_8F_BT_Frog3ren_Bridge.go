package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_BT_Frog3ren_Bridge(UserID string) error {
	//Add Quest Progress
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_4_Start_Battle_2, 1)
	if err != nil {
		return err
	}
	return nil
}
