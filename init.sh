#!/bin/sh
kill $(pgrep mynozom)
./mynozom > /dev/null 2>&1 & 
disown
echo "running on prccedd id :" + $(pgrep mynozom)

