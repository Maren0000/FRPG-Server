CREATE TABLE "users" (
	"DeviceID"	TEXT,
	"ID"	TEXT,
	"Name"	TEXT,
	"Status"	INTEGER,
	"TeamID"	TEXT,
	"NewGame"	INTEGER,
	"UUID"	TEXT,
	FOREIGN KEY("TeamID") REFERENCES "teams"("TeamID") ON DELETE SET NULL ON UPDATE CASCADE,
	PRIMARY KEY("ID")
);

CREATE TABLE "teams" (
	"TeamID"	TEXT,
	"RoomID"	TEXT,
	"TeamName"	TEXT,
	PRIMARY KEY("TeamID")
);

CREATE TABLE "userBuildings" (
	"UserID"	TEXT,
	"Prefab"	TEXT,
	"Name"	TEXT,
	"Status"	TEXT,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userCharacters" (
	"UserID"	TEXT,
	"CharacterId"	INTEGER,
	"ItemId"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userGPS" (
	"UserID"	TEXT,
	"ID"	INTEGER	UNIQUE PRIMARY KEY	AUTOINCREMENT,
	"Name"	TEXT,
	"PinType"	TEXT,
	"PinColor"	TEXT,
	"Latitude"	REAL,
	"Longitude"	REAL,
	"LuaScript"	INTEGER,
	"BLocationEvent"	INTEGER,
	"QuestID"	INTEGER,
	"MapType"	TEXT,
	"MapNo"	TEXT,
	"IsRemove"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userLocalMap" (
	"UserID"	TEXT UNIQUE,
	"Name"	TEXT,
	"Floor"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE,
	PRIMARY KEY("UserID")
);

CREATE TABLE "userOnceEvent" (
	"UserID"	TEXT,
	"UInt"	INTEGER,
	"IsRemove"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userQuest" (
	"UserID"	TEXT,
	"ID"	INTEGER,
	"Value"	INTEGER,
	"IsClear"	INTEGER,
	PRIMARY KEY("ID"),
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userQuestItems" (
	"UserID"	TEXT,
	"QuestID"	INTEGER,
	"Name"	TEXT,
	"Type"	TEXT,
	FOREIGN KEY("QuestID") REFERENCES "userQuest"("ID") ON DELETE CASCADE ON UPDATE CASCADE,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userSave" (
	"UserID"	TEXT UNIQUE,
	"BIntro"	INTEGER,
	"NowHP"	INTEGER,
	"MaxHP"	INTEGER,
	"ColorID"	INTEGER,
	"BNewQuest"	INTEGER,
	"BattleID"	INTEGER,
	"BattleBadge1"	INTEGER,
	"BattleBadge2"	INTEGER,
	"BattleBadge3"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE,
	PRIMARY KEY("UserID")
);

CREATE TABLE "userScans" (
	"UserID"	TEXT,
	"ID"	INTEGER	UNIQUE PRIMARY KEY	AUTOINCREMENT,
	"Type"	INTEGER,
	"Tag"	TEXT,
	"Height"	REAL,
	"BMulti"	INTEGER,
	"LuaHash"	INTEGER,
	"IsRemove"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userItems" (
	"UserID"	TEXT,
	"Item"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE "userResume" (
	"UserID"	TEXT UNIQUE,
	"BResume"	INTEGER,
	"LuaHash"	INTEGER,
	"TagID"	INTEGER,
	"ResumeID"	INTEGER,
	FOREIGN KEY("UserID") REFERENCES "users"("ID") ON DELETE CASCADE ON UPDATE CASCADE,
	PRIMARY KEY("UserID")
);