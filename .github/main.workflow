workflow "Pipeline" {
  on = "push"
  resolves = ["Docker Build", "Docker Tag", "Docker Push"]
}

action "Docker Build" {
  uses = "actions/docker/cli@master"
  args = ["build", "-t", "kuafu", "."]
}

action "Docker Tag" {
  needs = ["Docker Build"]
  uses = "actions/docker/tag@master"
  args = "kuafu joway/kuafu"
  secrets = ["GITHUB_TOKEN"]
}

action "Docker Push" {
  needs = ["Docker Tag"]
  uses = "docker://docker:stable"
  args = "push joway/kuafu"
}
