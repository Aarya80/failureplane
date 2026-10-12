# FailurePlane — repository setup and first working session

This file is for **GitHub repo/local implementation**. A different set of
files named `chatgpt_project_sources/` belongs in the **new ChatGPT Project**.

## 0. What has already been done in this archive
The Go in-memory execution state machine was validated with normal tests,
race-enabled tests, vet, and its CLI demo. This repo was NOT pushed to GitHub.
PostgreSQL migrations, sandbox, HTTP API, multi-worker networking and replay
are planned, not validated.

## 1. Windows setup (recommended: WSL2 Ubuntu)
Open PowerShell **as Administrator** if WSL is not installed:

```powershell
wsl --install -d Ubuntu
```

Reboot if Windows requests it. Open Ubuntu, set up your Linux account, then:

```bash
sudo apt update
sudo apt install -y git build-essential make unzip ca-certificates
```

Install the current supported Go version from https://go.dev/doc/install.
Check `go version` is 1.23 or newer; do not depend on Ubuntu apt's potentially
older Go package. Install VS Code + WSL extension if you use VS Code.

**You do not need Docker or Rust for the first coding session.** Add Docker
Desktop with WSL integration when implementing PostgreSQL. Add Rust/rustup
and `rustup target add wasm32-wasip2` for the isolated runner later.

## 2. Extract repo
Unzip `FailurePlane_GitHub_Repo_Starter.zip` in your Windows Downloads folder
or another location, then copy the extracted `github_repo` contents into
`~/code/failureplane` inside Ubuntu (avoid running Linux-heavy repos directly
from `/mnt/c` where file IO can be slower):

```bash
mkdir -p ~/code/failureplane
cd ~/code/failureplane
# Copy the extracted github_repo/ files here (include hidden files).
```

A portable alternative from Windows Explorer is to copy the `github_repo`
folder into `\\wsl$\Ubuntu\home\<your-ubuntu-user>\code\`, then rename it to
`failureplane` if needed.

Verify:

```bash
pwd
ls -la
go version
git --version
```

## 3. First coding session: verify starting point
From the actual repository root (the folder containing `go.mod`):

```bash
go test ./...
go test -race ./...
go vet ./...
go run ./cmd/failureplane demo
python3 examples/faults/ambiguous_effect_spike.py
```

The demo must reject `Old worker completion` with a stale-generation error,
then accept `New worker completion` at generation 2.

Read:
1. `docs/guarantees.md`
2. `internal/execution/state.go`
3. `internal/execution/state_test.go`
4. `docs/adr/0001-postgres-coordination.md`

## 4. Create GitHub repo (once, manually)
On GitHub select **New repository**, name it `failureplane`, choose public or
private, and **DO NOT pre-add a README, license or .gitignore** (all are in the
starter). Public is reasonable when you're comfortable with the starter.

Back in Ubuntu:

```bash
git init
git branch -M main
git add .
git commit -m "chore: initialize FailurePlane engineering starter"
git remote add origin https://github.com/Aarya80/failureplane.git
git push -u origin main
```

You must authenticate GitHub in your own terminal. If the repository already
has a `main` branch or your GitHub username differs, adjust accordingly.
Never force-push over existing work.

## 5. Create the first 15 GitHub issues
The issue drafts are in `issues/`. You can create them one-by-one with the
GitHub UI, or after installing and authenticating GitHub CLI (`gh auth login`):

```bash
# Preview the issue drafts first.
ls issues/
# OPTIONAL: creates actual issues; run just once, after review.
bash scripts/publish-issues.sh
```

The issue publishing script is intentionally guarded and warns before it
creates anything on GitHub. You can skip it and make issues manually.

## 6. Actual implementation order

**First:** understand and edit the pure ownership state machine, add or refine
tests and ADRs. **Second:** set up local PostgreSQL and migrations. **Third:**
implement atomic claim, fenced result commit and a two-worker SQL race test.
**Only then:** build the network worker and bounded runner.

Read `docs/roadmap.md` and `issues/01-*` through `issues/05-*`.

## 7. PostgreSQL milestone (not needed on day 1)
Once Docker Desktop is installed and WSL integration enabled:

```bash
docker version
docker compose up -d postgres
docker compose ps
```

The initial schema is in `migrations/0001_init.up.sql`. Its schema and
migration apply/rollback test harness are issue #3. Do not pretend schema
acceptance has occurred merely because the SQL file exists.

## 8. Work with ChatGPT in the new project
Upload the curated source files there; use its custom instructions and paste
the `FIRST_CHAT_PROMPT.md` supplied with the ChatGPT Project pack.
At every session report what changed and what was actually tested. Keep
`docs/STATUS.md` current and re-upload it when ChatGPT's project sources lag.
