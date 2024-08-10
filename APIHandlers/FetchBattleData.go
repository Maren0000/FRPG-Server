package APIHandlers

import (
	Consts_Battle "FRPGServer/Consts/Battle"
	Consts_Item "FRPGServer/Consts/Item"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_Noise "FRPGServer/Consts/Noise"
	"strconv"
)

func FetchBattleData(BattleID int) (BattleData BattleInfoModel) {
	//To-Do: Randomize music for each battle (except maybe bosses?)
	BattleData.BGM_ID = 1

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
	}

	return BattleData
}

func FetchBattleLuaResult(BattleID int, Success int, NowHP int) uint32 {
	//To-Do: This may be removed?
	if NowHP == 0 {
		return Consts_LuaHash.Sys_BT_Lose
	}

	switch BattleID {
	case Consts_Battle.Oioi_5F_Crow:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Crow_Win
		} else {
			return 0
		}
	case Consts_Battle.Oioi_5F_Wolf:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Wolf_Win
		} else {
			return 0
		}
	case Consts_Battle.Oioi_5F_Jellyfish:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Jellyfish_Win
		} else {
			return 0
		}
	case Consts_Battle.Oioi_5F_Bear:
		if Success == 1 {
			return Consts_LuaHash.Oioi_5F_BT_Bear_Win
		} else {
			return 0
		}
	case Consts_Battle.Modi_Shark:
		if Success == 1 {
			return Consts_LuaHash.Modi_EV_Mission_Clear
		} else {
			return Consts_LuaHash.Modi_BT_Shark_Lose
		}
	default:
		return 0
	}
}
