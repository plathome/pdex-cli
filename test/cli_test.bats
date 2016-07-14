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
  run $pdex cr apps --app-name-suffix "test-app-1"
  [[ $output == *"app_id"* ]]
}

@test "create device" {
  run $pdex cr dg | jq .deid_prefix
  run $pdex cr devices --deid-prefix $output
  [[ $output == *"deid"* ]]
}

@test "create devicegroup" {
  run $pdex cr dg
  [[ $output == *"deid_prefix"* ]]
}

@test "create channels" {
  create_channel
  [[ $CH == *"channel_id"* ]]
}

# Listing API

@test "list apps" {
  list_apps
  [[ $LIST -ge 1 ]]
}

@test "list devicegroups" {
  list_dgs
  [[ $LIST -ge 1 ]]
}

@test "list devices" {
  list_devices
  [[ $LIST -eq 1 ]]
}

@test "list channels" {
  list_channel
  [[ $LIST -eq 1 ]]
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

@test "read messages" {
  read_message
  [[ $READ -eq 1 ]]
}

@test "read single message" {
  read_single_message
  [[ $READ == "message send for test" ]]
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
