#!/usr/bin/env bash

# git pre-commit hook
#
# To use, store as .git/hooks/pre-commit inside your repository and make sure
# it has execute permissions.
#
# This script does not handle file names that contain spaces.

# List all nonformatted files
files=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

# Some files are not formatted with gofmt. Print message.
nonformatted=$(gofmt -l $files)
if [ "$nonformatted" ]; then
  echo >&2 "Go files must be formatted with gofmt. Running:"
  for fn in $nonformatted; do
    echo >&2 "  gofmt -w $PWD/$fn"
    gofmt -w "$PWD/$fn"
    git add "$PWD/$fn"
  done
  printf "\n"
fi

# Some files are not formatted with golines. Print message.
nonformatted=$(golines -l $files)
if [ "$nonformatted" ]; then
  echo >&2 "Go files must be formatted with golines. Running:"
  for fn in $nonformatted; do
    echo >&2 "  golines -w $PWD/$fn"
    golines -w "$PWD/$fn"
    git add "$PWD/$fn"
  done
  printf "\n"
fi

# Run linter.
make linter || exit 1

# Build binaries to ensure program can be built
make build || exit 1

# Run test.
make test || exit 1

# Clean up unused dependency.
go mod tidy
git add go.mod
git add go.sum

echo ""
echo -e "\e[32mCommitting...\e[0m"
