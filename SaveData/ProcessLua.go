package SaveData

import (
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	"FRPGServer/SaveData/ProcessFiles"
	"fmt"
	"strconv"
)

func ProccessLua(UserID string, LuaHash uint32) error {
	switch LuaHash {
	case Consts_LuaHash.Introduction:
		err := ProcessFiles.ProcessIntroductionLua(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Deai_0:
		err := ProcessFiles.Oioi_8F_EV_Deai_0(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Rindo_0:
		err := ProcessFiles.Oioi_8F_EV_Rindo_0(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Fret_0:
		err := ProcessFiles.Oioi_8F_EV_Fret_0(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Nagi_0:
		err := ProcessFiles.Oioi_8F_EV_Nagi_0(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Shiba_0:
		err := ProcessFiles.Oioi_8F_EV_Shiba_0(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_B1:
		err := ProcessFiles.Oioi_8F_BT_Frog_B1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_R1:
		err := ProcessFiles.Oioi_8F_BT_Frog_R1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Y1:
		err := ProcessFiles.Oioi_8F_BT_Frog_Y1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Bridge:
		err := ProcessFiles.Oioi_8F_BT_Frog_Bridge(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_1:
		err := ProcessFiles.Oioi_8F_BT_Frog_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Shiba_1:
		err := ProcessFiles.Oioi_8F_EV_Shiba_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_B2:
		err := ProcessFiles.Oioi_8F_BT_Frog_B2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_R2:
		err := ProcessFiles.Oioi_8F_BT_Frog_R2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Y2:
		err := ProcessFiles.Oioi_8F_BT_Frog_Y2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Frog3ren_1:
		err := ProcessFiles.Oioi_8F_EV_Frog3ren_1(UserID)
		if err != nil {
			return err
		}
		return nil
	default:
		fmt.Println("Unknown LuaHash to process: " + strconv.FormatUint(uint64(LuaHash), 10))
		return nil
	}
}
