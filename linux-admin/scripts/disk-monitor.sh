#!/bin/bash

URL="https://eo5wtn983pri732.m.pipedream.net"

count=$(df -h | awk 'NR>1 && int($5) >= 20' | wc -l)

echo "Total high usage disks: $count"

df -h | awk 'NR>1 && int($5) >= 20' | while read -r FS SIZE USED UNUSED PERCENT MOUNT
do
    curl -X POST $URL \
    -H "Content-Type: application/json" \
    -d '{
        "Total Size": "'$SIZE'",
        "Used Size": "'$USED' '$PERCENT'",
        "Free Size": "'$UNUSED'",
        "Mounted on": "'$MOUNT'"
    }'

    echo "$SIZE, $USED, $UNUSED, $PERCENT, $MOUNT"
done
