# Show current Git branch (only if in a repo)
parse_git_branch() {
  git rev-parse --is-inside-work-tree &>/dev/null || return

  branch=$(git symbolic-ref --short HEAD 2>/dev/null || git describe --tags --exact-match 2>/dev/null)
  status=$(git status --porcelain 2>/dev/null)

  dirty=""
  [[ -n "$status" ]] && dirty="*"

  echo " (${branch}${dirty})"
}
