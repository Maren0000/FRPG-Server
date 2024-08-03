// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createNewTeam = `-- name: CreateNewTeam :one
INSERT INTO "teams" (
  "TeamID", "RoomID", "TeamName"
) VALUES (
  ?, ?, ?
)
RETURNING TeamID, RoomID, TeamName
`

type CreateNewTeamParams struct {
	TeamID   sql.NullString
	RoomID   sql.NullString
	TeamName sql.NullString
}

func (q *Queries) CreateNewTeam(ctx context.Context, arg CreateNewTeamParams) (Teams, error) {
	row := q.db.QueryRowContext(ctx, createNewTeam, arg.TeamID, arg.RoomID, arg.TeamName)
	var i Teams
	err := row.Scan(&i.TeamID, &i.RoomID, &i.TeamName)
	return i, err
}

const createNewUser = `-- name: CreateNewUser :one
INSERT INTO "users" (
  "DeviceID", "ID", "NewGame", "UUID", "Status"
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING DeviceID, ID, Name, Status, TeamID, NewGame, UUID
`

type CreateNewUserParams struct {
	DeviceID sql.NullString
	ID       sql.NullString
	NewGame  sql.NullInt64
	UUID     sql.NullString
	Status   sql.NullInt64
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createNewUser,
		arg.DeviceID,
		arg.ID,
		arg.NewGame,
		arg.UUID,
		arg.Status,
	)
	var i Users
	err := row.Scan(
		&i.DeviceID,
		&i.ID,
		&i.Name,
		&i.Status,
		&i.TeamID,
		&i.NewGame,
		&i.UUID,
	)
	return i, err
}

const createNewUserBuilding = `-- name: CreateNewUserBuilding :one
INSERT INTO "userBuildings" (
  "UserID", "Prefab", "Name", "Status"
) VALUES (
  ?, ?, ?, ?
)
RETURNING UserID, Prefab, Name, Status
`

type CreateNewUserBuildingParams struct {
	UserID sql.NullString
	Prefab sql.NullString
	Name   sql.NullString
	Status sql.NullString
}

func (q *Queries) CreateNewUserBuilding(ctx context.Context, arg CreateNewUserBuildingParams) (UserBuildings, error) {
	row := q.db.QueryRowContext(ctx, createNewUserBuilding,
		arg.UserID,
		arg.Prefab,
		arg.Name,
		arg.Status,
	)
	var i UserBuildings
	err := row.Scan(
		&i.UserID,
		&i.Prefab,
		&i.Name,
		&i.Status,
	)
	return i, err
}

const createNewUserCharacter = `-- name: CreateNewUserCharacter :one
INSERT INTO "userCharacters" (
  "UserID", "CharacterId", "ItemId"
) VALUES (
  ?, ?, ?
)
RETURNING UserID, CharacterId, ItemId
`

type CreateNewUserCharacterParams struct {
	UserID      sql.NullString
	CharacterId sql.NullInt64
	ItemId      sql.NullInt64
}

func (q *Queries) CreateNewUserCharacter(ctx context.Context, arg CreateNewUserCharacterParams) (UserCharacters, error) {
	row := q.db.QueryRowContext(ctx, createNewUserCharacter, arg.UserID, arg.CharacterId, arg.ItemId)
	var i UserCharacters
	err := row.Scan(&i.UserID, &i.CharacterId, &i.ItemId)
	return i, err
}

const createNewUserGPS = `-- name: CreateNewUserGPS :one
INSERT INTO "userGPS" (
  "UserID", "ID", "Name", "PinType", "PinColor", "Latitude", "Longitude", "LuaScript", "BLocationEvent", "QuestID", "MapType", "MapNo", "IsRemove"
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING UserID, ID, Name, PinType, PinColor, Latitude, Longitude, LuaScript, BLocationEvent, QuestID, MapType, MapNo, IsRemove
`

type CreateNewUserGPSParams struct {
	UserID         sql.NullString
	ID             sql.NullString
	Name           sql.NullString
	PinType        sql.NullString
	PinColor       sql.NullString
	Latitude       sql.NullFloat64
	Longitude      sql.NullFloat64
	LuaScript      sql.NullInt64
	BLocationEvent sql.NullInt64
	QuestID        sql.NullInt64
	MapType        sql.NullString
	MapNo          sql.NullString
	IsRemove       sql.NullInt64
}

func (q *Queries) CreateNewUserGPS(ctx context.Context, arg CreateNewUserGPSParams) (UserGPS, error) {
	row := q.db.QueryRowContext(ctx, createNewUserGPS,
		arg.UserID,
		arg.ID,
		arg.Name,
		arg.PinType,
		arg.PinColor,
		arg.Latitude,
		arg.Longitude,
		arg.LuaScript,
		arg.BLocationEvent,
		arg.QuestID,
		arg.MapType,
		arg.MapNo,
		arg.IsRemove,
	)
	var i UserGPS
	err := row.Scan(
		&i.UserID,
		&i.ID,
		&i.Name,
		&i.PinType,
		&i.PinColor,
		&i.Latitude,
		&i.Longitude,
		&i.LuaScript,
		&i.BLocationEvent,
		&i.QuestID,
		&i.MapType,
		&i.MapNo,
		&i.IsRemove,
	)
	return i, err
}

const createNewUserItem = `-- name: CreateNewUserItem :one
INSERT INTO "userItems" (
  "UserID", "Item"
) VALUES (
  ?, ?
)
RETURNING UserID, Item
`

type CreateNewUserItemParams struct {
	UserID sql.NullString
	Item   sql.NullInt64
}

func (q *Queries) CreateNewUserItem(ctx context.Context, arg CreateNewUserItemParams) (UserItems, error) {
	row := q.db.QueryRowContext(ctx, createNewUserItem, arg.UserID, arg.Item)
	var i UserItems
	err := row.Scan(&i.UserID, &i.Item)
	return i, err
}

const createNewUserLocalMap = `-- name: CreateNewUserLocalMap :one
INSERT INTO "userLocalMap" (
  "UserID", "Name", "Floor"
) VALUES (
  ?, ?, ?
)
RETURNING UserID, Name, Floor
`

type CreateNewUserLocalMapParams struct {
	UserID sql.NullString
	Name   sql.NullString
	Floor  sql.NullInt64
}

func (q *Queries) CreateNewUserLocalMap(ctx context.Context, arg CreateNewUserLocalMapParams) (UserLocalMap, error) {
	row := q.db.QueryRowContext(ctx, createNewUserLocalMap, arg.UserID, arg.Name, arg.Floor)
	var i UserLocalMap
	err := row.Scan(&i.UserID, &i.Name, &i.Floor)
	return i, err
}

const createNewUserOnceEvent = `-- name: CreateNewUserOnceEvent :one
INSERT INTO "userOnceEvent" (
  "UserID", "UInt", "IsRemove"
) VALUES (
  ?, ?, ?
)
RETURNING UserID, UInt, IsRemove
`

type CreateNewUserOnceEventParams struct {
	UserID   sql.NullString
	UInt     sql.NullInt64
	IsRemove sql.NullInt64
}

func (q *Queries) CreateNewUserOnceEvent(ctx context.Context, arg CreateNewUserOnceEventParams) (UserOnceEvent, error) {
	row := q.db.QueryRowContext(ctx, createNewUserOnceEvent, arg.UserID, arg.UInt, arg.IsRemove)
	var i UserOnceEvent
	err := row.Scan(&i.UserID, &i.UInt, &i.IsRemove)
	return i, err
}

const createNewUserQuest = `-- name: CreateNewUserQuest :one
INSERT INTO "userQuest" (
  "UserID", "ID", "Value", "IsClear"
) VALUES (
  ?, ?, ?, ?
)
RETURNING UserID, ID, Value, IsClear
`

type CreateNewUserQuestParams struct {
	UserID  sql.NullString
	ID      sql.NullInt64
	Value   sql.NullInt64
	IsClear sql.NullInt64
}

func (q *Queries) CreateNewUserQuest(ctx context.Context, arg CreateNewUserQuestParams) (UserQuest, error) {
	row := q.db.QueryRowContext(ctx, createNewUserQuest,
		arg.UserID,
		arg.ID,
		arg.Value,
		arg.IsClear,
	)
	var i UserQuest
	err := row.Scan(
		&i.UserID,
		&i.ID,
		&i.Value,
		&i.IsClear,
	)
	return i, err
}

const createNewUserQuestItem = `-- name: CreateNewUserQuestItem :one
INSERT INTO "userQuestItems" (
  "UserID", "QuestID", "Name", "Type"
) VALUES (
  ?, ?, ?, ?
)
RETURNING UserID, QuestID, Name, Type
`

type CreateNewUserQuestItemParams struct {
	UserID  sql.NullString
	QuestID sql.NullInt64
	Name    sql.NullString
	Type    sql.NullString
}

func (q *Queries) CreateNewUserQuestItem(ctx context.Context, arg CreateNewUserQuestItemParams) (UserQuestItems, error) {
	row := q.db.QueryRowContext(ctx, createNewUserQuestItem,
		arg.UserID,
		arg.QuestID,
		arg.Name,
		arg.Type,
	)
	var i UserQuestItems
	err := row.Scan(
		&i.UserID,
		&i.QuestID,
		&i.Name,
		&i.Type,
	)
	return i, err
}

const createNewUserResume = `-- name: CreateNewUserResume :one
INSERT INTO "userResume" (
  "UserID", "BResume", "LuaHash", "TagID", "ResumeID"
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING UserID, BResume, LuaHash, TagID, ResumeID
`

type CreateNewUserResumeParams struct {
	UserID   sql.NullString
	BResume  sql.NullInt64
	LuaHash  sql.NullInt64
	TagID    sql.NullInt64
	ResumeID sql.NullInt64
}

func (q *Queries) CreateNewUserResume(ctx context.Context, arg CreateNewUserResumeParams) (UserResume, error) {
	row := q.db.QueryRowContext(ctx, createNewUserResume,
		arg.UserID,
		arg.BResume,
		arg.LuaHash,
		arg.TagID,
		arg.ResumeID,
	)
	var i UserResume
	err := row.Scan(
		&i.UserID,
		&i.BResume,
		&i.LuaHash,
		&i.TagID,
		&i.ResumeID,
	)
	return i, err
}

const createNewUserSave = `-- name: CreateNewUserSave :one
INSERT INTO "userSave" (
  "UserID", "BIntro", "NowHP", "MaxHP", "ColorID", "BNewQuest"
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING UserID, BIntro, NowHP, MaxHP, ColorID, BNewQuest, AItemList
`

type CreateNewUserSaveParams struct {
	UserID    sql.NullString
	BIntro    sql.NullInt64
	NowHP     sql.NullInt64
	MaxHP     sql.NullInt64
	ColorID   sql.NullInt64
	BNewQuest sql.NullInt64
}

func (q *Queries) CreateNewUserSave(ctx context.Context, arg CreateNewUserSaveParams) (UserSave, error) {
	row := q.db.QueryRowContext(ctx, createNewUserSave,
		arg.UserID,
		arg.BIntro,
		arg.NowHP,
		arg.MaxHP,
		arg.ColorID,
		arg.BNewQuest,
	)
	var i UserSave
	err := row.Scan(
		&i.UserID,
		&i.BIntro,
		&i.NowHP,
		&i.MaxHP,
		&i.ColorID,
		&i.BNewQuest,
		&i.AItemList,
	)
	return i, err
}

const createNewUserScan = `-- name: CreateNewUserScan :one
INSERT INTO "userScans" (
  "UserID", "ID", "Type", "Tag", "Height", "BMulti", "LuaHash", "IsRemove"
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING UserID, ID, Type, Tag, Height, BMulti, LuaHash, IsRemove
`

type CreateNewUserScanParams struct {
	UserID   sql.NullString
	ID       sql.NullInt64
	Type     sql.NullInt64
	Tag      sql.NullString
	Height   sql.NullFloat64
	BMulti   sql.NullInt64
	LuaHash  sql.NullInt64
	IsRemove sql.NullInt64
}

func (q *Queries) CreateNewUserScan(ctx context.Context, arg CreateNewUserScanParams) (UserScans, error) {
	row := q.db.QueryRowContext(ctx, createNewUserScan,
		arg.UserID,
		arg.ID,
		arg.Type,
		arg.Tag,
		arg.Height,
		arg.BMulti,
		arg.LuaHash,
		arg.IsRemove,
	)
	var i UserScans
	err := row.Scan(
		&i.UserID,
		&i.ID,
		&i.Type,
		&i.Tag,
		&i.Height,
		&i.BMulti,
		&i.LuaHash,
		&i.IsRemove,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "DeviceID" = ?
`

func (q *Queries) DeleteUser(ctx context.Context, deviceid sql.NullString) error {
	_, err := q.db.ExecContext(ctx, deleteUser, deviceid)
	return err
}

const getUser = `-- name: GetUser :one
SELECT DeviceID, ID, Name, Status, TeamID, NewGame, UUID FROM "users"
WHERE "DeviceID" = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, deviceid sql.NullString) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, deviceid)
	var i Users
	err := row.Scan(
		&i.DeviceID,
		&i.ID,
		&i.Name,
		&i.Status,
		&i.TeamID,
		&i.NewGame,
		&i.UUID,
	)
	return i, err
}

const getUserLocalMap = `-- name: GetUserLocalMap :one
SELECT UserID, Name, Floor FROM "userLocalMap"
WHERE "UserID" = ? LIMIT 1
`

func (q *Queries) GetUserLocalMap(ctx context.Context, userid sql.NullString) (UserLocalMap, error) {
	row := q.db.QueryRowContext(ctx, getUserLocalMap, userid)
	var i UserLocalMap
	err := row.Scan(&i.UserID, &i.Name, &i.Floor)
	return i, err
}

const getUserQuestCurrent = `-- name: GetUserQuestCurrent :one
SELECT UserID, ID, Value, IsClear FROM "userQuest"
WHERE "UserID" = ? AND "IsClear" = 0 LIMIT 1
`

func (q *Queries) GetUserQuestCurrent(ctx context.Context, userid sql.NullString) (UserQuest, error) {
	row := q.db.QueryRowContext(ctx, getUserQuestCurrent, userid)
	var i UserQuest
	err := row.Scan(
		&i.UserID,
		&i.ID,
		&i.Value,
		&i.IsClear,
	)
	return i, err
}

const getUserResume = `-- name: GetUserResume :one
SELECT UserID, BResume, LuaHash, TagID, ResumeID FROM "userResume"
WHERE "UserID" = ? LIMIT 1
`

func (q *Queries) GetUserResume(ctx context.Context, userid sql.NullString) (UserResume, error) {
	row := q.db.QueryRowContext(ctx, getUserResume, userid)
	var i UserResume
	err := row.Scan(
		&i.UserID,
		&i.BResume,
		&i.LuaHash,
		&i.TagID,
		&i.ResumeID,
	)
	return i, err
}

const getUserSave = `-- name: GetUserSave :one
SELECT UserID, BIntro, NowHP, MaxHP, ColorID, BNewQuest, AItemList FROM "userSave"
WHERE "UserID" = ? LIMIT 1
`

func (q *Queries) GetUserSave(ctx context.Context, userid sql.NullString) (UserSave, error) {
	row := q.db.QueryRowContext(ctx, getUserSave, userid)
	var i UserSave
	err := row.Scan(
		&i.UserID,
		&i.BIntro,
		&i.NowHP,
		&i.MaxHP,
		&i.ColorID,
		&i.BNewQuest,
		&i.AItemList,
	)
	return i, err
}

const getUserScan = `-- name: GetUserScan :one
SELECT UserID, ID, Type, Tag, Height, BMulti, LuaHash, IsRemove FROM "userScans"
WHERE "UserID" = ? AND "ID" = ?
LIMIT 1
`

type GetUserScanParams struct {
	UserID sql.NullString
	ID     sql.NullInt64
}

func (q *Queries) GetUserScan(ctx context.Context, arg GetUserScanParams) (UserScans, error) {
	row := q.db.QueryRowContext(ctx, getUserScan, arg.UserID, arg.ID)
	var i UserScans
	err := row.Scan(
		&i.UserID,
		&i.ID,
		&i.Type,
		&i.Tag,
		&i.Height,
		&i.BMulti,
		&i.LuaHash,
		&i.IsRemove,
	)
	return i, err
}

const listUserBuildings = `-- name: ListUserBuildings :many
SELECT UserID, Prefab, Name, Status FROM "userBuildings"
WHERE "UserID" = ?
ORDER BY "Prefab"
`

func (q *Queries) ListUserBuildings(ctx context.Context, userid sql.NullString) ([]UserBuildings, error) {
	rows, err := q.db.QueryContext(ctx, listUserBuildings, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserBuildings
	for rows.Next() {
		var i UserBuildings
		if err := rows.Scan(
			&i.UserID,
			&i.Prefab,
			&i.Name,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserCharacters = `-- name: ListUserCharacters :many
SELECT UserID, CharacterId, ItemId FROM "userCharacters"
WHERE "UserID" = ?
ORDER BY "CharacterId"
`

func (q *Queries) ListUserCharacters(ctx context.Context, userid sql.NullString) ([]UserCharacters, error) {
	rows, err := q.db.QueryContext(ctx, listUserCharacters, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserCharacters
	for rows.Next() {
		var i UserCharacters
		if err := rows.Scan(&i.UserID, &i.CharacterId, &i.ItemId); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserGPS = `-- name: ListUserGPS :many
SELECT UserID, ID, Name, PinType, PinColor, Latitude, Longitude, LuaScript, BLocationEvent, QuestID, MapType, MapNo, IsRemove FROM "userGPS"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "ID"
`

func (q *Queries) ListUserGPS(ctx context.Context, userid sql.NullString) ([]UserGPS, error) {
	rows, err := q.db.QueryContext(ctx, listUserGPS, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserGPS
	for rows.Next() {
		var i UserGPS
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.Name,
			&i.PinType,
			&i.PinColor,
			&i.Latitude,
			&i.Longitude,
			&i.LuaScript,
			&i.BLocationEvent,
			&i.QuestID,
			&i.MapType,
			&i.MapNo,
			&i.IsRemove,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserGPSRemoved = `-- name: ListUserGPSRemoved :many
SELECT UserID, ID, Name, PinType, PinColor, Latitude, Longitude, LuaScript, BLocationEvent, QuestID, MapType, MapNo, IsRemove FROM "userGPS"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "ID"
`

func (q *Queries) ListUserGPSRemoved(ctx context.Context, userid sql.NullString) ([]UserGPS, error) {
	rows, err := q.db.QueryContext(ctx, listUserGPSRemoved, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserGPS
	for rows.Next() {
		var i UserGPS
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.Name,
			&i.PinType,
			&i.PinColor,
			&i.Latitude,
			&i.Longitude,
			&i.LuaScript,
			&i.BLocationEvent,
			&i.QuestID,
			&i.MapType,
			&i.MapNo,
			&i.IsRemove,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserItems = `-- name: ListUserItems :many
SELECT UserID, Item FROM "userItems"
WHERE "UserID" = ?
ORDER BY "Item"
`

func (q *Queries) ListUserItems(ctx context.Context, userid sql.NullString) ([]UserItems, error) {
	rows, err := q.db.QueryContext(ctx, listUserItems, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserItems
	for rows.Next() {
		var i UserItems
		if err := rows.Scan(&i.UserID, &i.Item); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserOnceEvent = `-- name: ListUserOnceEvent :many
SELECT UserID, UInt, IsRemove FROM "userOnceEvent"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "UInt"
`

func (q *Queries) ListUserOnceEvent(ctx context.Context, userid sql.NullString) ([]UserOnceEvent, error) {
	rows, err := q.db.QueryContext(ctx, listUserOnceEvent, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserOnceEvent
	for rows.Next() {
		var i UserOnceEvent
		if err := rows.Scan(&i.UserID, &i.UInt, &i.IsRemove); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserOnceEventRemoved = `-- name: ListUserOnceEventRemoved :many
SELECT UserID, UInt, IsRemove FROM "userOnceEvent"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "UInt"
`

func (q *Queries) ListUserOnceEventRemoved(ctx context.Context, userid sql.NullString) ([]UserOnceEvent, error) {
	rows, err := q.db.QueryContext(ctx, listUserOnceEventRemoved, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserOnceEvent
	for rows.Next() {
		var i UserOnceEvent
		if err := rows.Scan(&i.UserID, &i.UInt, &i.IsRemove); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserQuestItem = `-- name: ListUserQuestItem :many
SELECT UserID, QuestID, Name, Type FROM "userQuestItems"
WHERE "UserID" = ? AND "QuestID" = ?
ORDER BY "Name"
`

type ListUserQuestItemParams struct {
	UserID  sql.NullString
	QuestID sql.NullInt64
}

func (q *Queries) ListUserQuestItem(ctx context.Context, arg ListUserQuestItemParams) ([]UserQuestItems, error) {
	rows, err := q.db.QueryContext(ctx, listUserQuestItem, arg.UserID, arg.QuestID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserQuestItems
	for rows.Next() {
		var i UserQuestItems
		if err := rows.Scan(
			&i.UserID,
			&i.QuestID,
			&i.Name,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserQuests = `-- name: ListUserQuests :many
SELECT UserID, ID, Value, IsClear FROM "userQuest"
WHERE "UserID" = ?
ORDER BY "ID"
`

func (q *Queries) ListUserQuests(ctx context.Context, userid sql.NullString) ([]UserQuest, error) {
	rows, err := q.db.QueryContext(ctx, listUserQuests, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserQuest
	for rows.Next() {
		var i UserQuest
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.Value,
			&i.IsClear,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserScan = `-- name: ListUserScan :many
SELECT UserID, ID, Type, Tag, Height, BMulti, LuaHash, IsRemove FROM "userScans"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "ID"
`

func (q *Queries) ListUserScan(ctx context.Context, userid sql.NullString) ([]UserScans, error) {
	rows, err := q.db.QueryContext(ctx, listUserScan, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserScans
	for rows.Next() {
		var i UserScans
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.Type,
			&i.Tag,
			&i.Height,
			&i.BMulti,
			&i.LuaHash,
			&i.IsRemove,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserScanRemoved = `-- name: ListUserScanRemoved :many
SELECT UserID, ID, Type, Tag, Height, BMulti, LuaHash, IsRemove FROM "userScans"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "ID"
`

func (q *Queries) ListUserScanRemoved(ctx context.Context, userid sql.NullString) ([]UserScans, error) {
	rows, err := q.db.QueryContext(ctx, listUserScanRemoved, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserScans
	for rows.Next() {
		var i UserScans
		if err := rows.Scan(
			&i.UserID,
			&i.ID,
			&i.Type,
			&i.Tag,
			&i.Height,
			&i.BMulti,
			&i.LuaHash,
			&i.IsRemove,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserLocalMap = `-- name: UpdateUserLocalMap :exec
UPDATE "userLocalMap"
set "Name" = ?, "Floor" = ?
WHERE "UserID" = ?
`

type UpdateUserLocalMapParams struct {
	Name   sql.NullString
	Floor  sql.NullInt64
	UserID sql.NullString
}

func (q *Queries) UpdateUserLocalMap(ctx context.Context, arg UpdateUserLocalMapParams) error {
	_, err := q.db.ExecContext(ctx, updateUserLocalMap, arg.Name, arg.Floor, arg.UserID)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE "users"
set "Name" = ?
WHERE "DeviceID" = ?
`

type UpdateUserNameParams struct {
	Name     sql.NullString
	DeviceID sql.NullString
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserName, arg.Name, arg.DeviceID)
	return err
}

const updateUserNewGame = `-- name: UpdateUserNewGame :exec
UPDATE "users"
set "NewGame" = ?
WHERE "DeviceID" = ?
`

type UpdateUserNewGameParams struct {
	NewGame  sql.NullInt64
	DeviceID sql.NullString
}

func (q *Queries) UpdateUserNewGame(ctx context.Context, arg UpdateUserNewGameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserNewGame, arg.NewGame, arg.DeviceID)
	return err
}

const updateUserOnceEventRemove = `-- name: UpdateUserOnceEventRemove :exec
UPDATE "userOnceEvent"
set "IsRemove" = ?
WHERE "UserID" = ? AND "UInt" = ?
`

type UpdateUserOnceEventRemoveParams struct {
	IsRemove sql.NullInt64
	UserID   sql.NullString
	UInt     sql.NullInt64
}

func (q *Queries) UpdateUserOnceEventRemove(ctx context.Context, arg UpdateUserOnceEventRemoveParams) error {
	_, err := q.db.ExecContext(ctx, updateUserOnceEventRemove, arg.IsRemove, arg.UserID, arg.UInt)
	return err
}

const updateUserQuestClear = `-- name: UpdateUserQuestClear :exec
UPDATE "userQuest"
set "IsClear" = ?
WHERE "UserID" = ? AND "ID" = ?
`

type UpdateUserQuestClearParams struct {
	IsClear sql.NullInt64
	UserID  sql.NullString
	ID      sql.NullInt64
}

func (q *Queries) UpdateUserQuestClear(ctx context.Context, arg UpdateUserQuestClearParams) error {
	_, err := q.db.ExecContext(ctx, updateUserQuestClear, arg.IsClear, arg.UserID, arg.ID)
	return err
}

const updateUserResume = `-- name: UpdateUserResume :exec
UPDATE "userResume"
set "BResume" = ?, "LuaHash" = ?, "TagID" = ?, "ResumeID" = ?
WHERE "UserID" = ?
`

type UpdateUserResumeParams struct {
	BResume  sql.NullInt64
	LuaHash  sql.NullInt64
	TagID    sql.NullInt64
	ResumeID sql.NullInt64
	UserID   sql.NullString
}

func (q *Queries) UpdateUserResume(ctx context.Context, arg UpdateUserResumeParams) error {
	_, err := q.db.ExecContext(ctx, updateUserResume,
		arg.BResume,
		arg.LuaHash,
		arg.TagID,
		arg.ResumeID,
		arg.UserID,
	)
	return err
}

const updateUserResumeBool = `-- name: UpdateUserResumeBool :exec
UPDATE "userResume"
set "BResume" = ?
WHERE "UserID" = ?
`

type UpdateUserResumeBoolParams struct {
	BResume sql.NullInt64
	UserID  sql.NullString
}

func (q *Queries) UpdateUserResumeBool(ctx context.Context, arg UpdateUserResumeBoolParams) error {
	_, err := q.db.ExecContext(ctx, updateUserResumeBool, arg.BResume, arg.UserID)
	return err
}

const updateUserSaveIntro = `-- name: UpdateUserSaveIntro :exec
UPDATE "userSave"
set "BIntro" = ?
WHERE "UserID" = ?
`

type UpdateUserSaveIntroParams struct {
	BIntro sql.NullInt64
	UserID sql.NullString
}

func (q *Queries) UpdateUserSaveIntro(ctx context.Context, arg UpdateUserSaveIntroParams) error {
	_, err := q.db.ExecContext(ctx, updateUserSaveIntro, arg.BIntro, arg.UserID)
	return err
}

const updateUserSaveNewQuest = `-- name: UpdateUserSaveNewQuest :exec
UPDATE "userSave"
set "BNewQuest" = ?
WHERE "UserID" = ?
`

type UpdateUserSaveNewQuestParams struct {
	BNewQuest sql.NullInt64
	UserID    sql.NullString
}

func (q *Queries) UpdateUserSaveNewQuest(ctx context.Context, arg UpdateUserSaveNewQuestParams) error {
	_, err := q.db.ExecContext(ctx, updateUserSaveNewQuest, arg.BNewQuest, arg.UserID)
	return err
}

const updateUserScanLuaHash = `-- name: UpdateUserScanLuaHash :exec
UPDATE "userScans"
set "LuaHash" = ?
WHERE "UserID" = ? AND "Tag" = ?
`

type UpdateUserScanLuaHashParams struct {
	LuaHash sql.NullInt64
	UserID  sql.NullString
	Tag     sql.NullString
}

func (q *Queries) UpdateUserScanLuaHash(ctx context.Context, arg UpdateUserScanLuaHashParams) error {
	_, err := q.db.ExecContext(ctx, updateUserScanLuaHash, arg.LuaHash, arg.UserID, arg.Tag)
	return err
}

const updateUserStatus = `-- name: UpdateUserStatus :exec
UPDATE "users"
set "Status" = ?
WHERE "DeviceID" = ?
`

type UpdateUserStatusParams struct {
	Status   sql.NullInt64
	DeviceID sql.NullString
}

func (q *Queries) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateUserStatus, arg.Status, arg.DeviceID)
	return err
}

const updateUserTeam = `-- name: UpdateUserTeam :exec
UPDATE "users"
set "TeamID" = ?
WHERE "ID" = ?
`

type UpdateUserTeamParams struct {
	TeamID sql.NullString
	ID     sql.NullString
}

func (q *Queries) UpdateUserTeam(ctx context.Context, arg UpdateUserTeamParams) error {
	_, err := q.db.ExecContext(ctx, updateUserTeam, arg.TeamID, arg.ID)
	return err
}
