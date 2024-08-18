package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapPinColor "FRPGServer/Consts/MapPinColor"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Vari1_Tips(UserID string) error {
	Q19_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_19_Nagi_Battle)
	if err != nil {
		return err
	}

	if Q19_isCurrent {
		err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Hint, "on")
		if err != nil {
			return err
		}
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserGPSPin(UserID, "Q19_Elephant_R", Consts_MapPin.Noise_Elephant, Consts_MapPinColor.Noise_Elephant_R, Consts_Coords.Magnet_Elephant_lat, Consts_Coords.Magnet_Elephant_long, Consts_Quest.Quest_7_Branch, Consts_MapType.MagnetMap, "6")
		if err != nil {
			return err
		}
		//Handled by the Kubo files anyways
		/*err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_R, "off")
		if err != nil {
			return err
		}*/

	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserGPSPin(UserID, "Q19_Elephant_B", Consts_MapPin.Noise_Elephant, Consts_MapPinColor.Noise_Elephant_B, Consts_Coords.Magnet_Elephant_lat, Consts_Coords.Magnet_Elephant_long, Consts_Quest.Quest_7_Branch, Consts_MapType.MagnetMap, "6")
		if err != nil {
			return err
		}
		//Handled by the Kubo files anyways
		/*err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_B, "off")
		if err != nil {
			return err
		}*/

	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserGPSPin(UserID, "Q19_Elephant_Y", Consts_MapPin.Noise_Elephant, Consts_MapPinColor.Noise_Elephant_Y, Consts_Coords.Magnet_Elephant_lat, Consts_Coords.Magnet_Elephant_long, Consts_Quest.Quest_7_Branch, Consts_MapType.MagnetMap, "6")
		if err != nil {
			return err
		}
		//Handled by the Kubo files anyways
		/*err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_19_Nagi_Battle, Consts_QuestItem.Quest_19_Elephant_Y, "off")
		if err != nil {
			return err
		}*/

	}

	return nil
}
