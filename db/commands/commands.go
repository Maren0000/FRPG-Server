package db_commands

import (
	Consts_Building "FRPGServer/Consts/Building"
	Consts_Login "FRPGServer/Consts/Login"
	Consts_LuaHash "FRPGServer/Consts/LuaHash"
	Consts_MapPin "FRPGServer/Consts/MapPin"
	Consts_Quest "FRPGServer/Consts/Quest"
	Consts_QuestItem "FRPGServer/Consts/QuestItem"
	Consts_ScanTag "FRPGServer/Consts/ScanTag"
	"FRPGServer/Utils"
	"FRPGServer/db"
	"context"
	"database/sql"
	"fmt"
	"math/rand"

	_ "embed"

	_ "github.com/glebarez/go-sqlite"
)

//go:embed schema.sql
var ddl string

var ctx context.Context
var dbconn *sql.DB
var queries *db.Queries

func Opendb() error {
	var err error
	ctx = context.Background()

	dbconn, err = sql.Open("sqlite", "file:FRPG.sqlite3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened DB")

	queries = db.New(dbconn)

	if _, err := dbconn.ExecContext(ctx, ddl); err != nil {
		return err
	}

	return nil
}

func CheckDeviceIdExists(Did string) bool {
	_, err := queries.GetUser(ctx, sql.NullString{String: Did, Valid: true})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CreateNewUser(Did string) (user db.Users) {
	user, err := queries.CreateNewUser(ctx, db.CreateNewUserParams{
		DeviceID: sql.NullString{String: Did, Valid: true},
		ID:       sql.NullString{String: Utils.GenRandId(), Valid: true},
		NewGame:  sql.NullInt64{Int64: 1, Valid: true},
		UUID:     sql.NullString{String: "00000000-0000-0000-0000-000000000001", Valid: true},
		Status:   sql.NullInt64{Int64: Consts_Login.CREATING_USER, Valid: true},
	})
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func GetUser(Did string) (user db.Users, err error) {
	user, err = queries.GetUser(ctx, sql.NullString{String: Did, Valid: true})
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserID(DeviceID string) (UserID string, err error) {
	var User sql.NullString
	User, err = queries.GetUserID(ctx, sql.NullString{String: DeviceID, Valid: true})
	if err != nil {
		return UserID, err
	}
	return User.String, nil
}

func SetUserName(Did string, Name string) (err error) {
	err = queries.UpdateUserName(ctx, db.UpdateUserNameParams{
		DeviceID: sql.NullString{String: Did, Valid: true},
		Name:     sql.NullString{String: Name, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func SetUserStatus(Did string, Status int) (err error) {
	err = queries.UpdateUserStatus(ctx, db.UpdateUserStatusParams{
		DeviceID: sql.NullString{String: Did, Valid: true},
		Status:   sql.NullInt64{Int64: int64(Status), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func SetUserTeam(UserID string, TeamID string) (err error) {
	err = queries.UpdateUserTeam(ctx, db.UpdateUserTeamParams{
		ID:     sql.NullString{String: UserID, Valid: true},
		TeamID: sql.NullString{String: TeamID, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func SetUserNewGame(Did string, Value int) (err error) {
	err = queries.UpdateUserNewGame(ctx, db.UpdateUserNewGameParams{
		DeviceID: sql.NullString{String: Did, Valid: true},
		NewGame:  sql.NullInt64{Int64: int64(Value), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateTeam(TeamName string) (team db.Teams, err error) {
	team, err = queries.CreateNewTeam(ctx, db.CreateNewTeamParams{
		TeamID:   sql.NullString{String: Utils.GenRandId(), Valid: true},
		RoomID:   sql.NullString{String: Utils.GenRandId(), Valid: true},
		TeamName: sql.NullString{String: TeamName, Valid: true},
	})
	if err != nil {
		return team, err
	}
	return team, nil
}

func InitSaveData(UserID string) (err error) {
	//To-Do: Move ColorID Random to Utils
	_, err = queries.CreateNewUserSave(ctx, db.CreateNewUserSaveParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		BIntro:    sql.NullInt64{Int64: 1, Valid: true},
		NowHP:     sql.NullInt64{Int64: 5, Valid: true},
		MaxHP:     sql.NullInt64{Int64: 5, Valid: true},
		ColorID:   sql.NullInt64{Int64: rand.Int63n(3) + 1, Valid: true},
		BNewQuest: sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		Name:      sql.NullString{String: "MaruiPin", Valid: true},
		PinType:   sql.NullString{String: Consts_MapPin.MapPin, Valid: true},
		Latitude:  sql.NullFloat64{Float64: 35.660866, Valid: true},
		Longitude: sql.NullFloat64{Float64: 139.701024, Valid: true},
		QuestID:   sql.NullInt64{Int64: 1, Valid: true},
		IsRemove:  sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}

	//To-Do: add the rest of the OnceEvents since they are all added at the start
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Introduction, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Oioi_5F_EV_FirstBattle, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Sys_EV_Burger_FirstDamage, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Modi_EV_Kariya_1, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Modi_EV_FirstBattle, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Loft_EV_Uduki_1, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Magnet_EV_Kubou_1, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserBuilding(ctx, db.CreateNewUserBuildingParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		Prefab: sql.NullString{String: Consts_Building.Marui, Valid: true},
		Name:   sql.NullString{String: "マルイ", Valid: true},
		Status: sql.NullString{String: "2", Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserQuest(ctx, db.CreateNewUserQuestParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		ID:        sql.NullInt64{Int64: Consts_Quest.Quest_1_Start, Valid: true},
		Value:     sql.NullInt64{Int64: 0, Valid: true},
		IsCurrent: sql.NullInt64{Int64: 1, Valid: true},
		IsClear:   sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserQuestItem(ctx, db.CreateNewUserQuestItemParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		QuestID: sql.NullInt64{Int64: Consts_Quest.Quest_1_Start, Valid: true},
		Name:    sql.NullString{String: Consts_QuestItem.Quest_1_JoinBadge, Valid: true},
		Type:    sql.NullString{String: "off", Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserScan(ctx, db.CreateNewUserScanParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		Type:     sql.NullInt64{Int64: 3, Valid: true},
		Tag:      sql.NullString{String: Consts_ScanTag.QR_Q1_JoinBadge, Valid: true},
		BMulti:   sql.NullInt64{Int64: 0, Valid: true},
		LuaHash:  sql.NullInt64{Int64: Consts_LuaHash.Oioi_8F_EV_Deai_0, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserResume(ctx, db.CreateNewUserResumeParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		BResume:  sql.NullInt64{Int64: 0, Valid: true},
		LuaHash:  sql.NullInt64{Int64: Consts_LuaHash.Introduction, Valid: true},
		TagID:    sql.NullInt64{Int64: 0, Valid: true},
		ResumeID: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserLocalMap(ctx, db.CreateNewUserLocalMapParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		Name:   sql.NullString{Valid: false},
		Floor:  sql.NullInt64{Valid: false},
	})
	if err != nil {
		return err
	}

	return nil
}

func GetUserSavaData(UserID string) (SaveData db.UserSave, err error) {
	SaveData, err = queries.GetUserSave(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return SaveData, err
	}
	return SaveData, nil
}

func GetUserSaveColor(UserID string) (ColorID int, err error) {
	var UserSaveColor sql.NullInt64
	UserSaveColor, err = queries.GetUserSaveColor(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return ColorID, err
	}
	return int(UserSaveColor.Int64), nil
}

func UpdateUserSaveIntro(UserID string, Value int) (err error) {
	err = queries.UpdateUserSaveIntro(ctx, db.UpdateUserSaveIntroParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		BIntro: sql.NullInt64{Int64: int64(Value), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveNewQuest(UserID string, Bool int) (err error) {
	err = queries.UpdateUserSaveNewQuest(ctx, db.UpdateUserSaveNewQuestParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		BNewQuest: sql.NullInt64{Int64: int64(Bool), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveNowHP(UserID string, NowHp int) (err error) {
	err = queries.UpdateUserSaveNowHP(ctx, db.UpdateUserSaveNowHPParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		NowHP:  sql.NullInt64{Int64: int64(NowHp), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveBattleId(UserID string, BattleId int) (err error) {
	err = queries.UpdateUserSaveBattleID(ctx, db.UpdateUserSaveBattleIDParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		BattleID: sql.NullInt64{Int64: int64(BattleId), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveBattleBadge1(UserID string, BattleBadge int) (err error) {
	err = queries.UpdateUserSaveBattleBadge1(ctx, db.UpdateUserSaveBattleBadge1Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge1: sql.NullInt64{Int64: int64(BattleBadge), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveBattleBadge2(UserID string, BattleBadge int) (err error) {
	err = queries.UpdateUserSaveBattleBadge2(ctx, db.UpdateUserSaveBattleBadge2Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge2: sql.NullInt64{Int64: int64(BattleBadge), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserSaveBattleBadge3(UserID string, BattleBadge int) (err error) {
	err = queries.UpdateUserSaveBattleBadge3(ctx, db.UpdateUserSaveBattleBadge3Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge3: sql.NullInt64{Int64: int64(BattleBadge), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func EmptyUserSaveBadges(UserID string) (err error) {
	err = queries.UpdateUserSaveBattleBadge1(ctx, db.UpdateUserSaveBattleBadge1Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge1: sql.NullInt64{Valid: false},
	})
	if err != nil {
		return err
	}
	err = queries.UpdateUserSaveBattleBadge2(ctx, db.UpdateUserSaveBattleBadge2Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge2: sql.NullInt64{Valid: false},
	})
	if err != nil {
		return err
	}
	err = queries.UpdateUserSaveBattleBadge3(ctx, db.UpdateUserSaveBattleBadge3Params{
		UserID:       sql.NullString{String: UserID, Valid: true},
		BattleBadge3: sql.NullInt64{Valid: false},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserQuest(UserID string, QuestID int) (UserQuest db.UserQuest, err error) {
	UserQuest, err = queries.GetUserQuest(ctx, db.GetUserQuestParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		ID:     sql.NullInt64{Int64: int64(QuestID), Valid: true},
	})
	if err != nil {
		return UserQuest, err
	}
	return UserQuest, nil
}

func CheckUserQuestExists(UserID string, QuestID int) (Exists bool, err error) {
	var UserQuests []db.UserQuest
	UserQuests, err = queries.GetUserQuestAll(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return false, err
	}
	for _, quest := range UserQuests {
		if quest.ID.Int64 == int64(QuestID) {
			return true, nil
		}
	}
	return false, nil
}

func GetCurrentUserQuest(UserID string) (UserQuest db.UserQuest, err error) {
	UserQuest, err = queries.GetUserQuestCurrent(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserQuest, err
	}
	return UserQuest, nil
}

func CheckCurrentUserQuest(UserID string, QuestID int) (check bool, err error) {
	var UserQuest db.UserQuest
	UserQuest, err = queries.GetUserQuestCurrent(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return false, err
	}
	if UserQuest.ID.Int64 == int64(QuestID) {
		return true, nil
	}
	return false, nil
}

func CreateUserQuestItem(UserID string, QuestID int, ItemID string, Type string) (err error) {
	_, err = queries.CreateNewUserQuestItem(ctx, db.CreateNewUserQuestItemParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		QuestID: sql.NullInt64{Int64: int64(QuestID), Valid: true},
		Name:    sql.NullString{String: ItemID, Valid: true},
		Type:    sql.NullString{String: Type, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserQuestClear(UserID string, QuestID int, IsClear int) (err error) {
	err = queries.UpdateUserQuestClear(ctx, db.UpdateUserQuestClearParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		ID:      sql.NullInt64{Int64: int64(QuestID), Valid: true},
		IsClear: sql.NullInt64{Int64: int64(IsClear), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func CheckUserQuestCleared(UserID string, QuestID int) (Exists bool, err error) {
	var UserQuests []db.UserQuest
	UserQuests, err = queries.GetUserQuestAll(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return false, err
	}
	for _, quest := range UserQuests {
		if quest.ID.Int64 == int64(QuestID) && quest.IsClear.Int64 == 1 {
			return true, nil
		}
	}
	return false, nil
}

func UpdateUserQuestCurrent(UserID string, QuestID int, IsCurrent int) (err error) {
	err = queries.UpdateUserQuestCurrent(ctx, db.UpdateUserQuestCurrentParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		ID:        sql.NullInt64{Int64: int64(QuestID), Valid: true},
		IsCurrent: sql.NullInt64{Int64: int64(IsCurrent), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserQuestProgress(UserID string, QuestID int, Increment int) (err error) {
	var UserQuest db.UserQuest
	UserQuest, err = queries.GetUserQuest(ctx, db.GetUserQuestParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		ID:     sql.NullInt64{Int64: int64(QuestID), Valid: true},
	})
	if err != nil {
		return err
	}

	err = queries.UpdateUserQuestProgress(ctx, db.UpdateUserQuestProgressParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		ID:     sql.NullInt64{Int64: int64(QuestID), Valid: true},
		Value:  sql.NullInt64{Int64: UserQuest.Value.Int64 + int64(Increment), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateUserGPSPin(UserID string, Name string, PinType string, PinColor string, Latitude float64, Longitude float64, QuestId int, MapType string, MapFloor string) (err error) {
	//We need to check if the entry already exists or not.
	//If it does then remove IsRemove, Otherwise make entry
	GPSPins, err := queries.ListUserGPSRemoved(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return err
	}

	if len(GPSPins) > 0 {
		for _, Pin := range GPSPins {
			if Pin.Name.String == Name {
				err = queries.UpdateUserGPSRemove(ctx, db.UpdateUserGPSRemoveParams{
					UserID:   sql.NullString{String: UserID, Valid: true},
					Name:     sql.NullString{String: Name, Valid: true},
					IsRemove: sql.NullInt64{Int64: 0, Valid: true},
				})
				if err != nil {
					return err
				}
				return nil
			} else {
				continue
			}
		}
		_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
			UserID:    sql.NullString{String: UserID, Valid: true},
			Name:      sql.NullString{String: Name, Valid: true},
			PinType:   sql.NullString{String: PinType, Valid: true},
			PinColor:  sql.NullString{String: PinColor, Valid: true},
			Latitude:  sql.NullFloat64{Float64: Latitude, Valid: true},
			Longitude: sql.NullFloat64{Float64: Longitude, Valid: true},
			QuestID:   sql.NullInt64{Int64: int64(QuestId), Valid: true},
			MapType:   sql.NullString{String: MapType, Valid: true},
			MapNo:     sql.NullString{String: MapFloor, Valid: true},
			IsRemove:  sql.NullInt64{Int64: 0, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
			UserID:    sql.NullString{String: UserID, Valid: true},
			Name:      sql.NullString{String: Name, Valid: true},
			PinType:   sql.NullString{String: PinType, Valid: true},
			PinColor:  sql.NullString{String: PinColor, Valid: true},
			Latitude:  sql.NullFloat64{Float64: Latitude, Valid: true},
			Longitude: sql.NullFloat64{Float64: Longitude, Valid: true},
			QuestID:   sql.NullInt64{Int64: int64(QuestId), Valid: true},
			MapType:   sql.NullString{String: MapType, Valid: true},
			MapNo:     sql.NullString{String: MapFloor, Valid: true},
			IsRemove:  sql.NullInt64{Int64: 0, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func CreateUserGPSEventPin(UserID string, Name string, Latitude float64, Longitude float64, LuaHash uint32) (err error) {
	_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
		UserID:         sql.NullString{String: UserID, Valid: true},
		Name:           sql.NullString{String: Name, Valid: true},
		PinType:        sql.NullString{String: "", Valid: true},
		PinColor:       sql.NullString{String: "", Valid: true},
		Latitude:       sql.NullFloat64{Float64: Latitude, Valid: true},
		Longitude:      sql.NullFloat64{Float64: Longitude, Valid: true},
		QuestID:        sql.NullInt64{Int64: Consts_Quest.Quest_59651, Valid: true},
		BLocationEvent: sql.NullInt64{Int64: 1, Valid: true},
		LuaScript:      sql.NullInt64{Int64: int64(LuaHash), Valid: true},
		MapType:        sql.NullString{String: "", Valid: true},
		MapNo:          sql.NullString{String: "", Valid: true},
		IsRemove:       sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserGPSQuest(UserID string, name string, QuestID int) (err error) {
	err = queries.UpdateUserGPSQuest(ctx, db.UpdateUserGPSQuestParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		Name:    sql.NullString{String: name, Valid: true},
		QuestID: sql.NullInt64{Int64: int64(QuestID), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func SetUserLocalMap(UserID string, Buildingname string, Floor int) (err error) {
	err = queries.UpdateUserLocalMap(ctx, db.UpdateUserLocalMapParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		Name:   sql.NullString{String: Buildingname, Valid: true},
		Floor:  sql.NullInt64{Int64: int64(Floor), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateUserQuest(UserID string, QuestID int, IsCurrent int) (err error) {
	_, err = queries.CreateNewUserQuest(ctx, db.CreateNewUserQuestParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		ID:        sql.NullInt64{Int64: int64(QuestID), Valid: true},
		Value:     sql.NullInt64{Int64: 0, Valid: true},
		IsCurrent: sql.NullInt64{Int64: 1, Valid: true},
		IsClear:   sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}
	err = UpdateUserSaveNewQuest(UserID, 1)
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentUserLocalMap(UserID string) (UserLocalMap db.UserLocalMap, err error) {
	UserLocalMap, err = queries.GetUserLocalMap(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserLocalMap, err
	}
	return UserLocalMap, nil
}

func ListUserGPS(UserID string) (UserGPS []db.UserGPS, err error) {
	UserGPS, err = queries.ListUserGPS(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserGPS, err
	}
	return UserGPS, nil
}

func UpdateUserGPSRemove(UserID string, GPSName string, Remove int) (err error) {
	err = queries.UpdateUserGPSRemove(ctx, db.UpdateUserGPSRemoveParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		Name:     sql.NullString{String: GPSName, Valid: true},
		IsRemove: sql.NullInt64{Int64: int64(Remove), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func ListUserGPSRemoved(UserID string) (UserGPS []db.UserGPS, err error) {
	UserGPS, err = queries.ListUserGPSRemoved(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserGPS, err
	}
	return UserGPS, nil
}

func ListUserQuestItems(UserID string, QuestID int) (UserQuestItems []db.UserQuestItems, err error) {
	UserQuestItems, err = queries.ListUserQuestItem(ctx, db.ListUserQuestItemParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		QuestID: sql.NullInt64{Int64: int64(QuestID), Valid: true},
	})
	if err != nil {
		return UserQuestItems, err
	}
	return UserQuestItems, nil
}

func UpdateUserQuestItem(UserID string, QuestID int, Name string, Type string) (err error) {
	err = queries.UpdateUserQuestItem(ctx, db.UpdateUserQuestItemParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		QuestID: sql.NullInt64{Int64: int64(QuestID), Valid: true},
		Name:    sql.NullString{String: Name, Valid: true},
		Type:    sql.NullString{String: Type, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func CreateUserScan(UserID string, Type int, Tag string, Bmulti int, LuaHash uint32) (err error) {
	//We need to check if the entry already exists in Removed or not.
	//If it does then remove IsRemove, Otherwise make entry
	Scans, err := queries.ListUserScanRemoved(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return err
	}
	if len(Scans) > 0 {
		for _, Scan := range Scans {
			if Scan.Tag.String == Tag {
				err = queries.UpdateUserScanRemove(ctx, db.UpdateUserScanRemoveParams{
					UserID:   sql.NullString{String: UserID, Valid: true},
					Tag:      sql.NullString{String: Tag, Valid: true},
					IsRemove: sql.NullInt64{Int64: 0, Valid: true},
				})
				if err != nil {
					return err
				}
				return nil
			} else {
				continue
			}
		}
		_, err = queries.CreateNewUserScan(ctx, db.CreateNewUserScanParams{
			UserID:   sql.NullString{String: UserID, Valid: true},
			Type:     sql.NullInt64{Int64: int64(Type), Valid: true},
			Tag:      sql.NullString{String: Tag, Valid: true},
			BMulti:   sql.NullInt64{Int64: int64(Bmulti), Valid: true},
			LuaHash:  sql.NullInt64{Int64: int64(LuaHash), Valid: true},
			IsRemove: sql.NullInt64{Int64: 0, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	} else {
		_, err = queries.CreateNewUserScan(ctx, db.CreateNewUserScanParams{
			UserID:   sql.NullString{String: UserID, Valid: true},
			Type:     sql.NullInt64{Int64: int64(Type), Valid: true},
			Tag:      sql.NullString{String: Tag, Valid: true},
			BMulti:   sql.NullInt64{Int64: int64(Bmulti), Valid: true},
			LuaHash:  sql.NullInt64{Int64: int64(LuaHash), Valid: true},
			IsRemove: sql.NullInt64{Int64: 0, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	}
}

func ListUserScan(UserID string) (UserScans []db.UserScans, err error) {
	UserScans, err = queries.ListUserScan(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserScans, err
	}
	return UserScans, nil
}

func UpdateUserScanLuaHash(UserID string, ScanTag string, LuaHash uint32) (err error) {
	err = queries.UpdateUserScanLuaHash(ctx, db.UpdateUserScanLuaHashParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		Tag:     sql.NullString{String: ScanTag, Valid: true},
		LuaHash: sql.NullInt64{Int64: int64(LuaHash), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserScanRemove(UserID string, ScanTag string, Remove int) (err error) {
	err = queries.UpdateUserScanRemove(ctx, db.UpdateUserScanRemoveParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		Tag:      sql.NullString{String: ScanTag, Valid: true},
		IsRemove: sql.NullInt64{Int64: int64(Remove), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserScan(UserID string, ScanId int) (UserScan db.UserScans, err error) {
	UserScan, err = queries.GetUserScan(ctx, db.GetUserScanParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		ID:     int64(ScanId),
	})
	if err != nil {
		return UserScan, err
	}
	return UserScan, nil
}

func ListUserScanRemoved(UserID string) (UserScans []db.UserScans, err error) {
	UserScans, err = queries.ListUserScanRemoved(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserScans, err
	}
	return UserScans, nil
}

func ListUserBuilding(UserID string) (UserBuildings []db.UserBuildings, err error) {
	UserBuildings, err = queries.ListUserBuildings(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserBuildings, err
	}
	return UserBuildings, nil
}

func CreateUserBuilding(UserID string, Prefab string, Name string, Status string) (err error) {
	_, err = queries.CreateNewUserBuilding(ctx, db.CreateNewUserBuildingParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		Prefab: sql.NullString{String: Prefab, Valid: true},
		Name:   sql.NullString{String: Name, Valid: true},
		Status: sql.NullString{String: Status, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserBuildingStatus(UserID string, Prefab string, Status string) (err error) {
	err = queries.UpdateUserBuildingStatus(ctx, db.UpdateUserBuildingStatusParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		Prefab: sql.NullString{String: Prefab, Valid: true},
		Status: sql.NullString{String: Status, Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func ListUserOnceEvent(UserID string) (UserOnceEvent []db.UserOnceEvent, err error) {
	UserOnceEvent, err = queries.ListUserOnceEvent(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserOnceEvent, err
	}
	return UserOnceEvent, nil
}

func ListUserOnceEventRemoved(UserID string) (UserOnceEvent []db.UserOnceEvent, err error) {
	UserOnceEvent, err = queries.ListUserOnceEventRemoved(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserOnceEvent, err
	}
	return UserOnceEvent, nil
}

func CreateUserCharacter(UserID string, CharacterId int, ItemId int) (Character db.UserCharacters, err error) {
	Character, err = queries.CreateNewUserCharacter(ctx, db.CreateNewUserCharacterParams{
		UserID:      sql.NullString{String: UserID, Valid: true},
		CharacterId: sql.NullInt64{Int64: int64(CharacterId), Valid: true},
		ItemId:      sql.NullInt64{Int64: int64(ItemId), Valid: true},
	})
	if err != nil {
		return Character, err
	}
	return Character, nil
}

func ListUserCharacters(UserID string) (UserCharacters []db.UserCharacters, err error) {
	UserCharacters, err = queries.ListUserCharacters(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserCharacters, err
	}
	return UserCharacters, nil
}

func CreateUserItem(UserID string, ItemId int) (Item db.UserItems, err error) {
	//Check if user already has item
	UserItems, err := queries.ListUserItems(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return Item, err
	}

	if len(UserItems) > 0 {
		for _, item := range UserItems {
			if item.Item.Int64 == int64(ItemId) {
				return item, nil
			} else {
				continue
			}
		}
		Item, err = queries.CreateNewUserItem(ctx, db.CreateNewUserItemParams{
			UserID: sql.NullString{String: UserID, Valid: true},
			Item:   sql.NullInt64{Int64: int64(ItemId), Valid: true},
		})
		if err != nil {
			return Item, err
		}
	} else {
		Item, err = queries.CreateNewUserItem(ctx, db.CreateNewUserItemParams{
			UserID: sql.NullString{String: UserID, Valid: true},
			Item:   sql.NullInt64{Int64: int64(ItemId), Valid: true},
		})
		if err != nil {
			return Item, err
		}
		return Item, nil
	}
	return Item, nil
}

func ListUserItems(UserID string) (UserItems []db.UserItems, err error) {
	UserItems, err = queries.ListUserItems(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserItems, err
	}
	return UserItems, nil
}

func CreateUserOnceEvent(UserID string, hash uint32, Remove int) (err error) {
	_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: int64(hash), Valid: true},
		IsRemove: sql.NullInt64{Int64: int64(Remove), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserOnceEvent(UserID string, hash uint32, Remove int) (err error) {
	err = queries.UpdateUserOnceEventRemove(ctx, db.UpdateUserOnceEventRemoveParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: int64(hash), Valid: true},
		IsRemove: sql.NullInt64{Int64: int64(Remove), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserResume(UserID string) (ResumeData db.UserResume, err error) {
	ResumeData, err = queries.GetUserResume(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return ResumeData, err
	}
	return ResumeData, nil
}

func UpdateUserResume(UserID string, Bool int, LuaHash uint32, TagId int) (err error) {
	err = queries.UpdateUserResume(ctx, db.UpdateUserResumeParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		BResume: sql.NullInt64{Int64: int64(Bool), Valid: true},
		LuaHash: sql.NullInt64{Int64: int64(LuaHash), Valid: true},
		TagID:   sql.NullInt64{Int64: int64(TagId), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserResumeLuaResume(UserID string, Resume int) (err error) {
	err = queries.UpdateUserResumeLuaResumeID(ctx, db.UpdateUserResumeLuaResumeIDParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		ResumeID: sql.NullInt64{Int64: int64(Resume), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserResumeBool(UserID string, Bool int) (err error) {
	err = queries.UpdateUserResumeBool(ctx, db.UpdateUserResumeBoolParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		BResume: sql.NullInt64{Int64: int64(Bool), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserCharacterItem(UserID string, CharacterID int, ItemID int) (err error) {
	err = queries.UpdateUserCharacterItem(ctx, db.UpdateUserCharacterItemParams{
		UserID:      sql.NullString{String: UserID, Valid: true},
		CharacterId: sql.NullInt64{Int64: int64(CharacterID), Valid: true},
		ItemId:      sql.NullInt64{Int64: int64(ItemID), Valid: true},
	})
	if err != nil {
		return err
	}
	return nil
}
