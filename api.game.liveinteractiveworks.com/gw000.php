<?php
header('Content-type: application/octet-stream');

function decrypt($packet)
{
	return openssl_decrypt($packet, 'DES-CBC', 'FrPg-109', OPENSSL_RAW_DATA, 'sX5nJEQL');
}
function encrypt($packet)
{
	return openssl_encrypt(json_encode($packet), 'DES-CBC', 'FrPg-109', OPENSSL_RAW_DATA, 'sX5nJEQL');
}

/*
RES Numbers:
SUCCESS 0;
MAINTENANCE 1;
VERSION_UP 2;
SERVICE_CLOSE 3;
FAILURE_NETWORK 1001;
FAILURE_PARAM 8000;
FAILURE_SESSION 8001;
FAILURE_MULTI_LOGIN 8002;
FAILURE_CONTINUOUS_ACCESS 8003;
FAILURE_SERVER 9000;
FAILURE_READ_FILE 9001;
*/

$websocketServer = "ws://127.0.0.1:6112";

$postRequest = file_get_contents('php://input');
if(isset($postRequest))
{
	$decryptedData = json_decode(decrypt($postRequest), true);
	file_put_contents('output.log', print_r($decryptedData, true));
	$packet = array();
	switch($decryptedData['pid'])
	{
		case 101: //P_LOGIN_START
			header('Set-Cookie: mdx=48cl46l7g5c8n3fe59vmb7gn8d; path=/;');
			$packet = array('res' => 0,
				"nativeToken" => "FHxveQcSfC+xZ9nc6B5euzQX7sUQ44PVTTNizkQRhvHrVDPnytbXqaXjGLBlfkWpHoFR\/g8g0+sgK3YTxznENOMBxJnJ7nKSVumBZCTqDcXN\/7j0NBAeARwu0kzjh0FdP9j0rXMBcqG="
			);
		break;
		case -2: //P_SESSION_CREATE
			file_put_contents('exportedRequests/P_SESSION_CREATE.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 102: //P_LOGIN_SESSION
			$packet = array('res' => 0,
				"newGame" => 1
			);
		break;
		case 103: //P_LOGIN_GAME
			/*
			[pid] => 103
			[ver] => 3
			[deviceType] => 2
			[terminalId] => Unique Device ID
			[CC] => 3
			[advertisingId] => 00000000-0000-0000-0000-000000000000
			[isTrackingEnabled] => 0		
			*/	
		
			//New player with no ticket
			/*$packet = array('res' => 0,
				"result" => 10,
				"linkUUID" => null,
				"player" => array(
					"id" => "kPjrXd",
					"name" => null
				)
			);*/
			
			//Reuslt
			/*10 - Deactivate (No Ticket)
			11 - Creating Player
			12 - Creating Party
			13 - Playing
			*/
			
			//For new player with ticket
			/*$packet = array('res' => 0,
				"result" => 11,
				"linkUUID" => "00000000-0000-0000-0000-000000000001",
				"player" => array(
					"id" => "kPjrXd",
					"name" => null
				)
			);*/
			
			$packet = array('res' => 0,
				"result" => 13,
				"linkUUID" => "00000000-0000-0000-0000-000000000001",
				"player" => array(
					"id" => "kPjrXd",
					"name" => "Maren"
				)
			);
			
			
		break;
		case 110: //P_IAP_START
			file_put_contents('exportedRequests/P_IAP_START.log', print_r($decryptedData, true));
			$packet = array('res' => 0,
				"bridgeTransId" => '',
				"result" => 0
			);
		break;
		case 111: //P_ACTIVATION
			file_put_contents('exportedRequests/P_ACTIVATION.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 120: //P_CREATE_USER
			$packet = array('res' => 0);
		break;
		case 121: //P_PARTY_CREATE
			$packet = array('res' => 0);
		break;
		case 122: //P_PARTY_START
			file_put_contents('exportedRequests/P_PARTY_START.log', print_r($decryptedData, true));
			$packet = array('res' => 0,
				'partyId' => '1',
				'webSockesServer' => $websocketServer
			);
		break;
		case 115: //P_PARTY_CANCEL
			file_put_contents('exportedRequests/P_PARTY_CANCEL.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 116: //P_PARTY_STATE
			file_put_contents('exportedRequests/P_PARTY_STATE.log', print_r($decryptedData, true));
			$packet = array('res' => 0,
				"state" => 0
			);
		break;
		case 117: //P_PARTY_KICK
			file_put_contents('exportedRequests/P_PARTY_KICK.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 200: //P_HOME
			$packet = array('res' => 0,
				"aTeamUser" => array(
					array(
						"id" => "kPjrXd",
						"name" => "Maren"
					)
				),
				"bIntro" => boolval(true),
				"eventInfo" => array(
					"bResume" => boolval(false),
					"lua" => intval(crc32('ComicEvent/introduction.lua')),
					"resumeId" => intval(0),
					"tagId" => intval(0)
				),
				"roomId" => "1",
				"teamId" => "1",
				"teamName" => "Maren's Team",
				"webSockesServer" => $websocketServer,
				"aCharacter" => array(),
				"aScan" => array(),
				"aRemoveScan" => array(),
				"aGPS" => array(
					array(
						"id" => intval(1),
						"name" => "marui",
						"pinType" => "main",
						"pinColor" => "y",
						"latitude" => floatval(35.646530),
						"longitude" => floatval(139.708430),
						"LuaScript" => intval(crc32('ComicEvent/oioi8f/oioi8f_ev_deai_000.lua')),
						"bLocationEvent" => intval(0),
						"quest" => array(
							"id" => 1,
							"value" => 1
						),
						"mapType" => "GPSMap",
						"mapNo" => "",			
					)
				),
				"aRemoveGPS" => array(),
				"aOnceEvent" => array(),
				"aRemoveOnceEvent" => array(),
				"aBuildings" => array(),
				"nowHp" => intval(50),
				"maxHp" => intval(100),
				"colorId" => intval(25),
				"quest" => array(
					"id" => 1,
					"value" => 1
				),
				"bNewQuest" => boolval(true),
				"aItemList" => array(),
				"localMap" => array()
			);
		break;
		case 201: //P_EVENT
			$packet = array('res' => 0,
				"aCharacter" => array(),
				"aScan" => array(),
				"aRemoveScan" => array(),
				"aGPS" => array(
					array(
						"id" => intval(1),
						"name" => "marui",
						"pinType" => "main",
						"pinColor" => "y",
						"latitude" => floatval(35.646530),
						"longitude" => floatval(139.708430),
						"LuaScript" => intval(crc32('ComicEvent/oioi8f/oioi8f_ev_deai_000.lua')),
						"bLocationEvent" => intval(0),
						"quest" => array(
							"id" => 1,
							"value" => 1
						),
						"mapType" => "GPSMap",
						"mapNo" => "",			
					)
				),
				"aRemoveGPS" => array(),
				"aOnceEvent" => array(),
				"aRemoveOnceEvent" => array(),
				"aBuildings" => array(),
				"nowHp" => intval(50),
				"maxHp" => intval(100),
				"colorId" => intval(25),
				"quest" => array(
					"id" => 1,
					"value" => 1
				),
				"bNewQuest" => boolval(false),
				"aItemList" => array(),
				"localMap" => array()
			);
		break;
		case 202: //P_EVENT_SAVE_RESUME
			file_put_contents('exportedRequests/P_EVENT_SAVE_RESUME.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 203: //P_SCAN_TAG
			file_put_contents('exportedRequests/P_SCAN_TAG.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 204: //P_QUEST_CHECKED
			file_put_contents('exportedRequests/P_QUEST_CHECKED.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);	
		break;
		case 205: //P_GET_QUEST_INFO
			file_put_contents('exportedRequests/P_GET_QUEST_INFO.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 206: //P_EVENT_CHECK_RESUME
			file_put_contents('exportedRequests/P_EVENT_CHECK_RESUME.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 300: //P_BATTLE_IN
			file_put_contents('exportedRequests/P_BATTLE_IN.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 301: //P_BATTLE_RESULT
			file_put_contents('exportedRequests/P_BATTLE_RESULT.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 302: //P_BATTLE_ATTACK
			file_put_contents('exportedRequests/P_BATTLE_ATTACK.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 900: //P_SHOP_BENEFIT
			file_put_contents('exportedRequests/P_SHOP_BENEFIT.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 901: //P_SHOP_IDENTIFY_START
			file_put_contents('exportedRequests/P_SHOP_IDENTIFY_START.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 902: //P_SHOP_IDENTIFY_END
			file_put_contents('exportedRequests/P_SHOP_IDENTIFY_END.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 1000: //P_DATA_LINKAGE_SESSION_SIGN_UP
			file_put_contents('exportedRequests/P_DATA_LINKAGE_SESSION_SIGN_UP.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 1001: //P_DATA_LINKAGE_SESSION_ADD_TERMINAL
			file_put_contents('exportedRequests/P_DATA_LINKAGE_SESSION_ADD_TERMINAL.log', print_r($decryptedData, true));
			$packet = array('res' => 0,
				"url" => "http://127.0.0.1:2090/PurchaseTicket.txt",
			);
		break;
		case 1002: //P_DATA_LINKAGE_GET_LINK_UUID
			file_put_contents('exportedRequests/P_DATA_LINKAGE_GET_LINK_UUID.log', print_r($decryptedData, true));
			$packet = array('res' => 0,
				"linkUUID" => "00000000-0000-0000-0000-000000000001",
			);
		break;
		case 1100: //P_NG_WORD
			$packet = array('res' => 0,
				"bOK" => 1
			);
		break;
		case 9000: //P_NETWORK_TEST
			file_put_contents('exportedRequests/P_NETWORK_TEST.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
		case 9001: //P_DEBUG_ACTIVATION
			file_put_contents('exportedRequests/P_DEBUG_ACTIVATION.log', print_r($decryptedData, true));
			$packet = array('res' => 1,
			
			);
		break;
	}
	echo encrypt($packet);
}
?>