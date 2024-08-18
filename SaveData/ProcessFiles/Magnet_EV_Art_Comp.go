package ProcessFiles

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Art_Comp(UserID string) error {
	err := db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_16_Nagi_Badge, 0)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_18_Nagi_Badge_3, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_18_Nagi_Badge_3, Consts_QuestItem.Quest_18_Kubo, "off")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_4)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q15_Kubo", 0)
	if err != nil {
		return err
	}
	return nil
}
