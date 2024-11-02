# Anyver

Anything Version Manager - A version manager to temporarily replace any program with any script on your PATH.

## What does Anyver solve?

### The Problem

If your developer machine is like mine, it's filled with small CLI tools and utilities that are installed directly on the machine with a package manager like Homebrew.

Examples of such tools in my experience can be jq, yq, buf, protoc or mockery.

When working in (or for) a company on a project, they might use some of the same set of tools and utilities, but different (typically older) versions.

The dilemma is then; Do you adapt your precious developer machine on a per-company or per-project basis? How do you support having different versions of the same programs depending on your workload?

Perhaps the most obvious solution is to run the tools inside a Docker image, but if you have to use a custom bash function or alias to run a secondary version of a program, it might not work with a project's checked-in scripts in Makefiles or package.json (that rely on executing the programs as they are installed by default), without having to change the scripts to use your executable names which then conflicts with other developers or build servers using the same scripts.

### The Solution

Anyver is a small utility that aims to remedy this problem by managing a set of alternative scripts for any given program name. When a script is activated by Anyver, an executable (bash) script with the same name as the program will be put in `~/.anyver/apps`. The content of this script file will be whatever script content you have configured in a global Anyver yaml config file.

For example, if I want to create a spoofed version of `jq`, I can run `anyver use jq the-other-version`. If the script `the-other-version` has the content `docker exec --workdir $PWD my-toolbox-image $@`, then that will be the content of the file `~/.anyver/apps/jq` created by Anyver.

Since I've made sure that `~/.anyver/apps` is always first on my `PATH` environment variable, I know that Anyver's spoofed version of the executable will "win" when a `jq` command is executed in a shell.

To restore `jq` to its original executable (`/usr/local/bin/jq` in my case), I'll run `anyver restore jq` and Anyver will remove the file `~/.anyver/apps/jq` so that the next (default) occurrance of jq on my `PATH` will be found and executed.

Therefore, Anyver can replace any program with any script and serve as a version manager for anything.

## Installation

### 1. Install the Anyver exe

Either install it to your `$GOPATH/bin` and make sure `$GOPATH/bin` is on your `PATH`:

```
go install github.com/eaardal/anyver
```

Or clone the repo, build it and make sure the built executable is somewhere on your `PATH`:

```
go build -o anyver main.go && cp <REPO_DIR> <DIR_ON_PATH>
```

### 2. Create the inital Anyver config file

In a terminal, run:

```
anyver init
```

This will create a default `~/.anyver/config.yaml` file.

### 3. Ensure `~/.anyver/apps` is _first_ on your `PATH`

Add this to your `.zshrc`, `.bashrc` or other shell config script:

```shell
export PATH=$HOME/.anyver/apps:$PATH
```

The path is not required to be first-first, but it should be in your `PATH` before wherever your programs are usually found, such as `/usr/local/bin` and `/usr/bin`.

You can inspect your `PATH` with: `echo $PATH`.

## Usage

```
NAME:
   Anything Version Manager

USAGE:
   Anything Version Manager [global options] command [command options] [arguments...]

COMMANDS:
   config       Print the entire Anyver config file as-is
   use          Set the given version as the active command when running the app
   restore      Removes the Anyver shortcut which should make the system's default executable the main one
   restore-all  Removes all active Anyver shortcuts which should make the system's default executables the main ones for all apps
   context      Manage multiple apps in a context
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

### `config`

Prints the file `~/.anyver/config.yaml` as it's found on disk.

Example:

```shell
> anyver config

active:
  jq: _system
  mockery: my-company
apps:
  jq:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image jq $@
  mockery:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image mockery $@
contexts:
  default:
    jq: _system
    mockery: _system
  all-apps-my-company:
    jq: my-company
    mockery: my-company
```

### `use`

Activates an app alias. Anyver will find the script for the given app name and make a executable file with that app name in `~/.anyver/apps`.

Example:

Given the config file

```yaml
active:
  jq: _system
  mockery: my-company
apps:
  jq:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image jq $@
  mockery:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image mockery $@
contexts:
  default:
    jq: _system
    mockery: _system
  all-apps-my-company:
    jq: my-company
    mockery: my-company
```

Alias and verify an app:

```shell
# No aliased apps before we start
> ls ~/.anyver/apps
(no files found)

# Using system/default jq
> which jq
/usr/local/bin/jq

# Tell Anyver to make an alias for jq
> anyver use jq my-company
Now using version "my-company" for app "jq"

# Now there is 1 aliased app
> ls ~/.anyver/apps
jq

# Content of the file is our script
> cat ~/.anyver/apps/jq
docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image jq $@

# Using aliased jq from Anyver
> which jq
/Users/<name>/.anyver/apps/jq
```

### `restore`

Removes the Anyver shortcut which should make the system's default executable the main one.

Will simply delete the file under `~/.anyver/apps` if it exists.

Example:

```
> anyver restore jq
Restored app "jq"
```

### `restore-all`

Removes all active Anyver shortcuts which should make the system's default executables the main ones for all apps.

Same as `restore`, but it deletes all files under `~/.anyver/apps` instead of just one.

### `context use`

Alias all apps listed under the given context

Context lets you define multiple apps under one name. Useful if there are multiple apps you have to switch between when working on different projects.

Example:

Given the config file

```yaml
active:
  jq: _system
  mockery: my-company
apps:
  jq:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image jq $@
  mockery:
    my-company: docker run --rm -it --volume $PWD:$PWD --workdir $PWD my-toolbox-image mockery $@
contexts:
  default:
    jq: _system
    mockery: _system
  all-apps-my-company:
    jq: my-company
    mockery: my-company
```

Applying a context

```
> anyver context use all-apps-my-company
Now using version "my-company" for app "jq"
Now using version "my-company" for app "mockery"
```

### `context restore`

Same as restore or restore-all but only deletes apps in the given context.
