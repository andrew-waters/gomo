workflow "On Push" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "andrew-waters/actions/go/test@master"
}
