package ProcessFiles

import (
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Shiba_3(UserID string) error {

	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_22_Last_Boss, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_1000_Burger, 1)
	if err != nil {
		return err
	}

	return nil
}
