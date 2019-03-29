workflow "On Push" {
  on = "push"
  resolves = ["go test", "codeclimate report"]
}

action "go test" {
  uses = "andrew-waters/actions/go/test@master"
  secrets = ["CLIENT_ID", "CLIENT_SECRET"]
}

action "codeclimate report" {
  uses = "andrew-waters/actions/go/codeclimate@master"
  secrets = ["CLIENT_ID", "CLIENT_SECRET", "CC_TEST_REPORTER_ID"]
}
