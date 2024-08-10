package ProcessFiles

import (
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
)

// To-Do: Fix Quest IDs
func Sys_EV_Branch_Switch(UserID string) error {
	Color, err := db_commands.GetUserSavaColor(UserID)
	if err != nil {
		return err
	}

	Q10_Clear, err := db_commands.CheckUserQuestExists(UserID, Consts_Quest.Quest_10_Rindo_Battle)
	if err != nil {
		return err
	}
	if !Q10_Clear {
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Modi_Kariya, 0)

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Modi_Shark_R, 0)
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Modi_Shark_B, 0)
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Modi_Shark_Y, 0)
		}
	}

	Fret_Cleared, err := db_commands.CheckUserQuestCleared(UserID, Consts_Quest.Quest_1000_Burger)
	if err != nil {
		return err
	}
	if !Fret_Cleared {
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Uzuki, 0)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Yuusen1, 0)
		if err != nil {
			return err
		}

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_R, 0)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_B, 0)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_Y, 0)
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
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Uzuki, 0)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Yuusen1, 0)
		if err != nil {
			return err
		}

		if Color == Consts_PlayerColor.Color_Red {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_R, 0)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_B, 0)
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			err = db_commands.UpdateUserScanRemove(UserID, Consts_ScanTag.QR_Q_Chameleon_Y, 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
