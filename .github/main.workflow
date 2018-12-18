workflow "Test on Push" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "docker://go"
  args = "test ./..."
}
