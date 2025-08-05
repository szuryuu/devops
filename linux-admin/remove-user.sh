#!/bin/bash

for i in {1001..1100}; do
    userdel -r user-$i 2>/dev/null
done

groupdel testusers 2>/dev/null
