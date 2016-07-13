#!/usr/bin/env bash

setup() {
  export PDEX_END_POINT_URL=http://localhost:9292/api/v1
  export PDEX_ACCESS_KEY=c15722ed5cdb
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
}
