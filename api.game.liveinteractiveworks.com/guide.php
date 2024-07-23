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

/*Request Params:
{"pid":-1,"ver":3,"deviceType":2,"terminalId":"Device Unique ID","CC":0}
*/

$postRequest = file_get_contents('php://input');
if(isset($postRequest))
{
	$decryptedData = json_decode(decrypt($postRequest), true);
	
	$packet = array();
	switch($decryptedData['pid'])
	{
		case -1: //P_GET_URL
			$packet = array('res' => 0,
				'api' => 'http://127.0.0.1:2090/gw000.php',
				'nsc' => 'http://127.0.0.1:2090/nativeBridge/native/session.php',
				'mainteUrl' => 'https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki',
				'termsUrl' => 'https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/terms/',
				'privacyUrl' => 'https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/privacy/',
				'licenseUrl' => 'https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/license',
				'commerceUrl' => 'https://www.liveinteractiveworks.com/contents/frpg/shinsubarashiki/commerce',
				'ticketUrl' => 'http://127.0.0.1:2090/ticket.php',
				'storeUrl' => 'https://play.google.com/store/apps/details?id=com.square_enix.android_googleplay.NTWEWYXFIELDWALKRPG',
				'eventEndDate' => '1937679599',
				//'eventEndDate' => '1637679599',			
			);
		break;
	}
	echo encrypt($packet);
}
?>