package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Oioi_5F_EV_Pure2_1(UserID string) error {
	//Add item to quest panel
	err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_6_Week_Battle, Consts_QuestItem.Quest_6_Hint_3, "on")
	if err != nil {
		return err
	}
	//Not in SERVER_SCRIPT. Helps guide user with mission update.
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}
	//Remove Current GPS markers
	err = db_commands.UpdateUserGPSRemove(UserID, "Q6_Pure_1", 1)
	if err != nil {
		return err
	}
	return nil
}
