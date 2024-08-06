package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Oioi_5F_EV_Pure3_1(UserID string) error {
	//Add item to quest panel
	err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Hint_2, "on")
	if err != nil {
		return err
	}
	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q6_Pure_2", 1)
	if err != nil {
		return err
	}
	return nil
}
