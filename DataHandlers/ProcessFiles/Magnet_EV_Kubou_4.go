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

func Magnet_EV_Kubou_4(UserID string) error {
	err := db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_5)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q15_Kubo", 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_16_Nagi_Badge, 1)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_18_Nagi_Badge_3, 1)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_19_Nagi_Battle, 1)
	if err != nil {
		return err
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_R, "off")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_B, "off")
		if err != nil {
			return err
		}

	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_Y, "off")
		if err != nil {
			return err
		}

	}

	_, err = db_commands.CreateUserItem(UserID, Consts_Item.Item_6)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserCharacterItem(UserID, Consts_Character.Nagi, Consts_Item.Item_6)
	if err != nil {
		return err
	}

	return nil
}
