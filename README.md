# unTrace
[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)


unTrace is a tool for professional Penetration Testers who want to exfiltrate data of a system and cover their tracks.
The main focus lies, as the name implies, in covering your tracks.

It has some built in exploits for Linux (all the DirtyC0w variants). They are used for privelidge escalation and may not work depending on your Linux Kernel version.

## We are working on the following add ons:
- Replacing shell scripts with GO functions for portabillity

- Adding easier and more intuitive shell commands

- Adding Enumeration Functions for both Linux and Windows

- More privelidge escalation exploits for both Windows and Linux

- Support for MacOS and Linux (right now only runs on Windows)

- Other exploits to exfiltrate hashes etc..

- persistance settings

## Done so far:
- Added go-prompt functions
- added structure

## Next:
- finish structure
- add cross platform support
- run some commands for testing
- migrate some unTrace functions over(write unTrace package for import)


## Commands

### `unt help`

Shows a list of commands with their description.

### `unt vanish`

(Tries to) remove all traces of your activity on the target system and then disconnects from it.

### `unt escalate`

Shows a list of priviledge escalation exploits and for which kernel/OS Version they are designed for.

### `unt escalate <exploit>`

Runs the exploit defined in `<exploit>`.

### `unt enum`

Collects important information about the target system in a file called report.txt.

### `unt download`

Downloads a file from the server to the client via scp. This command will ask for the filepath.
