package ProcessFiles

import (
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_BT_Frog_1(UserID string) error {
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}

	//Update quest panel based on ColorID
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_R_1, "on")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_B_1, "on")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_3_Start_Battle_1, Consts_QuestItem.Quest_3_Frog_Y_1, "on")
		if err != nil {
			return err
		}
	}
	return nil
}
