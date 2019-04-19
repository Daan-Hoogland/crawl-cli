# Crawl CLI
[![Go Report Card](https://goreportcard.com/badge/github.com/Daan-Hoogland/crawl-cli)](go-report-card)  ![Latest Release Shield](https://img.shields.io/github/release/daan-hoogland/crawl-cli.svg  "latest-release") ![License](https://img.shields.io/github/license/daan-hoogland/crawl-cli.svg  "license")

**This application is a work in progress.**
It's mostly just a tool for me to work on while learning Go. Any functional suggestions, tips or pull requests are more than welcome.

This tool is designed to scan a file system for files or running processes matching certain parameters. 

## Getting started
The application can be compiled from the source or downloaded from the [release page](https://github.com/Daan-Hoogland/crawl-cli/releases). The latest release can be found in the shield at the top. 

There ~~are~~ will be prebuilt binaries available for:
- Linux (64 bit)
- OSX (64 bit)
- Windows (32 bit and 64 bit)

### Building
Building the project from source will require a valid go installation. Instructions for your OS will be widely available on the web.

Included in the repository is a `makefile`. This file includes all actions that need to be taken to build the application. Afterwards, the generated binary can be found in `/build/{os}`. 

### Installing
#### Linux & OSX
Start off by executing `echo $PATH` in a terminal of your choice and verify `/usr/local/bin` is in your path, if not instructions on how to add it can be found [here](https://unix.stackexchange.com/questions/26047/how-to-correctly-add-a-path-to-path) . 

If the directory is part of your `PATH`, the executable can be copied to `/usr/local/bin`. Linux: `cp build/linux/crawl /usr/local/bin`, OSX: `cp build/osx/crawl /usr/local/bin`. After this, the application should be available from any terminal of choice. 

If when executed from the terminal the following command shows: 
`-bash: crawl: command not found`
This means that the application doesn't have sufficient permissions to be executed. These can be added by executing `chmod +x /usr/local/bin/crawl`

##### Arch Linux
- [ ] add application to AUR.

#### Windows 32/64bit
Move the `crawl-cli` executable to a location of your choice. The executable itself can be found in `build/windows/32/` or `build/windows/64/` depending on which version your Windows supports.

In order to access the `crawl` application from `cmd.exe` or PowerShell, the directory which you have chosen needs to be added to your user's path variable. How to do this depends on your version of Windows, but there are plenty of [examples](https://stackoverflow.com/a/9546345)  available online to show you how to.


### Usage
There are 4 general actions. These can be accessed by typing `crawl-cli <action> <flags>`. The  actions and their flags are briefly described below.

#### help
Shows the list of actions along with their flags. Can be executed on an action by adding the `-h` flag. Example: `crawl help -h` will output the following:
```
Help provides help for any command in the application.
Simply type crawl help [path to command] for full details.

Usage:
  crawl help [command] [flags]

Flags:
  -h, --help   help for help

Global Flags:
  -D, --debug        turn debug mode on or off
      --dev          turn trace mode on or off
  -l, --log string   file that the log will be written to
  -v, --verbose      verbose level of the logger

```

#### connect
The connect subcommand is used to test the connection to a webserver. The webserver is generally ran on a server that all the clients can connect to in order to not lose the output when executed on multiple clients.

```
Usage:
  crawl connect [flags]

Flags:
  -h, --help            help for connect
  -p, --port int        port used to connect to web service (default 9000)
  -t, --target string   external ipv4 address pointing to web service (default "127.0.0.1")

Global Flags:
  -D, --debug        turn debug mode on or off
      --dev          turn trace mode on or off
  -l, --log string   file that the log will be written to
  -v, --verbose      verbose level of the logger
```
Any errors/success messages will be shown in the output of the command.

#### analyse
`analyse` executes the application without connecting to a server. This means that the results of the application **will not** be saved. Instead, they will be shown in the terminal as the output of the command, in a sorted table. 

```
The analyse command scans the filesystem for files or running processes.
Unlike the scan command, the analyse command does not send the results
to a running web application.

Usage:
  crawl analyse [flags]

Flags:
  -a, --algorithm string   the hash algorithm used
  -d, --directory string   directory the application starts in
  -H, --hash string        hash of the file to search for
  -h, --help               help for analyse
  -n, --name strings       name(s) of files to search for
  -s, --size int           file size that the target file must match
  -T, --threads int        number of processes to spawn. (default 2)

Global Flags:
  -D, --debug        turn debug mode on or off
      --dev          turn trace mode on or off
  -l, --log string   file that the log will be written to
  -v, --verbose      verbose level of the logger
```

#### scan
Does the same as `analyse`, except the results are posted to a running web service configured by flags available to the `scan` command. The output of the command will not be shown in the console, but can instead be viewed in the interface of the web service the application posted to.

```
The scan commands scans the filesystem for files or running applications
and compares them to the given input. Any matches will be sent
to the connected web application.

Usage:
  crawl scan [flags]

Flags:
  -a, --algorithm string   the hash algorithm used
  -d, --directory string   directory the application starts in
  -H, --hash string        hash of the file to search for
  -h, --help               help for scan
  -n, --name strings       name(s) of files to search for
  -p, --port int           port used to connect to web service (default 9000)
  -s, --size int           file size that the target file must match
  -t, --target string      external ipv4 address pointing to web service (default "127.0.0.1")
  -T, --threads int        number of processes to spawn. (default 2)

Global Flags:
  -D, --debug        turn debug mode on or off
      --dev          turn trace mode on or off
  -l, --log string   file that the log will be written to
  -v, --verbose      verbose level of the logger
```

## Todo
- [ ] create `makefile`
- [ ] unit test fs operations
- [ ] configure jenkins to automatically build and test
- [ ] allow scanning of running processes
- [ ] start working on web service
- [ ] start working on web interface

## Used libraries
- [Logrus](https://github.com/sirupsen/logrus) 
	- [Nested Logrus Formatter](https://github.com/antonfisher/nested-logrus-formatter)  
- [Cobra](https://github.com/spf13/cobra) 
- [Spinner](github.com/briandowns/spinner) 
- [walk](https://github.com/MichaelTJones/walk) 
