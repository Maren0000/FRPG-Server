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

func Loft_EV_Yuusen2_Tips(UserID string) error {
	Q14_isCurrent, err := db_commands.CheckCurrentUserQuest(UserID, Consts_Quest.Quest_14_Fret_Battle)
	if err != nil {
		return err
	}

	if Q14_isCurrent {
		err := db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_14_Fret_Battle, Consts_QuestItem.Quest_14_Hint, "on")
		if err != nil {
			return err
		}
	}

	Color, err := db_commands.GetUserSaveColor(UserID)
	if err != nil {
		return err
	}
	if Color == Consts_PlayerColor.Color_Red {
		err = db_commands.CreateUserGPSPin(UserID, "Q14_Chameleon_R", Consts_MapPin.Noise_Chameleon, Consts_MapPinColor.Noise_Chameleon_R, Consts_Coords.Loft_Chameleon_lat, Consts_Coords.Loft_Chameleon_long, Consts_Quest.Quest_7_Branch, Consts_MapType.LoftMap, "6")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Blue {
		err = db_commands.CreateUserGPSPin(UserID, "Q14_Chameleon_B", Consts_MapPin.Noise_Chameleon, Consts_MapPinColor.Noise_Chameleon_B, Consts_Coords.Loft_Chameleon_lat, Consts_Coords.Loft_Chameleon_long, Consts_Quest.Quest_7_Branch, Consts_MapType.LoftMap, "6")
		if err != nil {
			return err
		}
	} else if Color == Consts_PlayerColor.Color_Yellow {
		err = db_commands.CreateUserGPSPin(UserID, "Q14_Chameleon_Y", Consts_MapPin.Noise_Chameleon, Consts_MapPinColor.Noise_Chameleon_Y, Consts_Coords.Loft_Chameleon_lat, Consts_Coords.Loft_Chameleon_long, Consts_Quest.Quest_7_Branch, Consts_MapType.LoftMap, "6")
		if err != nil {
			return err
		}
	}

	return nil
}
