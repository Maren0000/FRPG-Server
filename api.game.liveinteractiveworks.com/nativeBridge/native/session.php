<?php
header('Connection', 'keep-alive');
header("Content-type: text/javascript; charset=utf-8");
header('Cache-Control', 'private, no-store, no-cache');
header('Pragma: no-cache');
ob_start("ob_gzhandler");
echo '{"sharedSecurityKey":"QetBknGIxXLS4zxfoUQIkFUKChDhvILG","nativeSessionId":"73745dbcc387e6b90e04c98b433a1320"}';
ob_end_flush();
?>