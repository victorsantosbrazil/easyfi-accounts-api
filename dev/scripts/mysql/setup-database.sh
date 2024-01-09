#!/bin/bash

# Usage: ./script.sh -h <DB_HOST> -u <DB_USER> -p <DB_PASSWORD> -d <DB_NAME> -s <SQL_DIR>

# MySQL connection parameters for the remote server
DB_HOST="localhost"
DB_USER="root"
DB_PASSWORD="root"
DB_NAME="institution"

# Parse command-line arguments
while getopts ":h:u:p:d:s:" opt; do
  case $opt in
    h) DB_HOST="$OPTARG" ;;
    u) DB_USER="$OPTARG" ;;
    p) DB_PASSWORD="$OPTARG" ;;
    d) DB_NAME="$OPTARG" ;;
    s) SQL_DIR="$OPTARG" ;;
    \?) echo "Invalid option: -$OPTARG" >&2; exit 1 ;;
    :)  echo "Option -$OPTARG requires an argument." >&2; exit 1 ;;
  esac
done

BASEDIR=$(dirname $0)
SQL_DIR="$BASEDIR/../../../database/schema"

# Execute the SQL query to create the table on the remote server
RUN_SCRIPT_COMMAND_PREFIX="mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD --protocol tcp $DB_NAME"

#Iterate over SQL files in the directory
for SQL_FILE in "$SQL_DIR"/*.sql; do
    # Check if the file exists and is a regular file
    if [ -f "$SQL_FILE" ]; then
        # Execute the SQL file
        echo "Executing $SQL_FILE ..."
        if $RUN_SCRIPT_COMMAND_PREFIX < "$SQL_FILE"; then
            echo "$SQL_FILE executed successfully"
        else
            echo "Error executing $SQL_FILE"
            exit 1
        fi
    else
        echo "Skipping $SQL_FILE - not a regular file."
    fi
done

exit 0