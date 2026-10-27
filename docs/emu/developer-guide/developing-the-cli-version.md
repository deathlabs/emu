# Environment Setup for Developing the CLI Version

## Install Go

EMU requires Go version 1.23.7 or higher.

**Step 1. Install Go**  
Follow the [official documentation](https://go.dev/doc/install) to install the Go toolchain.

**Step 2. Verify Installation**
```bash
go version
```

Expected output:
```
go version go1.23.7 linux/amd64
```

## Install Golangci-lint

Golangci-lint is a fast linters runner for Go. It runs linters in parallel, uses caching, and supports YAML configuration.

**Step 1. Install golangci-lint**  
```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.5.0

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v2.5.0

# In Alpine Linux (as it does not come with curl by default)
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s v2.5.0
```

**Step 2. Verify Installation**  
```bash
golangci-lint --version
```

Expected output:
```
golangci-lint has version 1.63.4 built with go1.23.4 from c1149695 on 2025-01-03T19:49:42Z
```

**GNU Make**  

We use [GNU Make](https://www.gnu.org/software/make/) to build and manage EMU.

```bash
make --version
```

Expected output:
```
GNU Make 4.3
Built for x86_64-pc-linux-gnu
```

If not installed, run the command sentence below.
```bash
sudo apt-get install build-essential
```

## Clone CLI Repository

```bash
# Create project directory
mkdir -p ~/projects/emu
cd ~/projects/emu

# Clone the CLI repository
git clone https://code.cdso.army.mil/ai2c/emu/cli.git
cd cli
```

## Set Environment Variables

Create a `.env` file in the CLI directory with the required environment variables:

```bash
cat > .env << 'EOF'
export EMASS_API_URL="https://connect.emass.apps.mil"
export EMASS_API_UID="YOUR_ACTUAL_UID"
export EMASS_API_KEY="YOUR_ACTUAL_API_KEY"
export EMASS_SYSTEM_ID="YOUR_SYSTEM_ID"
export NPE_CERTIFICATE=$(cat /path/to/SomeApp.cer)
export NPE_KEY=$(cat /path/to/SomeApp.pem)
EOF

# For Windows builds, add:
echo 'export USER="YOUR_WINDOWS_USERNAME.MIL"' >> .env

# Load environment variables
source .env
```

> **Note:** Replace placeholder values with your actual credentials and paths.

## Building the EMU CLI

After completing all setup steps:

**Step 1.** Navigate to CLI directory.  
```bash
cd ~/projects/emu/cli
```

**Step 2.** Build the EMU binary 
```bash
make
```

This command will:
- Run quality assurance checks (`qa`)
- Execute tests (`tests`)
- Build the binary (`build`)

**Step 3. Verify the build
```bash
./emu -h
```
