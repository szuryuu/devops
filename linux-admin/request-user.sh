#!/bin/bash

groupadd testusers
for i in {1001..1100}; do
    useradd -u $i -g testusers user-$i
    curl -d '{
      \"username\": \"user-$i\",
      "uid": $i,
      "group": "testusers",
      "gid": $(getent group testusers | cut -d: -f3)
    }'   -H "Content-Type: application/json"   https://eo50x3rnudzgrdw.m.pipedream.net
done
