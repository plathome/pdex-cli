'use strict'

const Bleacon     = require('bleacon')
var   sqlite3     = require('sqlite3').verbose();
var   db          = new sqlite3.Database('base.db');
var   exec        = require('child_process').exec;
const startedAt   = new Date().getTime()
var   dateFormat  = require('dateformat');
const table_name  = 'sensor_master'
const devicegroup = "01.fc6b44"
const app_id      = "9bbfc362a7784c769a5be1bf23647146"

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
              let device_id = `${res['deid']}`
              let command = `pdex s msg --deid ${device_id} `
              var now = new Date();
              var event_time = dateFormat(now, "yyyymmddHHMMss");
              let message = `'{ time:${event_time}, deid:${device_id}, major:${major}, minor:${minor}, rssi:${beacon.rssi}, proximity:${beacon.proximity}, accuracy:${beacon.accuracy}, txpower:${beacon.measuredPower}}'`
              command += message
              exec(command, function(error, exam, stderr) {
                if(exam.indexOf("error") > -1) {
                  console.log(`error in the server-side : ${res['deid']} `)
                } else {
                  console.log(`beacon ${res['deid']} is sending message`)
                }

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
