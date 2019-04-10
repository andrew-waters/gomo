workflow "On Push" {
  on = "push"
  resolves = ["go test", "codeclimate report"]
}

action "go mod" {
  uses = "andrew-waters/actions/go/mod@master"
}

action "go test" {
  uses = "andrew-waters/actions/go/test@master"
  needs = "go mod"
  secrets = ["CLIENT_ID", "CLIENT_SECRET"]
}

action "codeclimate report" {
  uses = "andrew-waters/actions/go/codeclimate@master"
  needs = "go mod"
  secrets = ["CLIENT_ID", "CLIENT_SECRET", "CC_TEST_REPORTER_ID"]
}
