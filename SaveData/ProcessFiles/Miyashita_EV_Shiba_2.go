package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Shiba_2(UserID string) error {

	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_21_Last_Battle, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_22_Last_Boss, 1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_22_Last_Boss, Consts_QuestItem.Quest_22_Boss, "off")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Miyashita_Gorilla, 1, Consts_LuaHash.Miyashita_EV_Gorilla_Firstbattle)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Badge, Consts_LuaHash.Miyashita_EV_Mission_Cleared)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Kariya", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Kariya, Consts_LuaHash.Miyashita_EV_Kariya_3)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Kariya", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Kariya, Consts_LuaHash.Miyashita_EV_Kariya_3)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Uzuki", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Uzuki, Consts_LuaHash.Miyashita_EV_Uduki_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Kanon", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Kanon, Consts_LuaHash.Miyashita_EV_Kanon_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Pure", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Pure_1, Consts_LuaHash.Miyashita_EV_Pure1_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Vari_1", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Vari_1, Consts_LuaHash.Miyashita_EV_Vari2_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Vari_2", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Vari_2, Consts_LuaHash.Miyashita_EV_Vari3_2)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserGPSRemove(UserID, "Miyashita_Yuusen", 0)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Miyashita_Yuusen_1, Consts_LuaHash.Miyashita_EV_Yuusen1_2)
	if err != nil {
		return err
	}
	return nil
}
