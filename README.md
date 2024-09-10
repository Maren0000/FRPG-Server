# FRPG-Server

Back in 2021, a Japan exclusive mobile game made to promote the release of **NEO: The World Ends with You** was released called *NEO: The World Ends with You x Field Walk RPG*. It only lasted from September to August as limited time campaign. It heavily relied on an online server that was closed soon after the campaign ended.

FRPG-Server is a lightweight, easy to setup API server meant to bring the game back to life. The game can be fully played from start to finish, including all the side missions, using the official QR codes that were used on the original server.

## Missing Features

Almost all the game features have been fully recreated in this server. However, there are currently two game features that are missing:

1. Although the server has support for teams in the DB, the team system in-game uses a different communication protocol than the main API. So, for now teams aren't supported and only Solo mode works in-game.
2. The game features TFLITE object recognition for scanning. Currently the hardcoded strings needed to get this to work are unknown, so for now all the scans will need to be done using QR codes.

## Server Building and Setup

Setup of the server is very simple. There are currently two ways to setup it up:

### 1. Manual Setup

1. Install Go 1.23 or higher
2. Clone the repo
3. Run the command `go build`
4. Setup a `.env` file based on the example file provided
5. Run the executable

That's it! The server is now running and clients can now connect to the running servers.

### 2. Docker Setup

A dockerfile and docker-compose files are also provided and customizable if that's preferred. Just be sure to place the SQLite DB in a place that's saved safely using the `CUSTOM_SQLITE_PATH` and `SQLITE_PATH` settings in the env file.

## Client Setup

For client setup, please read the `Docs/client_setup.md` file.
