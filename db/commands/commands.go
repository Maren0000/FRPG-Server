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
	_, err = queries.CreateNewUserSave(ctx, db.CreateNewUserSaveParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		BIntro:    sql.NullInt64{Int64: 1, Valid: true},
		NowHP:     sql.NullInt64{Int64: 100, Valid: true},
		MaxHP:     sql.NullInt64{Int64: 100, Valid: true},
		ColorID:   sql.NullInt64{Int64: rand.Int63n(3-1) + 1, Valid: true},
		BNewQuest: sql.NullInt64{Int64: 1, Valid: true},
	})
	if err != nil {
		return err
	}

	_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		ID:        sql.NullString{String: "1", Valid: true},
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

	/*_, err = queries.CreateNewUserOnceEvent(ctx, db.CreateNewUserOnceEventParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		UInt:     sql.NullInt64{Int64: Consts_LuaHash.Introduction, Valid: true},
		IsRemove: sql.NullInt64{Int64: 0, Valid: true},
	})
	if err != nil {
		return err
	}*/

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
		UserID:  sql.NullString{String: UserID, Valid: true},
		ID:      sql.NullInt64{Int64: Consts_Quest.Quest_1_Start, Valid: true},
		Value:   sql.NullInt64{Int64: 0, Valid: true},
		IsClear: sql.NullInt64{Int64: 0, Valid: true},
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
		ID:       sql.NullInt64{Int64: 1, Valid: true},
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

func GetUserSavaColor(UserID string) (ColorID int, err error) {
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

func GetCurrentUserQuest(UserID string) (UserQuest db.UserQuest, err error) {
	UserQuest, err = queries.GetUserQuestCurrent(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserQuest, err
	}
	return UserQuest, nil
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

func UpdateUserQuestProgress(UserID string, QuestID int, Increment int) (err error) {
	var UserQuest db.UserQuest
	UserQuest, err = queries.GetUserQuestCurrent(ctx, sql.NullString{String: UserID, Valid: true})
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

func CreateUserGPSPin(UserID string, ID string, Name string, PinType string, PinColor string, Latitude float64, Longitude float64, QuestId int, MapType string, MapFloor string) (err error) {
	_, err = queries.CreateNewUserGPS(ctx, db.CreateNewUserGPSParams{
		UserID:    sql.NullString{String: UserID, Valid: true},
		ID:        sql.NullString{String: ID, Valid: true},
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

func CreateUserQuest(UserID string, QuestID int) (err error) {
	_, err = queries.CreateNewUserQuest(ctx, db.CreateNewUserQuestParams{
		UserID:  sql.NullString{String: UserID, Valid: true},
		ID:      sql.NullInt64{Int64: int64(QuestID), Valid: true},
		Value:   sql.NullInt64{Int64: 0, Valid: true},
		IsClear: sql.NullInt64{Int64: 0, Valid: true},
	})
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

func UpdateUserGPSRemove(UserID string, GPSId string, Remove int) (err error) {
	err = queries.UpdateUserGPSRemove(ctx, db.UpdateUserGPSRemoveParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		ID:       sql.NullString{String: GPSId, Valid: true},
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

func CreateUserScan(UserID string, ID int, Type int, Tag string, Bmulti int, LuaHash uint32) (err error) {
	_, err = queries.CreateNewUserScan(ctx, db.CreateNewUserScanParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		ID:       sql.NullInt64{Int64: int64(ID), Valid: true},
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

func GetUserScan(UserID string, ScanId int) (UserScan db.UserScans, err error) {
	UserScan, err = queries.GetUserScan(ctx, db.GetUserScanParams{
		UserID: sql.NullString{String: UserID, Valid: true},
		ID:     sql.NullInt64{Int64: int64(ScanId), Valid: true},
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

func ListUserCharacters(UserID string) (UserCharacters []db.UserCharacters, err error) {
	UserCharacters, err = queries.ListUserCharacters(ctx, sql.NullString{String: UserID, Valid: true})
	if err != nil {
		return UserCharacters, err
	}
	return UserCharacters, nil
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

func UpdateUserResume(UserID string, Bool int, LuaHash uint32, TagId int, ResumeId int) (err error) {
	err = queries.UpdateUserResume(ctx, db.UpdateUserResumeParams{
		UserID:   sql.NullString{String: UserID, Valid: true},
		BResume:  sql.NullInt64{Int64: int64(Bool), Valid: true},
		LuaHash:  sql.NullInt64{Int64: int64(LuaHash), Valid: true},
		TagID:    sql.NullInt64{Int64: int64(TagId), Valid: true},
		ResumeID: sql.NullInt64{Int64: int64(ResumeId), Valid: true},
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
