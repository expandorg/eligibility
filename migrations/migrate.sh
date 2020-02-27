#!/bin/sh

action=$1
mnumber=$2
db=$3

if [ "$action" = "composeup" ]; then
  /migrations/migrate \
    -source file:///migrations \
    -database "mysql://$ELIG_DB_USER:$ELIG_DB_PASSWORD@tcp($ELIG_DB_HOST:$ELIG_DB_PORT)/$ELIG_DB" \
    up
  exit 0
fi

if [ "$action" = "build" ]; then
  echo "building migrations"
  exit 0
fi

if [ "$action" != "goto" ] && [ "$action" != "force" ]; then
  echo "operation must be 'goto' or 'force'"
  exit 1
fi

if [ "$mnumber" = "" ]; then
  echo "a migration version is required"
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
