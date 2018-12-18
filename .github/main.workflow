workflow "Test on Push" {
  on = "push"
  resolves = ["Test coverage"]
}

action "Run tests" {
  uses = "./.github/actions/gotest"
}

action "Test coverage" {
  needs = ["Run tests"]
  uses = "./.github/actions/codecov"
  args = "-f ${GITHUB_WORKSPACE}/coverage.txt"
}
