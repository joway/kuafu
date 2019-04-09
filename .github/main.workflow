workflow "Pipeline" {
  on = "push"
  resolves = ["Docker Tag"]
}

action "Docker Tag" {
  uses = "actions/docker/tag@master"
  args = "kuafu github/kuafu"
  secrets = ["GITHUB_TOKEN"]
}
