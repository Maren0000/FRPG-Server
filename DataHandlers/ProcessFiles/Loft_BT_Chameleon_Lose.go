package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Loft_BT_Chameleon_Lose(UserID string) error {
	//Some if checks
	Q7_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_7_Branch)
	if err != nil {
		return err
	}
	if !Q7_isCurrent {
		return nil
	}

	Q12_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_12_Fret_Badge)
	if err != nil {
		return err
	}
	if Q12_isCurrent {
		return nil
	}

	Q12_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_12_Fret_Badge)
	if err != nil {
		return err
	}
	if Q12_isClear {
		return nil
	}

	Q11_isClear, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_11_Fret_Death)
	if err != nil {
		return err
	}
	Q11_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_11_Fret_Death)
	if err != nil {
		return err
	}

	if !Q11_isClear {
		if !Q11_isCurrent {
			//Set local map
			err = db_commands.SetUserLocalMap(UserID, Consts_MapType.LoftMap, 6)
			if err != nil {
				return err
			}

			//Create new quest
			err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_11_Fret_Death, 1)
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
			err = db_commands.CreateUserGPSPin(UserID, "Q11_Uzuki", Consts_MapPin.Uzuki, "", Consts_Coords.Loft_Uzuki_lat, Consts_Coords.Loft_Uzuki_long, Consts_Quest.Quest_11_Fret_Death, Consts_MapType.LoftMap, "6")
			if err != nil {
				return err
			}

			//Create quest icon
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_11_Fret_Death, Consts_QuestItem.Quest_11_Uzuki, "off")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
