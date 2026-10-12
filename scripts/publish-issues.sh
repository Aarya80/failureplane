#!/usr/bin/env bash
# Optional helper; creates remote GitHub issues after explicit terminal confirmation.
# Run once, from repo root, after `gh auth login` and remote setup.
set -euo pipefail

command -v gh >/dev/null || { echo 'Install GitHub CLI first: https://cli.github.com/'; exit 1; }
gh auth status >/dev/null
[ -d issues ] || { echo 'Run from repository root'; exit 1; }

printf 'Repository: '; gh repo view --json nameWithOwner --jq .nameWithOwner
printf '\nThis will CREATE 15 GitHub issues. Existing issues are not checked.\n'
read -r -p 'Type CREATE to continue: ' confirm
[ "$confirm" = 'CREATE' ] || { echo 'Cancelled'; exit 0; }

for file in $(find issues -maxdepth 1 -name '[0-9][0-9]-*.md' | sort); do
  title=$(sed -n '1s/^# //p' "$file")
  if [ -z "$title" ]; then
    echo "Cannot extract title: $file" >&2
    exit 1
  fi
  echo "Creating: $title"
  gh issue create --title "$title" --body-file "$file"
done
