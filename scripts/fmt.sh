find . \
  -type d -name .git -prune -o \
  -type f -name "*.go" -print \
  | xargs gofumpt -l -s -w
