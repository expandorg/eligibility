#!/bin/sh

action=$1
mnumber=$2
db=$3

if [ "$action" != "goto" ] && [ "$action" != "force" ] && [ "$action" != "up" ] && [ "$action" != "down" ]; then
  echo "operation must be 'goto' or 'force'"
  exit 1
fi

if [ "$mnumber" = "" ] && [ "$action" != "up" ]; then
  echo "a migration number is required"
  exit 1
fi

if [ "$db" = "" ]; then
  echo "a db address is required"
  exit 1
fi


/migrations/migrate \
    -source file:///migrations \
    -database $db \
    $action $mnumber
