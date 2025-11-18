#!/bin/bash

if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set"
  exit 1
fi

data=$(curl -s https://platform.zone01.gr/assets/superhero/all.json)

# Extract relatives and replace real newlines with literal \n
relatives=$(echo "$data" | jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives" | sed ':a;N;$!ba;s/\n/\\n/g')

echo "$relatives"
