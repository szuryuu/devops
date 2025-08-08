#!/bin/bash

URL="https://eo5wtn983pri732.m.pipedream.net/request-bin-recive-testusers"

if ! getent group testusers &>/dev/null; then
    groupadd testusers
fi

for i in {1001..1100}; do
    USERNAME="user-$i"

    if id $USERNAME &>/dev/null; then
        echo "User $USERNAME already exists"
        continue
    fi

    useradd -u $i -g testusers $USERNAME
    GROUPNAME=$(id -g -n $USERNAME)
    GID=$(id -g "$USERNAME")

    curl -X POST $URL \
    -H "Content-Type: application/json" \
    -d '{
        "username": "'$USERNAME'",
        "uid": '$i',
        "group": "'$GROUPNAME'",
        "gid": '$GID'
    }'

    echo "User $USERNAME created"
    sleep 0.05
done
