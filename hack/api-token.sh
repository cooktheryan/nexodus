#!/bin/sh

set -e

USERNAME="$1"
PASSWORD="$2"

token=$(curl -s -f -X POST \
    https://auth.try.nexodus.local/realms/nexodus/protocol/openid-connect/token \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -d "username=$USERNAME" \
    -d "password=$PASSWORD" \
    -d "scope=openid profile email" \
    -d "grant_type=password" \
    -d "client_id=nexodus-cli"
)

echo $token
