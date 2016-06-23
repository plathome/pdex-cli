#!/bin/bash

#: Version      : 0.0.1

usage="$(basename "$0") [-h] -- Simple tool to send messages using curl

usage:
    -k  key
    -h  Show this help text
    -d  device id
    -m  Message to send
"

BASE_URL="http://localhost:9292/api/v1/"

die () {
    echo >&2 "$@"
    exit 1
}

while getopts 'k:d:m:h' opt; do
    case "$opt" in
        h)  die "$usage"
            ;;
        m)  MESSAGE=$OPTARG
            ;;
        d)  DEVICE=$OPTARG
            ;;
        k)  KEY=$OPTARG
            ;;
        \?) printf "illegal option: -%s\n" "$OPTARG" >&2
            die "$usage"
            ;;
    esac
done
shift $((OPTIND - 1))

if [[ -z $DEVICE  ]]; then
    echo "$usage"
    die "please provide the device-id."
fi
if [[ -z $MESSAGE ]]; then
    echo "$usage"
    die "please input the message."
fi
if [[ -z $KEY ]]; then
    echo "$usage"
    die "please input the key."
fi

#curl -w "\n" -d "key=e63ab20b0771" -d "message=01.8529e6.26387394" -X POST http://localhost:9292/api/v1/utils/hmac
HMAC_URL=$BASE_URL/utils/hmac

hmac=$(curl \
-d "key=$KEY" \
-d "message=$MESSAGE" \
-X POST $HMAC_URL)

echo $hmac

#curl -w "\n" -H "Authorization: Bearer Co8lc_LkNrs26woftVQeEnRHwzs4PUCTF5MuAQ==" -H "Content-Type: application/json" -d '{"a":1}' -X POST  http://localhost:9292/api/v1/devices/01.8529e6.26387394/message

URL=$BASE_URL/devices/$DEVICE/message

curl \
-H  "Authorization: Bearer Co8lc_LkNrs26woftVQeEnRHwzs4PUCTF5MuAQ==" \
-H "Content-Type: application/json" \
-X POST --data '
    {
        "message": "'"$MESSAGE"'"
    }' $URL
echo

