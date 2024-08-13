package ProcessFiles

import (
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_BT_Frog3ren_2(UserID string) error {
	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}

	//Update quest panel based on ColorID
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_R_2, "on")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_B_2, "on")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_4_Start_Battle_2, Consts_QuestItem.Quest_4_Frog_Y_2, "on")
		if err != nil {
			return err
		}
	}
	return nil
}
