# Database Guide
The aim of the guide is to provide guidelines for the creation of the database and its migration files.

## Migrations

A migration typically consists of two distinct files, one for moving the database to a new state (referred to as "up") and another for reverting the changes made to the previous state (referred to as "down").

The format of those files for SQL are:

```
{sequence_number}_{title}.down.sql
{sequence_number}_{title}.up.sql
```

⚠️<i> Obs: The migrations are executed by the application by [golang-migrate](https://github.com/golang-migrate/migrate) tool</i>