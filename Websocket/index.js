const http = require('http');
const ws = require('ws');
const CRC32 = require('crc32');

let server = http.createServer({}, (req, res) => {});
server.listen(6113, () => console.log('MQTT running on port ' + 6113));

function parseData(data)
{
    let rawData = data.toString();
    console.log("Raw Data: " + rawData);
    let parse1 = rawData.split("[");
    let returndata = {
        'unk': parse1[0],
        //'data': []
    };
    if(parse1.length > 1)
    {
        parse1[1] = parse1[1].replace("]", "").replace("{", "").replace("}", "");
        let splitData2 = parse1[1].split(",");

        returndata['evName'] = splitData2[0].replace('"', '');
        if(splitData2.length > 1)
        {
            returndata['str'] = splitData2[1].split(":")[1].replace('"', '');
            returndata['crc'] = splitData2[2].split(":")[1].replace('"', '');

            /*for(let i = 1; i < splitData2.length; i++)
            {
                let splitKeyValue = splitData2[i].split(":");
                returndata['data'][i - 1] = {
                    'key': splitKeyValue[0],
                    'value': splitKeyValue[1]
                }
            }*/
        }
    }
    return returndata;
}

const wss = new ws.Server({server, path: '/'});
wss.on('connection', function connection(ws) {
    ws.on('message', (data) => {
        let requestData = parseData(data);
        console.log(requestData);

        let returnData = '';
        switch(requestData['unk'])
        {
            case 2:
                returnData = "3KYJWzVFudDQypXv+W2TAQA8JmrcyfL8ibR4pPkG49k9cgStC56nuIqleX2693xehk40zlKjXT0=";     
            break;
            case 42:

            break;
        }

        ws.send(Buffer.from('2["maintenance",{"data":"'+returnData+'","crc":"'+CRC32(returnData)+'"}]', 'hex'));
        //To-Do Parse data properly
    });
});