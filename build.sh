#!/bin/bash

# Get the full go repo url
REPO=$(pwd |  rev | cut -d'/' -f-3 | rev)

# Get the name of the app
APP="${PWD##*/}"

# Get this tag as the version
VERSION=$(git describe --abbrev=0 --tags)

# Compile
xgo \
  -go 1.10.x \
  --targets=darwin-10.11/amd64 \
  --ldflags "-X main.version=${VERSION}" \
  ../$APP

rm kombustion
mv kombustion-darwin-10.11-amd64 kombustion