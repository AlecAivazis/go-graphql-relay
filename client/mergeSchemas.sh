#!/bin/bash

# Check if at least one argument (directory) is provided
if [ $# -lt 1 ]; then
    echo "Usage: $0 <directory>"
    exit 1
fi

# Set the directory as the first argument
directory="$1"

# Check if the directory exists
if [ ! -d "$directory" ]; then
    echo "Directory '$directory' does not exist."
    exit 1
fi

# Initialize an empty variable to store the concatenated content
concatenated_content=""

# Iterate over all files in the directory
for file in "$directory"/*; do
    # Check if the item is a file (not a directory)
    if [ -f "$file" ]; then
        # Concatenate the file's content to the concatenated_content variable
        concatenated_content+="$(cat "$file")"
    fi
done

# Print the concatenated content (you can also redirect it to a file)
echo "$concatenated_content" > merged-schema.graphql
