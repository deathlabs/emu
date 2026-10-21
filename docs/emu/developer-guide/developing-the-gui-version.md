# Environment Setup for Developing the GUI Version

## Install Wails

**Step 1.** Install Wails CLI. Install the Wails framework for building desktop applications:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Fix PATH if "wails command not found".
```bash
# Check Go path
go env GOPATH

# Add to PATH
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

**Verify installation:**
```bash
wails doctor
```

**Step 2.** Install System Dependencies.

```bash
# Install basic dependencies
sudo apt install build-essential pkg-config libgtk-3-dev

# Fix Ubuntu 24.04 webkit issue
sudo add-apt-repository "deb http://gb.archive.ubuntu.com/ubuntu jammy main"
sudo apt update
sudo apt install libwebkit2gtk-4.0-dev
```

**Verify dependencies.**  
```bash
wails doctor
```
All required dependencies should show "Installed".

**Note:** Docker and NSIS are optional dependencies (marked with *)

**Step 3.** Install Node.js v22

```bash
node --version
npm --version
```

If Node.js is not v22 or higher, install NVM.
```bash
# Install NVM
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash

# Load NVM
source ~/.bashrc

# Install Node.js v22
nvm install 22
nvm use 22

# Verify
node --version  # Should show v22.x.x
npm --version   # Should show 10.x.x or higher
```

**Step 4.** Setup Project.

Navigate to project directory.
```bash
cd /src/emu/desktop/
```

Install Go dependencies.
```bash
go mod download
go mod tidy
```

**Step 5.** Create .netrc File.

Create authentication file for private repositories:

```bash
nano ~/.netrc
```

Add your GitLab credentials:
```
machine code.cdso.army.mil
login <your_gitlab_username>
password <your_personal_access_token>
```

Save and secure the file:
```bash
chmod 600 ~/.netrc
```

**Step 6.** Run Development Mode.

Start development server.
```bash
wails dev
```

What happens:
1. Downloads Go dependencies
2. Installs frontend dependencies  
3. Builds frontend assets
4. Generates Go bindings
5. Compiles application
6. Opens desktop application window
7. Starts development server at `http://localhost:5174/`

Success indicators:
- Desktop application window opens with EMU dashboard
- Browser window may also open (this is normal for debugging)
- Terminal shows "Development mode exited"

**Step 7.** Build Production Version.

Build for Linux
```bash
wails build
```

Build for Windows (cross-compilation)
```bash
# Install Windows cross-compilation tools
sudo apt install gcc-mingw-w64

# Build Windows executable
wails build -platform windows/amd64
```

Executables location:
* Linux: `./build/bin/emu` 
* Windows: `./build/bin/emu-desktop.exe`

**Step 8.** Run Production Build.

**Linux:**
```bash
./build/bin/emu
```

**Windows:**
Transfer `emu-desktop.exe` to Windows machine and double-click to run.

## Troubleshooting

**Wails doctor shows missing dependencies:**
- Re-run the dependency installation commands
- Verify all packages installed successfully

**"Package not found" errors:**
- Ensure you've added the Jammy repository for webkit
- Run `sudo apt update` after adding repositories

**Build fails:**
- Check internet connection for dependency downloads
- Verify .netrc file if using private GitLab repositories
- Ensure all dependencies show "Installed" in `wails doctor`

**Cross-compilation fails:**
- Install mingw-w64: `sudo apt install gcc-mingw-w64`
- Or build directly on target platform

## Development Workflow

**For ongoing development:**
1. Make code changes
2. `wails dev` provides hot reload for frontend changes
3. Backend changes require restart
4. Test with `wails dev`
5. Build production with `wails build`

**Key directories:**
- `frontend/` - React/TypeScript frontend code
- `*.go` files - Go backend code  
- `build/bin/` - Compiled executables
- `wails.json` - Wails configuration
