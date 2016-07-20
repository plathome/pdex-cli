#!/usr/bin/env bash

setup() {
  export PDEX_END_POINT_URL=http://localhost:9292/api/v1
  export PDEX_ACCESS_KEY=c15722ed5cd1
  pdex="bin/pdex"
}

# assign the device group
create_dg() {
	DG=$(pdex cr dg | jq .deid_prefix)
}

# lust devicegroups
list_dgs() {
	create_dg
	CMD="pdex ls dg"
	LIST=$(eval $CMD | jq .count)
}

# create device
create_device() {
	create_dg
	CMD="pdex cr devices --deid-prefix $DG"
	DE=$(eval $CMD | jq .deid)
}

# list devices
list_devices() {
	create_channel
	CMD="pdex ls de --deid-prefix $DG"
	LIST=$(eval $CMD | jq .count)
}

# create app
create_app() {
	APP=$(pdex cr apps --app-name-suffix 'test-create-channel' | jq .app_id)
}

# list app
list_apps() {
	create_app
	CMD="pdex ls apps"
	LIST=$(eval $CMD | jq .count)
}

# create channel
create_channel() {
	create_device
	create_app
	CMD="pdex cr ch --deid $DE --app-id $APP"
	CHID=$(eval $CMD | jq .channel_id)
	CH=$(eval $CMD)
}

# list channel
list_channel() {
	create_channel
	CMD="pdex ls ch --deid $DE"
	LIST=$(eval $CMD | jq .count)
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
	READ=$(eval $CMD | jq .count)
	MSGID=$(eval $CMD | jq '.messages[0].msgid')
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
	READ=$(eval $CMD | jq .count)
	CMDID=$(eval $CMD | jq '.commands[0].cmdid')
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

# delete channel

# create after delete and send message then read message

# devicegroup tag create

# devicegroup tag list after create

# devicegroup tag update

# devicegroup tag list after update

# devicegroup tag delete

# device tag list after delete

# device tag create

# device tag list after create

# device tag update

# device tag list after update

# device tag delete

# device tag list after delete


# app tag create

# app tag list after create

# app tag update

# app tag list after update

# app tag delete

# app tag list after delete



