package ProcessFiles

import (
	db_commands "FRPGServer/db/commands"
)

func GPS_EV_Road_6(UserID string) error {
	//Remove GPS
	err := db_commands.UpdateUserGPSRemove(UserID, "GPS_VoiceOnly_6", 1)
	if err != nil {
		return err
	}
	return nil
}
