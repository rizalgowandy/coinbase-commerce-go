#!/usr/bin/env bash

echo -e "\e[32mRunning:\e[33m setup.\e[0m\n"

echo -e "\e[32mInstalling:\e[33m gin for live reload.\e[0m"
command -v gin 2>/dev/null || GO111MODULE=off go get -v github.com/codegangsta/gin
echo ""

echo -e "\e[32mInstalling:\e[33m air for live reload.\e[0m"
command -v air 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin # v1.12.1
echo ""

echo -e "\e[32mInstalling:\e[33m mockgen for mock generator.\e[0m"
command -v mockgen 2>/dev/null || GO111MODULE=off go get -v github.com/golang/mock/mockgen # v1.4.4
echo ""

echo -e "\e[32mInstalling:\e[33m golangci-lint for linter.\e[0m"
command -v golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin v1.33.0
echo ""

echo -e "\e[32mInstalling:\e[33m wire for compile time dependency injection.\e[0m"
command -v wire 2>/dev/null || GO111MODULE=off go get -v github.com/google/wire/cmd/wire # v0.4.0
echo ""

echo -e "\e[32mInstalling:\e[33m swag for open api documentation.\e[0m"
command -v swag 2>/dev/null || GO111MODULE=off go get -v github.com/swaggo/swag/cmd/swag # v1.6.7
echo ""

echo -e "\e[32mInstalling:\e[33m golines for formatting long lines code.\e[0m"
command -v golines 2>/dev/null || GO111MODULE=off go get -v github.com/segmentio/golines
echo ""

echo -e "\e[32mSetup:\e[33m pre-commit hook.\e[0m"
file=.git/hooks/pre-commit
cp scripts/pre-commit.sh $file
chmod +x $file
test -f $file && echo "$file exists."
echo ""

echo -e "\e[32mSetup:\e[33m success.\e[0m"
