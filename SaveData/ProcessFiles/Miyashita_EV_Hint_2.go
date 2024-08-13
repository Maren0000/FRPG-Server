package ProcessFiles

import (
	Consts_Item "FRPGServer/Consts/Item"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Hint_2(UserID string) error {
	_, err := db_commands.CreateUserItem(UserID, Consts_Item.Item_13)
	if err != nil {
		return err
	}

	return nil
}
