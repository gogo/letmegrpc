#!/bin/bash
set -xe
git-anchor ./src/github.com/gogo/letmegrpc/deps.json > ./src/github.com/gogo/letmegrpc/deps.sh && chmod +x ./src/github.com/gogo/letmegrpc/deps.sh
