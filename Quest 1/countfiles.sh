#!/bin/bash

# Count all regular files and directories recursively
count=$(find . -type f -o -type d | wc -l)

# Print only the number
echo $count
