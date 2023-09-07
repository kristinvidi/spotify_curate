#!/bin/bash

# Load the environment variables from the .env file
source .env

# Run the Liquibase update command
liquibase --changeLogFile=${CHANGELOG_FILE} --url=jdbc:postgresql://localhost:5432/${POSTGRES_DB} --username=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} update


# Rollback!!
# liquibase --changeLogFile=${CHANGELOG_FILE} --url=jdbc:postgresql://localhost:5432/${POSTGRES_DB} --username=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} rollbackToDate "2023-08-01"