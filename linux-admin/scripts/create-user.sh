#!/bin/bash

groupadd testusers
for i in {1001..1100}; do
    USERNAME="user-$i"

    useradd -u $i -g testusers $USERNAME
    GROUPNAME=$(id -g -n $USERNAME)

    curl -X POST https://eo50x3rnudzgrdw.m.pipedream.net \
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
