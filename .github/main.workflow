workflow "Pipeline" {
  on = "push"
  resolves = ["Docker Tag", "Docker Push"]
}

action "Docker Tag" {
  uses = "actions/docker/tag@master"
  args = "kuafu joway/kuafu"
  secrets = ["GITHUB_TOKEN"]
}

action "Docker Push" {
  needs = ["Docker Tag"]
  uses = "docker://docker:stable"
  args = "push joway/kuafu"
}
