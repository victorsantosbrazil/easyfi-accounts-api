#!/bin/bash

# MySQL connection parameters for the remote server
DB_HOST="localhost"
DB_USER="root"
DB_PASSWORD="root"
DB_DATABASE="institution"
DB_TABLE="institution"

# Define the SQL query to create the table
CREATE_TABLE_QUERY="USE $DB_DATABASE; CREATE TABLE IF NOT EXISTS $DB_TABLE  (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);"

# Execute the SQL query to create the table on the remote server
docker run --network host --rm mysql:8.2.0 mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD --protocol tcp -e "$CREATE_TABLE_QUERY"

# Check if the execution was successful
if [ $? -eq 0 ]; then
    echo "Table '$DB_TABLE' created successfully on the remote server."
else
    echo "Error creating table '$DB_TABLE': $?"
    exit 1
fi

exit 0