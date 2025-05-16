#!/bin/bash

# Load the environment variables from the .env file
source .env

# Run Liquibase through Docker Compose
docker-compose up liquibase
