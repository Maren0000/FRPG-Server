package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_BT_Harisenbon_Lose(UserID string) error {
	err := db_commands.UpdateUserOnceEvent(UserID, Consts_LuaHash.Miyashita_EV_Noise_Firstbattle, 1)
	if err != nil {
		return err
	}

	return nil
}
