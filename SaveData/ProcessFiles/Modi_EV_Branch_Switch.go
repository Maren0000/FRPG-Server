package ProcessFiles

import (
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

// To-Do: Chnage quest IDs
func Modi_EV_Branch_Switch(UserID string) error {
	Color, err := db_commands.GetUserSavaColor(UserID)
	if err != nil {
		return err
	}

	Fret_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_1000_Burger)
	if err != nil {
		return err
	}
	if !Fret_Cleared {
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Uzuki, 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Yuusen1, 1)
		if err != nil {
			return err
		}

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_R, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_B, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_Y, 1)
			if err != nil {
				return err
			}
		}
	}

	Nagi_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_1000_Burger)
	if err != nil {
		return err
	}
	if !Nagi_Cleared {
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Uzuki, 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Yuusen1, 1)
		if err != nil {
			return err
		}

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_R, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_B, 1)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_Y, 1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
