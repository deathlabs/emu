# Prerequisites

## Create Personal Access Token

To clone from a GitLab repository, you will need to have developer access and create a personal access token.

**Step 1.** Log into GitLab and click on your user icon, then click `Edit profile` on the dropdown.

**Step 2.** In user settings, click on `Access Tokens` in the sidebar.

**Step 3.** On the `Personal Access Tokens` page, click the `Add new token` button.

**Step 4.** Configure your token:
- Give the token a descriptive name
- Set the expiration date (maximum 6 months by policy)
- Select scopes: `write_repository` (minimum), `api` and `write_registry` (recommended)

**Step 5.** Click `Create`.

> **Important:** You will only see the token ONCE. Copy it immediately and store it securely (e.g., in a password manager or secure text file).

## Configure Git Authentication

The .netrc file allows Git to authenticate with private GitLab repositories automatically.

**Step 1.** Create .netrc file in your home directory:
```bash
cat > ~/.netrc << EOF
machine code.cdso.army.mil
login YOUR_GITLAB_USERNAME
password YOUR_PERSONAL_ACCESS_TOKEN
EOF
```

**Step 2.** Set secure permissions (required):
```bash
chmod 600 ~/.netrc
```

**Step 3.** Verify permissions:
```bash
ls -la ~/.netrc
# Should show: -rw-------
```

> **Note:** Replace YOUR_GITLAB_USERNAME and YOUR_PERSONAL_ACCESS_TOKEN with your actual credentials from the previous step.

## Configure Go for Private Modules

These settings are required for both CLI and Desktop development:

```bash
# Mark cDSO modules as private
go env -w GOPRIVATE=code.cdso.army.mil

# Skip checksum verification for private modules
go env -w GONOSUMDB=code.cdso.army.mil

# Fetch private modules directly
go env -w GONOPROXY=code.cdso.army.mil

# Use direct proxy
go env -w GOPROXY=direct

# Set Go binary directory
go env -w GOBIN=~/go/bin

# Verify settings
go env | grep PRIVATE
```
