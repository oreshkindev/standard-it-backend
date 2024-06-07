#!/bin/bash

. ./env.sh

# Migrate the database
# To install migrate cli read this post https://www.freecodecamp.org/news/database-migration-golang-migrate/
#
# Usage:
#   migrate.sh [-up] [-down] [-drop] [-create <name>]
#
# Options:
#   -up
#       Move the database up to the latest version.
#
#   -down
#       Move the database down to the previous version.
#
#   -drop
#       Drop the database and all of its tables.
#
#   -create <name>
#       Create a new migration file with the given name.
#
# Arguments:
#   None

while [[ "$#" -gt 0 ]]; do
  case "$1" in
    -up)
      # Move the database up to the latest version.
      migrate -path migrations -database "${DATABASE_URL}" up
      shift
      ;;
    -down)
      # Move the database down to the previous version.
      yes | migrate -path migrations -database "${DATABASE_URL}" down
      shift
      ;;
    -drop)
      # Drop the database and all of its tables.
      yes | migrate -path migrations -database "${DATABASE_URL}" drop
      shift
      ;;
    -create)
      # Create a new migration file with the given name.
      if [ -n "$2" ]; then
        migrate create -ext sql -dir migrations "$2"
        shift 2
      else
        echo "Error: Missing migration name." >&2
        exit 1
      fi
      ;;
    *)
      # Invalid option
      echo "Invalid option: $1" >&2
      exit 1
      ;;
  esac
done
