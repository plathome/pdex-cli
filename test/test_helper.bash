setup() {
  export PDEX_END_POINT_URL=http://localhost:9292/api/v1
  export PDEX_ACCESS_KEY=1234qweasd
  pdex="bin/pdex"
  $pdex c set --url $PDEX_END_POINT_URL --access-key $PDEX_ACCESS_KEY
}

teardown() {
  $pdex u ping
  echo $output
}

fixture() {
  echo "fixure"
}

