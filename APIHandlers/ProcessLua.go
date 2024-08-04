package APIHandlers

import (
	Consts_Coords "FRPGServer/Consts/Coords"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_MapPinColor "FRPGServer/Consts/MapPinColor"
	Consts_MapType "FRPGServer/Consts/MapType"
	Consts_PlayerColor "FRPGServer/Consts/PlayerColor"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	db_commands "FRPGServer/db/commands"
	"fmt"
	"strconv"
)

func proccessLua(UserID string, LuaHash uint32) error {
	//Welcome to hell
	switch LuaHash {
	case Consts_LuaHash.Introduction:
		err := db_commands.CreateUserOnceEvent(UserID, Consts_LuaHash.Introduction, 0)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserSaveIntro(UserID, 0)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Deai_0:
		//Clear and start new quest
		err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_1_Start, 1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_2_Start_Twisters)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
		if err != nil {
			fmt.Println(err)
		}

		//Add items for quest panel
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Rindo, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Fret, "off")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Nagi, "off")
		if err != nil {
			return err
		}

		//Update Scan Tag
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q1_JoinBadge, Consts_LuaHash.Sys_Error_1)
		if err != nil {
			return err
		}

		//Add new scan tags
		err = db_commands.CreateUserScan(UserID, 2, 3, Consts_ScanTag.QR_Q2_Rindo, 1, Consts_LuaHash.Oioi_8F_EV_Rindo_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 3, 3, Consts_ScanTag.QR_Q2_Fret, 1, Consts_LuaHash.Oioi_8F_EV_Fret_0)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserScan(UserID, 4, 3, Consts_ScanTag.QR_Q2_Nagi, 1, Consts_LuaHash.Oioi_8F_EV_Nagi_0)
		if err != nil {
			return err
		}

		//Update Local Map
		err = db_commands.SetUserLocalMap(UserID, Consts_MapType.MaruiMap, 8)
		if err != nil {
			return err
		}

		//Update GPS pins
		err = db_commands.CreateUserGPSPin(UserID, "2", "Q2_Rindo", Consts_MapPin.Rindo, "", Consts_Coords.Oioi_8F_Rindo_lat, Consts_Coords.Oioi_8F_Rindo_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "3", "Q2_Fret", Consts_MapPin.Fret, "", Consts_Coords.Oioi_8F_Fret_lat, Consts_Coords.Oioi_8F_Fret_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "4", "Q2_Nagi", Consts_MapPin.Nagi, "", Consts_Coords.Oioi_8F_Nagi_lat, Consts_Coords.Oioi_8F_Nagi_long, Consts_Quest.Quest_2_Start_Twisters, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}

		return nil
	case Consts_LuaHash.Oioi_8F_EV_Rindo_0:
		//Add Quest Progress
		err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
		if err != nil {
			return err
		}

		//Update quest panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Rindo, "on")
		if err != nil {
			return err
		}

		//Update lua file for scan
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Rindo, Consts_LuaHash.Oioi_8F_EV_Rindo_1)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Fret_0:
		//Add Quest Progress
		err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
		if err != nil {
			return err
		}

		//Update quest panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Fret, "on")
		if err != nil {
			return err
		}

		//Update lua file for scan
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Fret, Consts_LuaHash.Oioi_8F_EV_Fret_1)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Nagi_0:
		//Add Quest Progress
		err := db_commands.UpdateUserQuestProgress(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
		if err != nil {
			return err
		}

		//Update quest panel
		err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_2_Start_Twisters, Consts_QuestItem.Quest_2_Nagi, "on")
		if err != nil {
			return err
		}

		//Update lua file for scan
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Nagi, Consts_LuaHash.Oioi_8F_EV_Nagi_1)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Shiba_0:
		//Clear and start new quest
		err := db_commands.UpdateUserQuestClear(UserID, Consts_Quest.Quest_2_Start_Twisters, 1)
		if err != nil {
			return err
		}
		err = db_commands.CreateUserQuest(UserID, Consts_Quest.Quest_3)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserSaveNewQuest(UserID, 1)
		if err != nil {
			fmt.Println(err)
		}

		//Update lua files for scans
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Rindo, Consts_LuaHash.Oioi_8F_EV_Rindo_Def)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Fret, Consts_LuaHash.Oioi_8F_EV_Fret_Def)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserScanLuaHash(UserID, Consts_ScanTag.QR_Q2_Nagi, Consts_LuaHash.Oioi_8F_EV_Nagi_Def)
		if err != nil {
			return err
		}

		//Remove Current GPS markers
		err = db_commands.UpdateUserGPSRemove(UserID, "2", 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserGPSRemove(UserID, "3", 1)
		if err != nil {
			return err
		}
		err = db_commands.UpdateUserGPSRemove(UserID, "4", 1)
		if err != nil {
			return err
		}

		//Create new GPS for team members
		err = db_commands.CreateUserGPSPin(UserID, "5", "Q3_Pure_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Pure_1_lat, Consts_Coords.Oioi_8F_Pure_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "6", "Q3_Vari_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Vari_1_lat, Consts_Coords.Oioi_8F_Vari_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}
		err = db_commands.CreateUserGPSPin(UserID, "7", "Q3_Yuusen_1", Consts_MapPin.MapPin, "", Consts_Coords.Oioi_8F_Yuusen_1_lat, Consts_Coords.Oioi_8F_Yuusen_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
		if err != nil {
			return err
		}

		//Now the fun part of if checks
		Color, err := db_commands.GetUserSavaColor(UserID)
		if err != nil {
			return err
		}
		if Color == Consts_PlayerColor.Color_Red {
			//Create QR scans
			err = db_commands.CreateUserScan(UserID, 5, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R1)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 6, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_R2)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 7, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 8, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 9, 3, Consts_ScanTag.QR_Q4_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 10, 3, Consts_ScanTag.QR_Q4_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 11, 3, Consts_ScanTag.QR_Q4_Frog_B_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 12, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 13, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 14, 3, Consts_ScanTag.QR_Q4_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 15, 3, Consts_ScanTag.QR_Q4_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 16, 3, Consts_ScanTag.QR_Q4_Frog_Y_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}

			//Add items to quest panel
			err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_R_1, "off")
			if err != nil {
				return err
			}
			err = db_commands.UpdateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_R_2, "off")
			if err != nil {
				return err
			}

			//Add GPS markers
			err = db_commands.CreateUserGPSPin(UserID, "8", "Q3_Frog_R1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
			err = db_commands.CreateUserGPSPin(UserID, "9", "Q3_Frog_R2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_R, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Blue {
			//Create QR scans
			err = db_commands.CreateUserScan(UserID, 5, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B1)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 6, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_B2)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 7, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 8, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 9, 3, Consts_ScanTag.QR_Q4_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 10, 3, Consts_ScanTag.QR_Q4_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 11, 3, Consts_ScanTag.QR_Q4_Frog_R_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 12, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 13, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 14, 3, Consts_ScanTag.QR_Q4_Frog_Y_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 15, 3, Consts_ScanTag.QR_Q4_Frog_Y_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 16, 3, Consts_ScanTag.QR_Q4_Frog_Y_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}

			//Add items to quest panel
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_B_1, "off")
			if err != nil {
				return err
			}
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_B_2, "off")
			if err != nil {
				return err
			}

			//Add GPS markers
			err = db_commands.CreateUserGPSPin(UserID, "8", "Q3_Frog_B1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
			err = db_commands.CreateUserGPSPin(UserID, "9", "Q3_Frog_B2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_B, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
		} else if Color == Consts_PlayerColor.Color_Yellow {
			//Create QR scans
			err = db_commands.CreateUserScan(UserID, 5, 3, Consts_ScanTag.QR_Q3_Frog_Y_1, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y1)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 6, 3, Consts_ScanTag.QR_Q3_Frog_Y_2, 1, Consts_LuaHash.Oioi_8F_BT_Frog_Y2)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 7, 3, Consts_ScanTag.QR_Q3_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 8, 3, Consts_ScanTag.QR_Q3_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 9, 3, Consts_ScanTag.QR_Q4_Frog_R_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 10, 3, Consts_ScanTag.QR_Q4_Frog_R_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 11, 3, Consts_ScanTag.QR_Q4_Frog_R_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 12, 3, Consts_ScanTag.QR_Q3_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 13, 3, Consts_ScanTag.QR_Q3_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 14, 3, Consts_ScanTag.QR_Q4_Frog_B_1, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 15, 3, Consts_ScanTag.QR_Q4_Frog_B_2, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}
			err = db_commands.CreateUserScan(UserID, 16, 3, Consts_ScanTag.QR_Q4_Frog_B_3, 1, Consts_LuaHash.Sys_Noise_Error_0)
			if err != nil {
				return err
			}

			//Add items to quest panel
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_Y_1, "off")
			if err != nil {
				return err
			}
			err = db_commands.CreateUserQuestItem(UserID, Consts_Quest.Quest_3, Consts_QuestItem.Quest_3_Frog_Y_2, "off")
			if err != nil {
				return err
			}

			//Add GPS markers
			err = db_commands.CreateUserGPSPin(UserID, "8", "Q3_Frog_Y1", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_1_lat, Consts_Coords.Oioi_8F_Frog_1_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
			err = db_commands.CreateUserGPSPin(UserID, "9", "Q3_Frog_Y2", Consts_MapPin.Noise_Frog, Consts_MapPinColor.Noise_Frog_Y, Consts_Coords.Oioi_8F_Frog_2_lat, Consts_Coords.Oioi_8F_Frog_2_long, Consts_Quest.Quest_3, Consts_MapType.MaruiMap, "8")
			if err != nil {
				return err
			}
		}
		return nil
	default:
		fmt.Println("Unknown LuaHash to process: " + strconv.FormatUint(uint64(LuaHash), 10))
		return nil
	}
}

func fetchSaveData(UserID string) (save PlayerDataModel, bIntro bool, error error) {
	PlayerCharacters, err := db_commands.ListUserCharacters(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerCharacters {
		var temp CharacterModel
		temp.CharacterId = int(data.CharacterId.Int64)
		temp.ItemId = int(data.ItemId.Int64)

		save.ACharacter = append(save.ACharacter, temp)
	}

	PlayerScans, err := db_commands.ListUserScan(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerScans {
		var temp ScanModel
		temp.ID = int(data.ID.Int64)
		temp.Type = int(data.Type.Int64)
		temp.Tag = data.Tag.String
		temp.Height = float32(data.Height.Float64)
		if data.BMulti.Int64 == 0 {
			temp.BMulti = false
		} else {
			temp.BMulti = true
		}

		save.AScan = append(save.AScan, temp)
	}

	PlayerBuildings, err := db_commands.ListUserBuilding(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerBuildings {
		var temp BuildingModel
		temp.Name = data.Name.String
		temp.Prefab = data.Prefab.String
		temp.Status = data.Status.String

		save.ABuildings = append(save.ABuildings, temp)
	}

	PlayerQuest, err := db_commands.GetCurrentUserQuest(UserID)
	if err != nil {
		return save, false, err
	}
	var CurrentQuest QuestModel
	CurrentQuest.ID = int(PlayerQuest.ID.Int64)
	CurrentQuest.Value = int(PlayerQuest.Value.Int64)

	PlayerQuestItems, err := db_commands.ListUserQuestItems(UserID, int(PlayerQuest.ID.Int64))
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerQuestItems {
		var temp QuestItemModel
		temp.Name = data.Name.String
		temp.Type = data.Type.String

		CurrentQuest.ItemList = append(CurrentQuest.ItemList, temp)
	}

	PlayerGPS, err := db_commands.ListUserGPS(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerGPS {
		var temp GPSModel
		temp.ID = data.ID.String
		temp.Name = data.Name.String
		temp.PinType = data.PinType.String
		temp.PinColor = data.PinColor.String
		temp.Latitude = data.Latitude.Float64
		temp.Longitude = data.Longitude.Float64
		temp.BLocationEvent = int(data.BLocationEvent.Int64)
		temp.LuaScript = uint32(data.LuaScript.Int64)
		temp.Quest = &CurrentQuest
		temp.MapType = data.MapType.String
		temp.MapNo = data.MapNo.String

		save.AGPS = append(save.AGPS, temp)
	}

	PlayerOnceEvent, err := db_commands.ListUserOnceEvent(UserID)
	if err != nil {
		return save, false, err
	}
	var OnceEventTemp []uint32
	for _, data := range PlayerOnceEvent {
		OnceEventTemp = append(OnceEventTemp, uint32(data.UInt.Int64))
	}
	save.AOnceEvent = OnceEventTemp

	var RemovedScansTemp []string
	PlayerRemovedScans, err := db_commands.ListUserScanRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedScans {
		RemovedScansTemp = append(RemovedScansTemp, data.Tag.String)
	}
	save.ARemoveScan = RemovedScansTemp

	var RemovedGPSTemp []string
	PlayerRemovedGPS, err := db_commands.ListUserGPSRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedGPS {
		RemovedGPSTemp = append(RemovedGPSTemp, data.ID.String) //Might be ID idk
	}
	save.ARemoveGPS = RemovedGPSTemp

	var RemovedOnceEventTemp []uint32
	PlayerRemovedOnceEvent, err := db_commands.ListUserOnceEventRemoved(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerRemovedOnceEvent {
		RemovedOnceEventTemp = append(RemovedOnceEventTemp, uint32(data.UInt.Int64))
	}
	save.ARemoveOnceEvent = RemovedOnceEventTemp

	var itemTemp []int
	PlayerItems, err := db_commands.ListUserItems(UserID)
	if err != nil {
		return save, false, err
	}
	for _, data := range PlayerItems {
		itemTemp = append(itemTemp, int(data.Item.Int64))
	}
	save.AItemList = itemTemp

	PlayerCurrentLocalMap, err := db_commands.GetCurrentUserLocalMap(UserID)
	if err != nil {
		return save, false, err
	}
	var LocalmapTemp LocalMapModel
	LocalmapTemp.Name = PlayerCurrentLocalMap.Name.String
	LocalmapTemp.Floor = int(PlayerCurrentLocalMap.Floor.Int64)
	save.LocalMap = &LocalmapTemp

	PlayerSave, err := db_commands.GetUserSavaData(UserID)
	if err != nil {
		return save, false, err
	}

	save.NowHp = int(PlayerSave.NowHP.Int64)
	save.MaxHp = int(PlayerSave.MaxHP.Int64)
	save.ColorId = int(PlayerSave.ColorID.Int64)
	save.Quest = &CurrentQuest
	if PlayerSave.BNewQuest.Int64 == 0 {
		save.BNewQuest = false
	} else {
		save.BNewQuest = true
	}

	if PlayerSave.BIntro.Int64 == 0 {
		return save, false, nil
	} else {
		return save, true, nil
	}
}
