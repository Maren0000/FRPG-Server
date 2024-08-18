package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Uduki_2(UserID string) error {
	err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_22_Last_Boss, Consts_QuestItem.Quest_22_Hint_2, "on")
	if err != nil {
		return err
	}

	return nil
}
