#!/bin/bash

# Get the latest tag
latest_tag=$(git describe --tags $(git rev-list --tags --max-count=1))
IFS='.' read -r major minor patch <<< "$latest_tag"

# Increment the patch version
patch=$((patch + 1))

# Create the new tag
new_tag="$major.$minor.$patch"
git tag "$new_tag"
git push origin "$new_tag"

echo "Updated tag to $new_tag"