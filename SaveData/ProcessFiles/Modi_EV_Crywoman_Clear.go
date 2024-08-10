package ProcessFiles

import (
	Consts_Character "FRPGServer/Consts/Character"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

func Modi_EV_Crywoman_Clear(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_9_Rindo_Badge, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_10_Rindo_Battle, 1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_10_Rindo_Battle, Consts_QuestItem.Quest_10_Hint, "on")
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_CryWoman, Consts_LuaHash.Sys_Error_1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q9_Crywoman", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Modi_Kariya, Consts_LuaHash.Modi_EV_Kariya_3)
	if err != nil {
		return err
	}

	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserCharacterItem(UserID, Consts_Character.Rindo, Consts_Item.Item_2)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSavaColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_10_Rindo_Battle, Consts_QuestItem.Quest_10_Shark_R, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_10_Rindo_Battle, Consts_QuestItem.Quest_10_Shark_B, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_10_Rindo_Battle, Consts_QuestItem.Quest_10_Shark_Y, "off")
		if err != nil {
			return err
		}
	}

	//Not in SERVER_SCRIPT. Mainly to avoid issues later on.
	err = db_commands.UpdateUserResumeLuaResume(UserID, 0)
	if err != nil {
		return err
	}

	return nil
}
