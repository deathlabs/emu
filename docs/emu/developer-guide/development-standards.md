# Development Standards
This document covers EMU's development standards. This should help developers as they contribute to EMU.

## Repository Structure
When you navigate to the EMU project directory, you should see a structure similar to the following:
```bash
.
├── certs
├── docker
├── docs
├── go.mod
├── src
├── go.mod
├── Makefile
├── README.md
├── .gitignore
└── .gitlab-ci.yaml
```

As a quick overview of the files at the top level of the repository: the `go.mod` file contains the information for the EMU module including the go version and the additional go packages required for running the module, the `Makefile` contains the information for `make` recipies, the `README.md` is the top level EMU documentation file, the `.gitignore` file lists the files that should be ignored when commiting to git (including files like coverage reports, the EMU binary, and the `go.sum` file), and finally the `.gitlab-ci.yml` defines the Gitlab CI/CD pipeline for EMU.

There are also several directories. The `certs` directory should be used to hold NPE certificates for the eMASS API. However, the certificate files **SHOULD NOT** be pushed to the git repository. Please be careful when commiting and ensure you do not push certificates. 

The `docker` directory is for any dockerfiles for EMU.

The `docs` directory is for both user and developer documentation. In this directory you will find the markdown files, but we will also generate documentation in `.pdf` and other formats for easier user digestion.

The final directory is the `src` direcotry. This is where all EMU code lives. See the [Go Project Structure](#go-project-structure) section for more information on how the go source code is structured for EMU.

## Modules
We use the Go language for the EMU project.

Go is an open-source programming language that is extremely good at concurrency and handling distributed systems. If you are unfamiliar with Go, the community provides a pretty expansive collection of introdution documentation and hands-on tutorials to get you started. We highly recommend [Effective Go](https://go.dev/doc/effective_go) and [A Tour of Go](https://go.dev/tour/welcome/1). Documenation and additional tutorials can be found in the Go [Documentation](https://go.dev/doc/)

All code for EMU lives under the `src` directory. Under `src` there are a variety of sub-modules as well as `main.go`. `main.go` is the entry point for the EMU CLI. From there other sub-modules are called to execute the user's CLI command.

```bash
.
└── src
    ├── api
    ├── artifacts
    ├── artifacts_export
    ├── cloud_resource_results
    ├── container_scan_results
    ├── internal
    ├── poams
    ├── static_code_scans
    ├── test_results
    └── main.go
```

Generally, each sub-module represents an endpoint described by the eMASS API with a couple of exceptions. `api` defines a base api client that all other modules use to execute `REST API` requests. `internal` defines additional sub-modules that provide common functionailty across the other modules. This includes things like test helpers and formatting. A standard module structure looks like below.
```bash
.
└── container_scan_results
    ├── cmd.go
    ├── cmd_test.go
    ├── container_scan.go
    └── container_scan_test.go
```

where the `cmd.go` file handles the entrypoint into the module as well as the user exposed functions while `container_scan.go` handles type and struct definitions as well as any internal functions required for the user facing fucntions. Each file then has a `_test.go` file where unit tests are defined. This naming convention follows standard go practices.

## Code
We have a few additional Go standards that we follow for EMU development. We expect all contributors to EMU to follow these standards as well.
- Comments
  - Comments are encouraged throughout code, but should help the reader understand the code. Please do not write extraneous comments.
  - All functions should have a function level comment above the definiton that informs the reader what the function does.
  - All comments should end with a period.
- Variables
  - Variables should be defined at the beginning of each function using the `var <name> <type>` structure.
  - Variable types should be explicitly defined. Do not use the Go short assignment declaration `:=`.
  - When multiple variables are defined in a function, they should be defined in alphabetical order.

Examples of the expected variable definition for EMU are below.
```go
func delete(args []string, writer io.Writer, client *api.Client) error {
  var (
    containerId        string
    containerScan      ContainerScan
    err                error
    flagSet            *flag.FlagSet
    functionHelpString string
    response           *http.Response
    responseBody       []byte
    resultString       string
  )
  ...
}

func Cmd(args []string, writer io.Writer, client *api.Client) error {
	var err error
    ...
}
```

**Reviews**  
All code written for EMU will be reviewed by a senior member of the EMU team. In the EMU project the `main` branch and the next major version branch are locked. This means you cannot commit or make changes directly on the branch. Further, it means that a different individual on the team must review and approve you code before you will be able to merge anything.

If you are unfamiliar with code reviews or just need a refreshers, the standard practices found [here](https://google.github.io/eng-practices/review/reviewer/standard.html) are a good place to start. For a greater overview, see [How to do a code review](https://google.github.io/eng-practices/review/reviewer/). We try to follow these practices within EMU.

- We are trying to ensure the longterm health of this project
  - code should be well designed following best practices and style set forth by style guides and the project
  - complexity is limited
  - code is functional and clear
  - documentation and comments are useful and increase readability
- We will generally look at every line of code submitted
- We will provide lots of feedback
  - please do not take comments personally. Reviewers have a responsibility to ensure code being added improves the project. *The review is of the code, **NOT** you*
  - comments should provide guidance and as needed reasoning for the changes requested
  - we will differentiate between small `nits`, optional changes, or future changes
  - if soemthing is well done, we will also comment so you know
- We try to be quick on review turn arounds to prevent lag and frustration

## Formatting  
We use the Go toolchain's builtin support for source formatting (i.e. `go fmt ...`) to format Go source.

## Linting  
We use [`golangci-lint`](https://github.com/golangci/golangci-lint) as our generic linters runner for Go projects.

## Unit Tests  
We use the Go toolchain's builtin testing support (i.e. `go test ...`) to execute unit tests in Go projects and generate test coverage. From there, we use the Go toolchain's builtin support for coverage report transformation (i.e. `go tool cover ...`) to translate this coverage report to a format that is suitable for consumption in development (by humans) or in CI (by SonarQube).

## Make
We use `make` to manage common command line recipes for EMU. 

GNU `make` is a tool that helps control the generation of executables and other outputs for a project from the source files. It helps ensure consistenty during development and in the resulting builds for EMU. 

`make` commands and the operations they perform are defined in a `Makefile`. For EMU, the `Makefile` can be found in the top level of the EMU repository. We use our `Makefile` to define quality assurance, tests, and builds. We lay out the EMU `Makefile` with clear sections to help make the file easier to read especially for developers new to `make`.

At the top of each section, there will be a line that starts with `.PHONY` such as below.
```bash
.PHONY: qa/format qa/lint qa
```

This defines a list of targets for `make` to run. The targets are phony targets which ensures that make will run the recipe defined regardless of whether or not source files exist as defined in the recipe. But, to simplify, fundamentally this line defines commands that a developer can run via `make`. So, the line above allows developers to run the following in their terminal (within the EMU repository at the same level as the `Makefile`)
```bash
make qa

make qa/format

make qa/lint
```

Following the target name definitions, the actions performed by each target must be provided. These can either be a collection of cli/bash commands or a different target defined elsewhere in the `Makefile` or both.

```bash
qa/lint: prebuild
	@echo "[*] Executing linter on $(prog) source code..."
	golangci-lint run --fix
```

Above, `qa/lint` first runs a different target in the `Makefile` called `prebuild` which will build anything needed to execute the qauality assurance linting target. Then, two commands run which first updates the developer that the linter is about to run and then executes a `Go` linter called `golangci-lint` with a flag to automatically fix linting issues when they occur.

The following commands are common commands for EMU development. Developers are expected to use these commands as they develop on EMU. All `make` commands will be run within the EMU CI pipeline, so it is good practice to verify that all commands function as expected without errors prior to sumbitting merge requests.

**Build**  
The default target for EMU is `build`. This means just running `make` will take the same actions as `make build`. The build target will run qualtiy assurance and test targets as well as build the binary for EMU. 
```bash
make

make build
```

If you just want to recompile the binary without running additional checks use the commands below.
```bash
make build/self
```

**Quality Assurance**  
To check the formatting of the code and ensure it aligns with `Go` best practices, run the command below.
```bash
make qa
```

Code will not be merged until `qa` passes.

**Test**  
To run unit tests against EMU code, run one of the following commands below.
```bash
make test

make dev/test
```

The difference between the two commands is that `test` outputs a coverage report to `coverage.out`. `dev/test` just runs the tests without a coverage report. Both will fail and report errors on a failed test. Code will not be merged until `test` passes.

**Build EMU Container**  
To build a container with EMU that can be run as the base image in a CI/CD pipeline run the following command.
```bash
make build-image
```

Prior to submitting a merge request, you should be able to succesfully run all of the make commands without error. You will not be allowed to merge your code if any of these commands fail. 

## Continuous Integration
We utilize community DevSecOps (cDSO) jobs for our CI pipelines. These pipelines run against all code pushed to the EMU repository and include the following:
- pre-test: quality assurance checks and unit tests
- pre-build-scan: SAST, secrets, and additional linting
- build: builds the container for the project
- post-build-scan: DAST, malware, and entrypoint scans
- clean: performs clean up of the temporarily created resources
- review: verifies results of all previous stages and scans

In total there are currently 14 jobs that run in the EMU pipeline. No additional efforts are required on the part of the developer to run the pipelines. Because of how the repository is set up as well as the inclusion of the `.gitlab-ci.yml` file and its contents in the EMU repository, the core pipeline described above is triggered upon all commits to the repository (each time a developer pushes a local commit using `git push`).

In order for the code to be merged, the pipeline must pass. This means that at a minimum all require jobs that are not allowed to fail complete sucessfully. There are some jobs that are allowed to fail. When these jobs fail, instead of a red `x` and a complete halt of the pipeline, you instead get a yellow warning symbol and the pipeline will continue.

It is best practice for all jobs including jobs allowed to fail to pass for a pipeline. Depending on the failure, you reviewer may not approve your merge request until the job passes.

**Pipeline Debugging**  
Pipelines run everytime code is pushed. They generally take a few minutes to run as resources must be spun up to support and run the jobs. Once a pipeline is complete, you can use the Gitlab UI to explore the pipeline and jobs. This is usually the first place to start for determining why a pipeline failed. To find pipelines, jobs, and the artifacts jobs produce (which is common for scans), navigate to the repository and click on the build tab on the sidebar. From there you can select which page to go to. 

Determine which job(s) failed and navigate to those jobs. You should get an output from the job that should help indicate the point of failure. This may be something like a command failing or a scan not meeting a threshold. From there, you can determine what needs to change in your code to make the job pass.

Update your local branch with those changes, commit, and push. This will be the iterative cycle until all jobs pass.

