package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	db_commands "FRPGServer/db/commands"
)

func Miyashita_EV_Pure1_2(UserID string) error {
	GPSExists, err := db_commands.CheckUserGPSExists(UserID, "Miyashita_Gorilla")

	if !GPSExists {
		err = db_commands.CreateUserGPSPin(UserID, "Miyashita_Gorilla", Consts_MapPin.Noise_Gorilla, "", Consts_Coords.Miyashita_Gorilla_lat, Consts_Coords.Miyashita_Gorilla_long, Consts_Quest.Quest_22_Last_Boss, Consts_MapType.MiyashitaParkMap, "4")
		if err != nil {
			return err
		}
	}

	return nil
}
