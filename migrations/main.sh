#!/bin/sh

action=$1
mnumber=$2

if [ "$action" != "goto" ] && [ "$action" != "force" ] && [ "$action" != "up" ]; then
  echo "operation must be 'goto' or 'force'"
  exit 1
fi

if [ "$mnumber" = "" ] && [ "$action" != "up" ]; then
  echo "a migration number is required"
  exit 1
fi

/migrations/migrate \
    -source file:///migrations \
    -database "mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" \
    $action $mnumber
