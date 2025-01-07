#!/bin/bash

# Determine OS type
OS="$(uname)"
ZIP_URL="https://raw.githubusercontent.com/MJAZ93/wizzy/main/build/mac.zip"
ZIP_FILE="wizzy.zip"
EXTRACTED_FOLDER="."

rm .wizzy

# Function to exit script on error
exit_on_error() {
    echo "$1"
    exit 1
}

# Download the GitHub zip
echo "Downloading repository zip..."
curl -L -o "$ZIP_FILE" "$ZIP_URL" || exit_on_error "Download failed"

# Check if the zip file is valid
if [[ ! -f "$ZIP_FILE" || $(file --mime-type -b "$ZIP_FILE") != "application/zip" ]]; then
    exit_on_error "Downloaded file is not a valid zip archive"
fi

# Extract the zip to a temporary folder
echo "Extracting zip..."
mkdir -p "$EXTRACTED_FOLDER"
unzip -o "$ZIP_FILE" -d "$EXTRACTED_FOLDER" || exit_on_error "Unzip failed"

# Move the contents to the current folder
echo "Moving files to the root folder..."
shopt -s dotglob  # Include hidden files
mv mac/* . || exit_on_error "Failed to move files"
shopt -u dotglob  # Turn off dotglob after moving

# Go back and clean up the zip file
echo "Cleaning up..."
rm -rf mac.zip __MACOSX
