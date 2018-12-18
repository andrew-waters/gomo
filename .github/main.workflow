workflow "Test on Push" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "go"
  args = "go test ./..."
}
