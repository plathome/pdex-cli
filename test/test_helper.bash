#!/usr/bin/env bash

setup() {
  export PDEX_END_POINT_URL=http://localhost:9292/api/v1
  export PDEX_ACCESS_KEY=c15722ed5cd1
  pdex="bin/pdex"
}

# assign the device group
create_dg() {
	DGRESULT=$(pdex cr dg)
	DG=$(echo $DGRESULT | jq .deid_prefix)
}

# lust devicegroups
list_dgs() {
	create_dg
	CMD="pdex ls dg"
	LISTDG=$(eval $CMD | jq .count)
}

# create device
create_device() {
	create_dg
	CMD="pdex cr devices --deid-prefix $DG"
	DERESULT=$(eval $CMD)
	DE=$(echo $DERESULT | jq .deid)
}

# list devices
list_devices() {
	create_channel
	CMD="pdex ls de --deid-prefix $DG"
	DELISTRESULT=$(eval $CMD)
	DELIST=$(echo $DELISTRESULT | jq .count)
}

# create app
create_app() {
	sleep 1
	TAG=$(awk -v min=1 -v max=99999 'BEGIN{srand(); print int(min+rand()*(max-min+1))}')
	TIME=$(($(date +'%s * 1000 + %-N / 1000000')))
	CMD="pdex cr apps --app-name-suffix $TAG-$TIME"
#	echo CMD
	APPRESULT=$(eval $CMD)
#	echo $APPRESULT
#	echo $CMD
	APP=$(echo $APPRESULT | jq .app_id)
}

# list app
list_apps() {
	create_app
	CMD="pdex ls apps"
	APPLISTRESULT=$(eval $CMD)
	APPLIST=$(echo $APPLISTRESULT | jq .count)
}

# create channel
create_channel() {
	create_device
	create_app
	CMD="pdex cr ch --deid $DE --app-id $APP"
	CHRESULT=$(eval $CMD)
	CHID=$(echo $CHRESULT | jq .channel_id)
}

# list channel
list_channel() {
	create_channel
	CMD="pdex ls ch --deid $DE"
	LISTCHRESULT=$(eval $CMD)
	LISTCH=$(echo $LISTCHRESULT | jq .count)
}

# send message
send_message() {
	create_channel
  	CMD="pdex send msg --deid $DE 'message send for test'"
  	XFER=$(eval $CMD)
}

# read message
read_message() {
	send_message
	CMD="pdex read msg --app-id $APP"
	READRESULT=$(eval $CMD)
	READ=$(echo $READRESULT | jq .count)
	MSGID=$(echo $READRESULT | jq '.messages[0].msgid')
}

# Read single message
read_single_message() {
	read_message
	CMD="pdex read msg --app-id $APP --msgid $MSGID"
	READ=$(eval $CMD)
}

# send command
send_command() {
  create_channel
  CMD="pdex send cmd --channel-id $CHID --app-id $APP 'test command sending'"
  XFER=$(eval $CMD)
}

# read command
read_command() {
	send_command
	CMD="pdex read cmd --deid $DE"
	READCMDRESULT=$(eval $CMD)
	READ=$(echo $READCMDRESULT | jq .count)
	CMDID=$(echo $READCMDRESULT | jq '.commands[0].cmdid')
}

# command contetnt check
read_single_command() {
	read_command
	CMD="pdex read cmd --deid $DE --cmdid $CMDID"
	READ=$(eval $CMD)
}

# send bulk commands
send_command_bulk() {
  create_channel
  CMD="pdex send cmd --channel-id $CHID --app-id $APP 'test command sending'"
  for i in `seq 1 10`;
  do
	  XFER=$(eval $CMD)
  done
}

# send bulk messages
send_messages_bulk() {
	create_channel
  	CMD="pdex send msg --deid $DE 'message send for test'"
	for i in `seq 1 10`;
	do
	  XFER=$(eval $CMD)
	done
}

# read bulk messages
read_messages_bulk() {
	send_messages_bulk
	CMD="pdex read msg --app-id $APP"
	READ=$(eval $CMD | jq .count)
}

# send bulk commands
send_commands_bulk() {
  create_channel
  CMD="pdex send cmd --channel-id $CHID --app-id $APP 'test command sending'"
  for i in `seq 1 10`;
  do
	  XFER=$(eval $CMD)
	  echo $XFER
  done
}

# read bulk commands
read_commands_bulk() {
	send_commands_bulk
	CMD="pdex read cmd --deid $DE"
	READ=$(eval $CMD | jq .count)
}

# app name change
app_name_change() {
	create_app
	sleep 1
	TAG=$(awk -v min=1 -v max=99999 'BEGIN{srand(); print int(min+rand()*(max-min+1))}')
	TIME=$(($(date +'%s * 1000 + %-N / 1000000')))
	CMD="pdex up apps --app-id $APP --app-name-suffix  updated-$TAG-$TIME"
	UPDATEAPP=$(eval $CMD)
}

# delete channel
delete_channel() {
	create_channel
	CMD="pdex delete ch --channel-id $CHID --deid $DE --confirm true"
	CLOSECH=$(eval $CMD)
	CLOSECHAGAIN=$(eval $CMD)

}

# create after delete and send message then read message
channel_create_after_delete() {
 	create_channel
 	delete_channel
 	read_command
}

# devicegroup tag create
create_devicegroups_tags() {
	create_dg
	CMD="pdex cr dg-tags --deid-prefix $DG --key DGROUP --value OpenBlocks"
	DGTAGS=$(eval $CMD)
}

# devicegroup tag list after create
list_devicegroups_tags() {
	create_devicegroups_tags
	CMD="pdex ls dg-tags --deid-prefix $DG"
	DGTAGSLIST=$(eval $CMD)
}

# devicegroup tag update
update_devicegroups_tags() {
	create_devicegroups_tags
	CMD="pdex up dg-tags --deid-prefix $DG --key DGROUP --value OpenBlocksNext"
	DELIST=$(eval $CMD)
	CMD="pdex ls dg-tags --deid-prefix $DG"
	DGTAGSLIST=$(eval $CMD)
}

# devicegroup tag delete
delete_devicegroups_tags() {
	create_devicegroups_tags
	CMD="pdex delete dg-tags --deid-prefix $DG --key DGROUP"
	DELIST=$(eval $CMD)
	CMD="pdex ls dg-tags --deid-prefix $DG"
	DGTAGSLIST=$(eval $CMD)
}

# device tag create
create_device_tags() {
	create_device
	create_channel
	CMD="pdex cr de-tags --deid $DE --key brand --value VX-1"
	DETAGS=$(eval $CMD)
}

# device tag list after create
list_device_tags() {
	create_device_tags
	CMD="pdex ls de-tags --deid $DE"
	DELIST=$(eval $CMD)
}

# device tag update
update_device_tags() {
	create_device_tags
	CMD="pdex up de-tags --deid $DE --key brand --value BX-1"
	DELIST=$(eval $CMD)
	CMD="pdex ls de-tags --deid $DE"
	DELIST=$(eval $CMD)
}

# device tag delete
delete_device_tags() {
	create_device_tags
	CMD="pdex delete de-tags --deid $DE --key brand"
	DETAG=$(eval $CMD)
	CMD="pdex ls de-tags --deid $DE"
	DELIST=$(eval $CMD)
}

# app tag create
create_app_tags() {
	create_app
	CMD="pdex cr app-tags --app-id $APP --key name --value plc"
	APPTAG=$(eval $CMD)
	CMD="pdex ls app-tags --app-id $APP"
	APPTAGS=$(eval $CMD)
}

# app tag list after create
list_app_tags() {
	create_app_tags
	CMD="pdex ls app-tags --app-id $APP"
	APPTAGS=$(eval $CMD)
}
# app tag update
update_app_tags() {
	create_app_tags
	CMD="pdex up app-tags --app-id $APP --key name --value humidity"
	APPLIST=$(eval $CMD)
	CMD="pdex ls app-tags --app-id $APP"
	APPLIST=$(eval $CMD)
}
# app tag delete
delete_app_tags() {
	create_app_tags
	CMD="pdex delete app-tags --app-id $APP --key name"
	echo $CMD
	APPTAG=$(eval $CMD)
	CMD="pdex ls app-tags --app-id $APP"
	echo $CMD
	APPLIST=$(eval $CMD)
}