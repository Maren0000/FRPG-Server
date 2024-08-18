package ProcessFiles

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	Consts_ScanType "FRPGServer/Consts/ScanType"
	db_commands "FRPGServer/db/commands"
)

func Magnet_EV_Kubou_3(UserID string) error {
	err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_17_Nagi_Badge_2, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserQuestCurrent(UserID, Consts_Quest.Quest_16_Nagi_Badge, 1)
	if err != nil {
		return err
	}

	err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Magnet_Kubo, Consts_LuaHash.Magnet_EV_Kubou_2)
	if err != nil {
		return err
	}
	err = db_commands.UpdateUserGPSRemove(UserID, "Q15_Kubo", 1)
	if err != nil {
		return err
	}

	//to-do: add tflite
	//3F
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Art_2, 1, Consts_LuaHash.Magnet_EV_Art_3F_Headgear_Error)
	if err != nil {
		return err
	}

	//7F
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Art_3, 1, Consts_LuaHash.Magnet_EV_Art_7F_Signboard_Error)
	if err != nil {
		return err
	}

	//2F
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Art_4, 1, Consts_LuaHash.Magnet_EV_Art_2F_Dragon_Error)
	if err != nil {
		return err
	}

	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Pure", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Pure_lat, Consts_Coords.Magnet_Pure_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "7")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Yuusen", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Yuusen_lat, Consts_Coords.Magnet_Yuusen_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "3")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserGPSPin(UserID, "Magnet_Vari1", Consts_MapPin.MapPin, "", Consts_Coords.Magnet_Vari1_lat, Consts_Coords.Magnet_Vari1_long, Consts_Quest.Quest_16_Nagi_Badge, Consts_MapType.MagnetMap, "2")
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Pure, 1, Consts_LuaHash.Magnet_EV_Pure4_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Yuusen, 1, Consts_LuaHash.Magnet_EV_Yuusen4_1)
	if err != nil {
		return err
	}
	err = db_commands.CreateUserScan(UserID, Consts_ScanType.QR_CODE, Consts_ScanTag.QR_Magnet_Vari1, 1, Consts_LuaHash.Magnet_EV_Vari4_1)
	if err != nil {
		return err
	}
	return nil
}
