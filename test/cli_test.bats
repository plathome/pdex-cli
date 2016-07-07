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
  run $pdex cr devices --deid-prefix 01.817eb9
  [[ $output == *"deid"* ]]
}

@test "create devicegroup" {
  run $pdex cr dg
  [[ $output == *"deid_prefix"* ]]
}

@test "create channels" {
  run $pdex cr ch --deid 01.817eb9.bdd15c7e --app-id b38154ecfe5043af905858e33595a6fe
  [[ $output == *"channel_id"* ]]
}

# Listing API

@test "list apps" {
  run $pdex ls apps
  echo $output | grep "\"count\""
}

@test "list devicegroups" {
  run $pdex ls dg
  echo $output | grep "\"count\""
}

@test "list devices" {
  run $pdex ls de --deid-prefix 01.bbac01
  echo $output | grep "\"count\""
}

@test "list channels" {
  run $pdex ls ch --deid 01.817eb9.c1d6c837
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
  run $pdex send msg --deid 01.817eb9.bdd15c7e "message send for test"
  echo $output | grep "\"transaction_id\""
}

@test "read messages" {
  run $pdex read msg --app-id b38154ecfe5043af905858e33595a6fe
  echo $output | grep "\"count\""
}

@test "send command" {
  run $pdex send cmd --channel-id 10e0be9636e64317a04d8157342d1e54 --app-id b38154ecfe5043af905858e33595a6fe "test command sending"
  echo $output | grep "\"transaction_id\""
}

@test "read commands" {
  run $pdex read cmd --deid 01.817eb9.bdd15c7e
  echo $output | grep "\"count\""
}
