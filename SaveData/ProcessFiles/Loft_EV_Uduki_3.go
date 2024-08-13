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

func Loft_EV_Uduki_3(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_12_Fret_Badge, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_13_Fret_Badge_2, 1)
	if err != nil {
		return err
	}
	//This check is broken because Q14 can't exist beforehand afaik. Don't know if it's really needed.
	/*Q14_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_14_Fret_Battle)
	if err != nil {
		return err
	}
	if !Q14_isCurrent {
		err := db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_14_Fret_Battle, 1)
		if err != nil {
			return err
		}
		//(Not in SERVER_SCRIPT)
		err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
		if err != nil {
			return err
		}
	}*/

	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_14_Fret_Battle, 1)
	if err != nil {
		return err
	}
	//(Not in SERVER_SCRIPT)
	/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}*/

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_14_Chameleon_R, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_14_Chameleon_B, "off")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_14_Chameleon_Y, "off")
		if err != nil {
			return err
		}
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Loft_Uzuki, Consts_LuaHash.Loft_EV_Uduki_4)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q11_Uzuki", 1)
	if err != nil {
		return err
	}

	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_4)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserCharacterItem(UserID, Consts_Character.Fret, Consts_Item.Item_4)
	if err != nil {
		return err
	}

	return nil
}
