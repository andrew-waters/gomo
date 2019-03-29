workflow "On Push" {
  on = "push"
  resolves = ["build"]
}

action "build" {
  uses = "andrew-waters/actions/go/build@master"
}
