load test_helper

@test "setup" {
  run $pdex c set --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
  [ "$output" = "successfully configured" ]
}

@test "profile" {
  run $pdex c profile --name new --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
  [ "$output" = "successfully configured" ]
}

@test "ping" {
  run $pdex u ping
  [ "$output" = "{\"result\":\"pong\"}" ]
}

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


