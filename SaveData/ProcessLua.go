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
	case Consts_LuaHash.Oioi_8F_BT_Frog_B3:
		err := ProcessFiles.Oioi_8F_BT_Frog_B3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_R3:
		err := ProcessFiles.Oioi_8F_BT_Frog_R3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Y3:
		err := ProcessFiles.Oioi_8F_BT_Frog_Y3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog3ren_Bridge:
		err := ProcessFiles.Oioi_8F_BT_Frog3ren_Bridge(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog3ren_1:
		err := ProcessFiles.Oioi_8F_BT_Frog3ren_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_B4:
		err := ProcessFiles.Oioi_8F_BT_Frog_B4(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_R4:
		err := ProcessFiles.Oioi_8F_BT_Frog_R4(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Y4:
		err := ProcessFiles.Oioi_8F_BT_Frog_Y4(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog3ren_2:
		err := ProcessFiles.Oioi_8F_BT_Frog3ren_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_B5:
		err := ProcessFiles.Oioi_8F_BT_Frog_B5(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_R5:
		err := ProcessFiles.Oioi_8F_BT_Frog_R5(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_BT_Frog_Y5:
		err := ProcessFiles.Oioi_8F_BT_Frog_Y5(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_8F_EV_Shiba_2:
		err := ProcessFiles.Oioi_8F_EV_Shiba_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Wall_1:
		err := ProcessFiles.Oioi_5F_EV_Wall_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Vari3_1:
		err := ProcessFiles.Oioi_5F_EV_Vari3_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Pure3_1:
		err := ProcessFiles.Oioi_5F_EV_Pure3_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Pure2_1:
		err := ProcessFiles.Oioi_5F_EV_Pure2_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Yuusen3_1:
		err := ProcessFiles.Oioi_5F_EV_Yuusen3_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_BT_Crow_Win:
		err := ProcessFiles.Oioi_5F_BT_Crow_Win(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_BT_Wolf_Win:
		err := ProcessFiles.Oioi_5F_BT_Wolf_Win(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_BT_Jellyfish_Win:
		err := ProcessFiles.Oioi_5F_BT_Jellyfish_Win(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_BT_Bear_Win:
		err := ProcessFiles.Oioi_5F_BT_Bear_Win(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Oioi_5F_EV_Shiba_0:
		err := ProcessFiles.Oioi_5F_EV_Shiba_0(UserID)
		if err != nil {
			return err
		}
		return nil

	case Consts_LuaHash.GPS_1:
		err := ProcessFiles.GPS_EV_Road_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.GPS_2:
		err := ProcessFiles.GPS_EV_Road_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.GPS_3:
		err := ProcessFiles.GPS_EV_Road_3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.GPS_4:
		err := ProcessFiles.GPS_EV_Road_4(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.GPS_5:
		err := ProcessFiles.GPS_EV_Road_5(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.GPS_6:
		err := ProcessFiles.GPS_EV_Road_6(UserID)
		if err != nil {
			return err
		}
		return nil

	case Consts_LuaHash.Sys_EV_Branch_Switch_EX_On:
		err := ProcessFiles.Sys_EV_Branch_Switch_EX_On(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Sys_EV_Burger_FirstDamage:
		err := ProcessFiles.Sys_EV_Burger_FirstDamage(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Sys_EV_Burger_Refresh:
		err := ProcessFiles.Sys_EV_Burger_Refresh(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Sys_BT_Lose:
		err := ProcessFiles.Sys_BT_Lose(UserID)
		if err != nil {
			return err
		}
		return nil
	default:
		fmt.Println("Unknown LuaHash to process: " + strconv.FormatUint(uint64(LuaHash), 10))
		return nil
	}
}
