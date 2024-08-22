package DataHandlers

import (
	Consts_BGM "FRPGServer/Consts/BGM"
	Consts_Battle "FRPGServer/Consts/Battle"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Noise "FRPGServer/Consts/Noise"
	"FRPGServer/Models"
	"FRPGServer/Utils"
	"strconv"
)

var Battle_BGM_List = []int{Consts_BGM.BGM_002, Consts_BGM.BGM_003, Consts_BGM.BGM_004, Consts_BGM.BGM_005, Consts_BGM.BGM_013, Consts_BGM.BGM_014, Consts_BGM.BGM_018, Consts_BGM.BGM_031, Consts_BGM.BGM_033, Consts_BGM.BGM_048, Consts_BGM.BGM_049}

func FetchBattleData(BattleID int) (BattleData Models.BattleInfoModel) {
	BattleData.BGM_ID = Utils.RandomBGM(Battle_BGM_List)

	switch BattleID {
	case Consts_Battle.Oioi_5F_Crow:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Crow
		BattleData.NoiseID = Consts_Noise.NoiseID_Crow
		BattleData.Damage = 1
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_3)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Oioi_5F_Wolf:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Wolf
		BattleData.NoiseID = Consts_Noise.NoiseID_Wolf
		BattleData.Damage = 1
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_5)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Oioi_5F_Jellyfish:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Jellyfish
		BattleData.NoiseID = Consts_Noise.NoiseID_Jellyfish
		BattleData.Damage = 1
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_1)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Oioi_5F_Bear:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Bear
		BattleData.NoiseID = Consts_Noise.NoiseID_Bear
		BattleData.Damage = 1
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_7)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Modi_Shark:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Shark
		BattleData.NoiseID = Consts_Noise.NoiseID_Shark
		BattleData.Damage = 2
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_2)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Loft_Chameleon:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Chameleon
		BattleData.NoiseID = Consts_Noise.NoiseID_Chameleon
		BattleData.Damage = 2
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_4)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Magnet_Elephant:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Elephant
		BattleData.NoiseID = Consts_Noise.NoiseID_Elephant
		BattleData.Damage = 2
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_6)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Miyashita_Tyranno:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Tyranno
		BattleData.NoiseID = Consts_Noise.NoiseID_Tyranno
		BattleData.Damage = 3
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_8)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Miyashita_Harisenbon:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Harisenbon
		BattleData.NoiseID = Consts_Noise.NoiseID_Harisenbon
		BattleData.Damage = 3
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_2) + "," + strconv.Itoa(Consts_Item.Item_4)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Miyashita_Scorpion:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Scorpion
		BattleData.NoiseID = Consts_Noise.NoiseID_Scorpion
		BattleData.Damage = 3
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_6) + "," + strconv.Itoa(Consts_Item.Item_8)
		BattleData.BIgnoreInputOrder = 0
	case Consts_Battle.Miyashita_Gorilla:
		BattleData.NoiseSymbolPath = Consts_Noise.NoiseSymbol_Gorilla
		BattleData.NoiseID = Consts_Noise.NoiseID_Gorilla
		BattleData.Damage = 4
		BattleData.Badge = strconv.Itoa(Consts_Item.Item_2) + "," + strconv.Itoa(Consts_Item.Item_8) + "," + strconv.Itoa(Consts_Item.Item_6) + "," + strconv.Itoa(Consts_Item.Item_4)
		BattleData.BIgnoreInputOrder = 0
		BattleData.BGM_ID = Consts_BGM.BGM_042
	}

	return BattleData
}

func FetchBattleLuaResult(BattleID int, Success int, NowHP int) uint32 {

	switch BattleID {
	case Consts_Battle.Oioi_5F_Crow:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Crow_Win
		} else {
			return Consts_LuaHash.Sys_BT_Lose
		}
	case Consts_Battle.Oioi_5F_Wolf:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Wolf_Win
		} else {
			return Consts_LuaHash.Sys_BT_Lose
		}
	case Consts_Battle.Oioi_5F_Jellyfish:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Jellyfish_Win
		} else {
			return Consts_LuaHash.Sys_BT_Lose
		}
	case Consts_Battle.Oioi_5F_Bear:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Bear_Win
		} else {
			return Consts_LuaHash.Sys_BT_Lose
		}
	case Consts_Battle.Modi_Shark:
		if Success == 1 {
			return Consts_LuaHash.Modi_EV_Mission_Clear
		} else {
			return Consts_LuaHash.Modi_BT_Shark_Lose
		}
	case Consts_Battle.Loft_Chameleon:
		if Success == 1 {
			return Consts_LuaHash.Loft_EV_Mission_Clear
		} else {
			return Consts_LuaHash.Loft_BT_Chameleon_Lose
		}
	case Consts_Battle.Magnet_Elephant:
		if Success == 1 {
			return Consts_LuaHash.Magnet_EV_Mission_Clear
		} else {
			return Consts_LuaHash.Magnet_BT_Elephant_Lose
		}
	case Consts_Battle.Miyashita_Tyranno:
		if Success == 1 {
			return Consts_LuaHash.Miyashita_EV_Tyranno_Win
		} else {
			return Consts_LuaHash.Miyashita_BT_Tyranno_Lose
		}
	case Consts_Battle.Miyashita_Harisenbon:
		if Success == 1 {
			return Consts_LuaHash.Miyashita_EV_Harisenbon_Win
		} else {
			return Consts_LuaHash.Miyashita_BT_Harisenbon_Lose
		}
	case Consts_Battle.Miyashita_Scorpion:
		if Success == 1 {
			return Consts_LuaHash.Miyashita_EV_Scorpion_Win
		} else {
			return Consts_LuaHash.Miyashita_BT_Scorpion_Lose
		}
	case Consts_Battle.Miyashita_Gorilla:
		if Success == 1 {
			return Consts_LuaHash.Miyashita_EV_Gorilla_Win
		} else {
			return Consts_LuaHash.Miyashita_BT_Gorilla_Lose
		}
	default:
		return 0
	}
}
