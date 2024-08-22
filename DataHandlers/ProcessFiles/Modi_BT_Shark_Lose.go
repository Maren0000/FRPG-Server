package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Modi_BT_Shark_Lose(UserID string) error {
	//Some if checks
	Q7_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_7_Branch)
	if err != nil {
		return err
	}
	if !Q7_isCurrent {
		return nil
	}

	Q9_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_9_Rindo_Badge)
	if err != nil {
		return err
	}
	if Q9_isCurrent {
		return nil
	}

	Q9_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_9_Rindo_Badge)
	if err != nil {
		return err
	}
	if Q9_isClear {
		return nil
	}

	Q8_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_8_Rindo_Death)
	if err != nil {
		return err
	}
	Q8_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_8_Rindo_Death)
	if err != nil {
		return err
	}

	if !Q8_isClear {
		if !Q8_isCurrent {
			//Set local map
			err = db_commands.SetUserLocalMap(UserID, Consts_MapType.ModiMap, 7)
			if err != nil {
				return err
			}

			//Create new quest
			err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_8_Rindo_Death, 1)
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

			//Create Kariya Pin
			err = db_commands.CreateUserGPSPin(UserID, "Q8_Kariya", Consts_MapPin.Kariya, "", Consts_Coords.Modi_Kariya_lat, Consts_Coords.Modi_Kariya_long, Consts_Quest.Quest_8_Rindo_Death, Consts_MapType.ModiMap, "7")
			if err != nil {
				return err
			}

			//Create quest icon
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_8_Rindo_Death, Consts_QuestItem.Quest_8_Kariya, "off")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
