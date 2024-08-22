package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Magnet_BT_Elephant_Lose(UserID string) error {
	//Some if checks
	Q7_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_7_Branch)
	if err != nil {
		return err
	}
	if !Q7_isCurrent {
		return nil
	}

	Q16_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_16_Nagi_Badge)
	if err != nil {
		return err
	}
	if Q16_isCurrent {
		return nil
	}

	Q16_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_16_Nagi_Badge)
	if err != nil {
		return err
	}
	if Q16_isClear {
		return nil
	}

	Q15_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_15_Nagi_Death)
	if err != nil {
		return err
	}
	Q15_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_15_Nagi_Death)
	if err != nil {
		return err
	}

	if !Q15_isClear {
		if !Q15_isCurrent {
			//Set local map
			err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MagnetMap, 7)
			if err != nil {
				return err
			}

			//Create new quest
			err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_15_Nagi_Death, 1)
			if err != nil {
				return err
			}

			//(Not in SERVER_SCRIPT)
			/*err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
			if err != nil {
				return err
			}*/
			//(Not in SERVER_SCRIPT) set the branch quest to not current
			err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_7_Branch, 0)
			if err != nil {
				return err
			}

			//Create Uzuki Pin
			err = db_commands.CreateUserGPSPin(UserID, "Q15_Kubo", Consts_MapPin.Kubo, "", Consts_Coords.Magnet_Kubo_lat, Consts_Coords.Magnet_Kubo_long, Consts_Quest.Quest_15_Nagi_Death, Consts_MapType.MagnetMap, "7")
			if err != nil {
				return err
			}

			//Create quest icon
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_15_Nagi_Death, Consts_QuestItem.Quest_15_Kubo, "off")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
