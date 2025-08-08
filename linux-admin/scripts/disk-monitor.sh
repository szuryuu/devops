#!/bin/bash

# URL="https://eo5wtn983pri732.m.pipedream.net/request-bin-recive-disk"

count=$(df -h | awk 'NR>1 && int($5) >= 90' | wc -l)

echo "Total high usage disks: $count"

df -h | awk 'NR>1 && int($5) >= 90' | while read -r FS SIZE USED UNUSED PERCENT MOUNT
do
    curl -X POST $URL \
    -H "Content-Type: application/json" \
    -d '{
        "Filesystem": "'$FS'",
        "Total Size": "'$SIZE'",
        "Used Size": "'$USED' '$PERCENT'",
        "Free Size": "'$UNUSED'",
        "Mounted on": "'$MOUNT'"
    }'

    echo "$FS, $SIZE, $USED, $UNUSED, $PERCENT, $MOUNT"
done
