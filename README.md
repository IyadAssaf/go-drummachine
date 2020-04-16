# go-drummachine

A drum machine written in Go making use of fluidsynth

## Installation

Install `fluidsynth` beefore building the project:
- osx: `brew install fluidsynth`
- ubuntu: `apt-get install fluidsynth`

Run `make install` to build and install the program and supporting files 

## Usage

```
$ drummachine --kit 808

  #######    #####    #######
 ##     ##  ##   ##  ##     ##
 ##     ## ##     ## ##     ##
  #######  ##     ##  #######
 ##     ## ##     ## ##     ##
 ##     ##  ##   ##  ##     ##
  #######    #####    #######

┌──────────────────────────────────────────────────────────────────────────────────────────┐
│BassDrum (b)│ Snare (s)  │  Clap (c)  │HHClosed (h)│ HHOpen (j) │TomHigh (t) │ TomLow (y) │
│──────────────────────────────────────────────────────────────────────────────────────────│
│            │            │            │            │            │            │            │
└──────────────────────────────────────────────────────────────────────────────────────────┘
```
