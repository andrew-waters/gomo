workflow "On Push" {
  on = "push"
  resolves = ["go test"]
}

action "go test" {
  uses = "andrew-waters/actions/go/test@master"
}
