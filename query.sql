-- name: GetUser :one
SELECT * FROM "users"
WHERE "DeviceID" = ? LIMIT 1;

-- name: CreateNewUser :one
INSERT INTO "users" (
  "DeviceID", "ID", "NewGame", "UUID", "Status"
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUserName :exec
UPDATE "users"
set "Name" = ?
WHERE "DeviceID" = ?;

-- name: UpdateUserStatus :exec
UPDATE "users"
set "Status" = ?
WHERE "DeviceID" = ?;

-- name: UpdateUserNewGame :exec
UPDATE "users"
set "NewGame" = ?
WHERE "DeviceID" = ?;

-- name: UpdateUserTeam :exec
UPDATE "users"
set "TeamID" = ?
WHERE "ID" = ?;

-- name: DeleteUser :exec
DELETE FROM "users"
WHERE "DeviceID" = ?;

-- name: CreateNewTeam :one
INSERT INTO "teams" (
  "TeamID", "RoomID", "TeamName"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: CreateNewUserBuilding :one
INSERT INTO "userBuildings" (
  "UserID", "Prefab", "Name", "Status"
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: ListUserBuildings :many
SELECT * FROM "userBuildings"
WHERE "UserID" = ?
ORDER BY "Prefab";

-- name: CreateNewUserCharacter :one
INSERT INTO "userCharacters" (
  "UserID", "CharacterId", "ItemId"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: ListUserCharacters :many
SELECT * FROM "userCharacters"
WHERE "UserID" = ?
ORDER BY "CharacterId";

-- name: CreateNewUserGPS :one
INSERT INTO "userGPS" (
  "UserID", "ID", "Name", "PinType", "PinColor", "Latitude", "Longitude", "LuaScript", "BLocationEvent", "QuestID", "MapType", "MapNo", "IsRemove"
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ListUserGPS :many
SELECT * FROM "userGPS"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "ID";

-- name: ListUserGPSRemoved :many
SELECT * FROM "userGPS"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "ID";

-- name: CreateNewUserLocalMap :one
INSERT INTO "userLocalMap" (
  "UserID", "Name", "Floor"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: GetUserLocalMap :one
SELECT * FROM "userLocalMap"
WHERE "UserID" = ? LIMIT 1;

-- name: CreateNewUserOnceEvent :one
INSERT INTO "userOnceEvent" (
  "UserID", "UInt", "IsRemove"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: ListUserOnceEvent :many
SELECT * FROM "userOnceEvent"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "UInt";

-- name: UpdateUserOnceEventRemove :exec
UPDATE "userOnceEvent"
set "IsRemove" = ?
WHERE "UserID" = ? AND "UInt" = ?;

-- name: ListUserOnceEventRemoved :many
SELECT * FROM "userOnceEvent"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "UInt";

-- name: CreateNewUserQuest :one
INSERT INTO "userQuest" (
  "UserID", "ID", "Value"
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: GetUserQuest :one
SELECT * FROM "userQuest"
WHERE "UserID" = ? LIMIT 1;

-- name: CreateNewUserQuestItem :one
INSERT INTO "userQuestItems" (
  "UserID", "QuestID", "Name", "Type"
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: ListUserQuestItem :many
SELECT * FROM "userQuestItems"
WHERE "UserID" = ? AND "QuestID" = ?
ORDER BY "Name";

-- name: CreateNewUserSave :one
INSERT INTO "userSave" (
  "UserID", "BIntro", "NowHP", "MaxHP", "ColorID", "BNewQuest"
) VALUES (
  ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: GetUserSave :one
SELECT * FROM "userSave"
WHERE "UserID" = ? LIMIT 1;

-- name: UpdateUserSaveIntro :exec
UPDATE "userSave"
set "BIntro" = ?
WHERE "UserID" = ?;

-- name: UpdateUserSaveNewQuest :exec
UPDATE "userSave"
set "BNewQuest" = ?
WHERE "UserID" = ?;

-- name: CreateNewUserScan :one
INSERT INTO "userScans" (
  "UserID", "ID", "Type", "Tag", "Height", "BMulti", "LuaHash", "IsRemove"
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: ListUserScan :many
SELECT * FROM "userScans"
WHERE "UserID" = ? AND "IsRemove" = 0
ORDER BY "ID";

-- name: GetUserScan :one
SELECT * FROM "userScans"
WHERE "UserID" = ? AND "ID" = ?
LIMIT 1;

-- name: ListUserScanRemoved :many
SELECT * FROM "userScans"
WHERE "UserID" = ? AND "IsRemove" = 1
ORDER BY "ID";

-- name: CreateNewUserItem :one
INSERT INTO "userItems" (
  "UserID", "Item"
) VALUES (
  ?, ?
)
RETURNING *;

-- name: ListUserItems :many
SELECT * FROM "userItems"
WHERE "UserID" = ?
ORDER BY "Item";

-- name: GetUserResume :one
SELECT * FROM "userResume"
WHERE "UserID" = ? LIMIT 1;

-- name: CreateNewUserResume :one
INSERT INTO "userResume" (
  "UserID", "BResume", "LuaHash", "TagID", "ResumeID"
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUserResume :exec
UPDATE "userResume"
set "BResume" = ?, "LuaHash" = ?, "TagID" = ?, "ResumeID" = ?
WHERE "UserID" = ?;

-- name: UpdateUserResumeBool :exec
UPDATE "userResume"
set "BResume" = ?
WHERE "UserID" = ?;