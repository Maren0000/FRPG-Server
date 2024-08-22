package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Mpanel_KW_11(UserID string) error {
	err := db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_1, "off")
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_2, "off")
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_9_Rindo_Badge, Consts_QuestItem.Quest_9_3, "on")
	if err != nil {
		return err
	}

	return nil
}
