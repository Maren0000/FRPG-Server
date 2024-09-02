package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func EX_EV_Nazo_Start(UserID string) error {
	Q108_Clear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_108_EX_Nazo)
	if err != nil {
		return err
	}
	if !Q108_Clear {
		err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_108_EX_Nazo, 1)
		if err != nil {
			return err
		}
	}

	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_120_EX_Nazo_Playing, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_EX_Nazo, Consts_LuaHash.Miyashita_EV_Mission_Cleared)
	if err != nil {
		return err
	}

	//To-Do: Add tflite
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_EX_Nazo_Goal, 1, Consts_LuaHash.EX_EV_Nazo_Goal)
	if err != nil {
		return err
	}

	//GPS pin should be remade here but honestly no point
	//Update: Past me was dum
	err = db_commands.UpdateUserGPSQuest(UserID, "GPS_EX_Parco", Consts_Quest.Quest_120_EX_Nazo_Playing)
	if err != nil {
		return err
	}

	return nil
}
