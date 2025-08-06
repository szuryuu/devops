#!/bin/bash

URL="https://eo5wtn983pri732.m.pipedream.net"

groupadd testusers
for i in {1001..1100}; do
    USERNAME="user-$i"

    useradd -u $i -g testusers $USERNAME
    GROUPNAME=$(id -g -n $USERNAME)

    curl -X POST $URL \
    -H "Content-Type: application/json" \
    -d '{
        "username": "'$USERNAME'",
        "uid": '$i',
        "group": "'$GROUPNAME'",
        "gid": '$(getent group testusers | cut -d: -f3)'
    }'

    echo "User $USERNAME created"
    echo $GROUPNAME
    sleep 0.05
done
