'use strict'

const Bleacon     = require('bleacon')
var   sqlite3     = require('sqlite3').verbose();
var   db          = new sqlite3.Database('base.db');
var   exec        = require('child_process').exec;
const startedAt   = new Date().getTime()
const table_name  = 'sensor_master'
const devicegroup = "01.72da6d"
const app_id      = "4817e8ee00814e93af7a59c80b8625f9"

function pad(str, len) {
  while (str.length < len) {
    str = '0' + str
  }
  return str
}

Bleacon.on('discover', (beacon) => {
  const elapsed = new Date().getTime() - startedAt
  const uuid    = beacon.uuid
  const major   = pad(beacon.major.toString(16), 4)
  const minor   = pad(beacon.minor.toString(16), 4)
  let   info    = `${elapsed}: ${uuid} | ${major} | ${minor} | ${beacon.rssi} | ${beacon.proximity} | ${beacon.accuracy} | ${beacon.measuredPower}`
  let   message = `'{${elapsed}:${uuid}, major:${major}, minor:${minor}, rssi:${beacon.rssi}, proximity:${beacon.proximity}, accuracy:${beacon.accuracy}, txpower:${beacon.measuredPower}}'`
  const deid    = ""
  const chid    = ""

  db.serialize(function () {
    var db_create = new Promise(function (resolve, reject) {
      db.get('select count(*) from sqlite_master where type="table" and name=$name',{ $name: table_name }, function (err, res) {
        var exists = false;
        if (0 < res['count(*)']) { exists = true; }

        resolve(exists);
      });
    });

    db_create.then(function (exists) {
      if (!exists) {
        db.run('create table sensor_master (id integer primary key autoincrement, uuid text, deid text)');
        console.log("table created")
      } else {
        var table_create = new Promise(function (resolve, reject) {
          db.get('select count(*) from sensor_master where uuid=$uuid',{ $uuid: uuid }, function (err, res) {
            var table_exists = false;
            if (0 < res['count(*)']) { table_exists = true; }

            resolve(table_exists);
          });
        });

        table_create.then(function (table_exists) {
          if (!table_exists) {
            let create_deid=`pdex cr de --deid-prefix ${devicegroup} | jq -r .deid`
            exec(create_deid, function(error, deid, stderr) {
              deid = deid.trim()
              let sql_query = `insert into sensor_master(uuid, deid) values('${uuid}','${deid}')`
              db.run(sql_query)
              let create_channel = `pdex cr ch --deid '${deid}' --app-id '${app_id}' | jq -r .channel_id`
              exec(create_channel, function(error, chid, stderr) {
                console.log(`new beacon ${deid} registered with channel-id: ${chid}`)
                if (error !== null) {
                    console.log('exec error: ' + error);
                }
              });
              if (error !== null) {
                console.log("execution error " + error)
              }
            });
          } else {
            db.get('select deid from sensor_master where uuid=$uuid',{ $uuid: uuid }, function (err, res) {
              let command = `pdex s msg --deid ${res['deid']} `
              command += message
              exec(command, function(error, exam, stderr) {
                console.log(`beacon ${res['deid']} is sending message`)
                if (error !== null) {
                    console.log('exec error: ' + error);
                }
              });

            });
          }
        });
      }
    });
  });
})

Bleacon.startScanning()

console.log('Listening for iBeacons...')
