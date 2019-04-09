workflow "Pipeline" {
  on = "push"
  resolves = ["Docker Tag"]
}

action "Docker Tag" {
  uses = "actions/docker/tag@8cdf801b322af5f369e00d85e9cf3a7122f49108"
  args = "kuafu github/kuafu"
  secrets = ["GITHUB_TOKEN"]
}
