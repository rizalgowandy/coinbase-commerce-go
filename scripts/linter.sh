#!/usr/bin/env bash

echo -e "\e[32mRunning: \e[33mlinter.\e[0m"

command time -f %E golangci-lint run -c "$PWD"/.golangci.yaml || exit 1

echo -e "\e[32mLinter: \e[33msuccess.\e[0m"
