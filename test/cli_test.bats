#!/usr/bin/env bats

load test_helper

@test "setup" {
  run $pdex c set --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
  [ "$output" = "successfully configured" ]
}

@test "profile" {
  run $pdex c profile --name new --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
  [ "$output" = "successfully configured" ]
}

# Create API list

@test "create app" {
  create_app
  [[ $APPRESULT == *"app_id"* ]]
}

@test "create device" {
  create_device
  [[ $DERESULT == *"deid"* ]]
}

@test "create devicegroup" {
  create_dg
  [[ $DGRESULT == *"deid_prefix"* ]]
}

@test "create channels" {
  create_channel
  [[ $CHRESULT == *"channel_id"* ]]
}

# Listing API

@test "list apps" {
  list_apps
  [[ $APPLIST -ge 1 ]]
}

@test "list devicegroups" {
  list_dgs
  [[ $LISTDG -ge 1 ]]
}

@test "list devices" {
  list_devices
  [[ $DELIST -eq 1 ]]
}

@test "list channels" {
  list_channel
  [[ $LISTCH -eq 1 ]]
}

# Show API
@test "show me" {
  run $pdex sh me
  echo $output | grep $PDEX_ACCESS_KEY
}

# Util API
@test "ping" {
  run $pdex u ping
  [ "$output" = "{\"result\":\"pong\"}" ]
}

# Send and Read API
@test "send message" {
  send_message
  [[ $XFER == *"transaction_id"* ]]
}

@test "send command" {
  send_command
  [[ $XFER == *"transaction_id"* ]]
}

@test "read commands" {
  read_command
  [[ READ -eq 1 ]]
}

@test "read single command" {
  read_single_command
  [[ $READ == "test command sending" ]]
}

@test "sending bulk commands and then read" {
  read_commands_bulk
  [[ $READ -eq 10 ]]
}

@test "sending bulk messages and then read" {
  read_messages_bulk
  [[ $READ -eq 10 ]]
}

@test "read messages" {
  read_message
  [[ $READ -eq 1 ]]
}

@test "read single message" {
  read_single_message
  [[ $READ == "message send for test" ]]
}

@test "update app" {
  app_name_change
  [[ $UPDATEAPP == *"updated"* ]]
}

@test "close channel" {
  delete_channel
  [[ $CLOSECH == *"channel_id"* ]]
  [[ $CLOSECHAGAIN == *"Channel not found"* ]]
}

@test "create device tags" {
  create_device_tags
  [[ $DETAGS == *"VX-1"* ]]
}

@test "list-up device tags" {
  list_device_tags
  [[ $DELIST == *"VX-1"* ]]
}

@test "update device tags" {
  update_device_tags
  [[ $DELIST == *"BX-1"* ]]
}

@test "delete device tag" {
  delete_device_tags
  [[ $DELIST == *"[]"* ]]
}

@test "create dg tags" {
  create_devicegroups_tags
  [[ $DGTAGS == *"OpenBlocks"* ]]
}

@test "list up the dg tags" {
  list_devicegroups_tags
  [[ $DGTAGSLIST == *"OpenBlocks"* ]]
}

@test "update the dg tags" {
  update_devicegroups_tags
  [[ $DGTAGSLIST == *"OpenBlocksNext"* ]]
}

@test "delete devicegroup tag" {
  delete_devicegroups_tags
  echo $DGTAGSLIST
  [[ $DGTAGSLIST == *"[]"* ]]
}

@test "create app tag" {
  create_app_tags
  [[ $APPTAGS == *"plc"* ]]
}

@test "list app tag" {
  list_app_tags
  [[ $APPTAGS == *"plc"* ]]
}

@test "update the app tag" {
  update_app_tags
  [[ $APPLIST == *"humidity"* ]]
}

@test "delete app tag" {
  delete_app_tags
  echo $APPLIST
  [[ $APPLIST == *"[]"* ]]
}
