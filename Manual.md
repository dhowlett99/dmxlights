# DMX Lights

## Introduction

DMX lights is a Open Source Project to control DMX light fixtures.  It is designed for the small to medium mobile disco rig and be very easy to setup and use.

Firstly we use the Novation Launchpad Mini (Mk3) as the control surface and a cheap FTDI RS422 controller as the interface for the DMX lights.

Second design goal was to do away with having to have a large fixture library,  with DMX lights you simply add the description of what channels
do what to a simple text file.

## Instalation

clone the DMX lights repositary

git clone github.com/dhowlett99/dmxlights

set up the modules

```sh
export GO111MODULE=on
go mod init
go mod tidy

go build dmxlights.go
```

### Setting up your fixtures

The simplest of fixture definitions is shown below:-

fixtures.yaml

```yaml
fixtures:
- name: FOH PAR
  description: RGB PAR
  type: rgb
  address: 1
  group: 1
  channels:
  - number: 1
    name: Red1
  - number: 2
    name: Green1
  - number: 3
    name: Blue1
```

Fixture Definition
| Field |  Function |
|-|-|
| name | The name is arbitrary and is only used as a label.|
| description | The description is arbitrary and only used to record extra info on the fixture.|
| type | The type defines the sequence type.  Valid values are rgb, scanner, switch.|
| address | The address is the DMX start address you have programed your fixture at.|
| group | The group defines which sequence this fixture belongs too.|
| channels | The channels is the list of the fixtures available DMX traits, these have a number and a name. See below.|

Channel Definition

| Field |  Function |
|-|-|
| Number  |  The DMX channel number
| Name    |  Channel Names, the name is used to indicate channel funtion.
| Value     | Used to set a static value when the name is set to 'Static.
| MaxDegrees | Tells dmxlights what the fixture is capable of.
| Offset     | Offset allows you to position the fixture.
| Comment  | This is is arbitrarty text field.
| Settings  | Further details

### Settings example

You can specifty further details for a channel. This examples outlines the different functions for channel 3 named shutter.

``` yaml
- number: 3
    name: Shutter
    settings:
      - name: Open
        setting: 255
      - name: Closed
        setting: 0
      - name: Strobe
        setting: 100
```

## Channel Example

 Channel Names, the name is used to decide which number in the Sequence and what color it will be. Valid names are

RGB Fixture - Valid Channel Names or Keywords are :-
| Field |  Function |
|-|-|
|Red1 | Channel respondes to fixture 1 Red
|Green1 | Channel respondes to fixture 1 Green |
|Blue1 | Channel respondes to fixture 1 Blue
|Red2 | Channel respondes to fixture 2 Red
|Green2 | Channel respondes to fixture 2 Green
|Blue2 | Channel respondes to fixture 2 Blue|
|Red3 | Channel respondes to fixture 3 Red|
|Green3 | Channel respondes to fixture 3 Green|
|Blue3 | Channel respondes to fixture 3 Blue|
|Red4 | Channel respondes to fixture 4 Red|
|Green4 | Channel respondes to fixture 4 Green|
|Blue4 | Channel respondes to fixture 4 Blue|
|Red5 | Channel respondes to fixture 5 Red|
|Green5 | Channel respondes to fixture 5 Green|
|Blue5 | Channel respondes to fixture 5 Blue|
|Red6 | Channel respondes to fixture 6 Red|
|Green6 | Channel respondes to fixture 6 Green|
|Blue6 | Channel respondes to fixture 6 Blue|
|Red7 | Channel respondes to fixture 7 Red|
|Green7 | Channel respondes to fixture 7 Green|
|Blue7 | Channel respondes to fixture 7 Blue|
|Red8 | Channel respondes to fixture 8 Red|
|Green8 | Channel respondes to fixture 8 Green|
|Blue8 | Channel respondes to fixture 8 Blue|
|Dimmer|Channel respondes to Master Brightness ( alternative to Master above|
|Static| Channel is used to set static DMX value. See settings example above.

Scanner Fixture - Valid Channel Names or Keywords are :-
| Field |  Function |
|-|-|
|Shutter | Channel respondes Shutter Size |
|Pan| Channel respondes Scanner Pan|
|FinePan| Channel respondes Scanner Fine Pan - Not currently Used|
|Tilt| Channel respondes Scanner Tilt|
|FineTilt| Channel respondes Scanner Fine Tilt - Not currently Used|
|Gobo| Channel respondes to Gobo Selection|
|Color| Channel respondes to Color Selection|
|Master |Channel respondes to Master Brightness |
|Dimmer|Channel respondes to Master Brightness ( alternative to Master above|

## Running DMX lights

Plug the FTDI interface card and Novation Lauchpad using their respective USB cables.

```sh
./dmxlights
```

## LaunchPad Layout

The launchpad buttons are laid out in a simple manner, the very top row are global controls.
The next four top rows are reserved to control and display the sequence as defined in the fixtures.yaml file.
The bottom three rows are reserved as storage for your scenes. Once you have selected all the required sequence
charateristics you can press the SAVE buttton and the one of these preset buttons to store the scene.

![LaunchPad Layout](Layout.png)

The very bottom row of buttons give more controls but these are specific to the selected sequence.

The buttons on the far right allow you to select a sequence, save presets, start a sequence, stop a sequence
The botton far right is the blackout button.

## Sequences

A sequence is the basic control set in DMX Lights. A sequence can have a few different modes depending on
the sequence type.

## Chase Sequence

A basic chase sequence of 8 fixtures with 8 different colors.

## Scanner Sequence

A specific to a scanner, this type of sequence can scan in a circle, left to right, up and down and finally in saw tooth motion.

## Static Colors

A static color sequence is where you want to setup a set of uplighters with specific colors.

## Switch Sequence

A switch sequence is simply eight switches that can be used to control simple devices like projectors.
A swicth can have multiple states, for example you could set a fixture to have specific color, brightness or Gobo.
