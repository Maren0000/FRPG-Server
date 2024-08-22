package ProcessFiles

import (
	Consts_Character "FRPGServer/Consts/Character"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Oioi_8F_EV_Rindo_0(UserID string) error {
	//Add Quest Progress
	err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
	if err != nil {
		return err
	}

	//Update quest panel
	err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Rindo, "on")
	if err != nil {
		return err
	}

	//Update lua file for scan
	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Rindo, Consts_LuaHash.Oioi_8F_EV_Rindo_1)
	if err != nil {
		return err
	}

	//Add character (Not in SERVER_SCRIPT)
	_, err = db_commands.CreateUserCharacter(UserID, Consts_Character.Rindo, Consts_Item.Item_1)
	if err != nil {
		return err
	}

	return nil
}
