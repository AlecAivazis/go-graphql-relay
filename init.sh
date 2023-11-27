#!/bin/bash

# check if the user provided a new name
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <new_project_name>"
    exit 1
fi

new_project_name="$1"
search_string="projectname"

# walk through the project and replace the projectname with the new name
find . -type f -exec sed -i "s/$search_string/$new_project_name/g" {} \;

    mv "db/projectname.go" "db/$new_project_name.go"
fi
