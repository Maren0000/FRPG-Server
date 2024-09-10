package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Pure1_1(UserID string) error {
	err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_21_Last_Battle, Consts_QuestItem.Quest_21_Hint_2_1, "on")
	if err != nil {
		return err
	}
	//Not in SERVER_SCRIPT. Helps guide user with mission update.
	err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}

	return nil
}
