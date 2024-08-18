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

	case Consts_LuaHash.Modi_BT_Shark_Lose:
		err := ProcessFiles.Modi_BT_Shark_Lose(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Kariya_1:
		err := ProcessFiles.Modi_EV_Kariya_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Branch_Switch:
		err := ProcessFiles.Modi_EV_Branch_Switch(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Crywoman_1:
		err := ProcessFiles.Modi_EV_Crywoman_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Girl_2:
		err := ProcessFiles.Modi_EV_Girl_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Mpanel_KW_1:
		err := ProcessFiles.Modi_EV_Mpanel_KW_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Mpanel_KW_10:
		err := ProcessFiles.Modi_EV_Mpanel_KW_10(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Mpanel_KW_11:
		err := ProcessFiles.Modi_EV_Mpanel_KW_11(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Crywoman_Miss:
		err := ProcessFiles.Modi_EV_Crywoman_Miss(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Kanon_2:
		err := ProcessFiles.Modi_EV_Kanon_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Crywoman_Clear:
		err := ProcessFiles.Modi_EV_Crywoman_Clear(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Modi_EV_Mission_Clear:
		err := ProcessFiles.Modi_EV_Mission_Clear(UserID)
		if err != nil {
			return err
		}
		return nil

	case Consts_LuaHash.Loft_BT_Chameleon_Lose:
		err := ProcessFiles.Loft_BT_Chameleon_Lose(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Uduki_1:
		err := ProcessFiles.Loft_EV_Uduki_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Yuusen2_Tips:
		err := ProcessFiles.Loft_EV_Yuusen2_Tips(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_3F_Book:
		err := ProcessFiles.Loft_EV_3F_Book(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_3F_Cup:
		err := ProcessFiles.Loft_EV_3F_Cup(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_3F_Bottle:
		err := ProcessFiles.Loft_EV_3F_Bottle(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_4F_Clock:
		err := ProcessFiles.Loft_EV_4F_Clock(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_4F_Holder:
		err := ProcessFiles.Loft_EV_4F_Holder(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_4F_Radio:
		err := ProcessFiles.Loft_EV_4F_Radio(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_5F_Bag:
		err := ProcessFiles.Loft_EV_5F_Bag(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_5F_Lanthanum:
		err := ProcessFiles.Loft_EV_5F_Lanthanum(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_5F_Umbrella:
		err := ProcessFiles.Loft_EV_5F_Umbrella(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_6F_Balloon:
		err := ProcessFiles.Loft_EV_6F_Balloon(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_6F_Puzzle:
		err := ProcessFiles.Loft_EV_6F_Puzzle(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_6F_Stuffed:
		err := ProcessFiles.Loft_EV_6F_Stuffed(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Branch_Switch:
		err := ProcessFiles.Loft_EV_Branch_Switch(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Item_Comp:
		err := ProcessFiles.Loft_EV_Item_Comp(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Mission_Clear:
		err := ProcessFiles.Loft_EV_Mission_Clear(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Uduki_3:
		err := ProcessFiles.Loft_EV_Uduki_3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Loft_EV_Remaining:
		err := ProcessFiles.Loft_EV_Remaining(UserID)
		if err != nil {
			return err
		}
		return nil

	case Consts_LuaHash.Magnet_BT_Elephant_Lose:
		err := ProcessFiles.Magnet_BT_Elephant_Lose(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Kubou_1:
		err := ProcessFiles.Magnet_EV_Kubou_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Kubou_3:
		err := ProcessFiles.Magnet_EV_Kubou_3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Kubou_4:
		err := ProcessFiles.Magnet_EV_Kubou_4(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Vari1_Tips:
		err := ProcessFiles.Magnet_EV_Vari1_Tips(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_2F_Dragon:
		err := ProcessFiles.Magnet_EV_Art_2F_Dragon(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_3F_Headgear:
		err := ProcessFiles.Magnet_EV_Art_3F_Headgear(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_7F_Animal:
		err := ProcessFiles.Magnet_EV_Art_7F_Animal(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_7F_Signboard:
		err := ProcessFiles.Magnet_EV_Art_7F_Signboard(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_Comp:
		err := ProcessFiles.Magnet_EV_Art_Comp(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Art_Progress:
		err := ProcessFiles.Magnet_EV_Art_Progress(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Mission_Clear:
		err := ProcessFiles.Magnet_EV_Mission_Clear(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Branch_Switch:
		err := ProcessFiles.Magnet_EV_Branch_Switch(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Pure4_1:
		err := ProcessFiles.Magnet_EV_Pure4_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Vari4_1:
		err := ProcessFiles.Magnet_EV_Vari4_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Magnet_EV_Yuusen4_1:
		err := ProcessFiles.Magnet_EV_Yuusen4_1(UserID)
		if err != nil {
			return err
		}
		return nil

	case Consts_LuaHash.Miyashita_EV_Hint_Bridge:
		err := ProcessFiles.Miyashita_EV_Hint_Bridge(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Miyashita_EV_Hint_1:
		err := ProcessFiles.Miyashita_EV_Hint_1(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Miyashita_EV_Hint_2:
		err := ProcessFiles.Miyashita_EV_Hint_2(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Miyashita_EV_Hint_3:
		err := ProcessFiles.Miyashita_EV_Hint_3(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Miyashita_EV_Unlock:
		err := ProcessFiles.Miyashita_EV_Unlock(UserID)
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
	case Consts_LuaHash.Sys_EV_Branch_Switch_EX_Off:
		err := ProcessFiles.Sys_EV_Branch_Switch_EX_Off(UserID)
		if err != nil {
			return err
		}
		return nil
	case Consts_LuaHash.Sys_EV_Branch_Switch:
		err := ProcessFiles.Sys_EV_Branch_Switch(UserID)
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
