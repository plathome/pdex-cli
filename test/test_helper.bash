setup() {
  export PDEX_END_POINT_URL=http://localhost:9292/api/v1
  export PDEX_ACCESS_KEY=1234qweasd
  ew="bin/pdex"
}

teardown() {
  $ew c set --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
  $ew u ping
  echo $output
}

fixture() {
  echo "fixure"
}

