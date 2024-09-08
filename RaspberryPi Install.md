# Install Notes for Raspberry Pi

Tested on the following plaforms.
Raspberry Pi 4
Linux raspberrypi 6.6.31+rpt-rpi-v8 #1 SMP PREEMPT Debian 1:6.6.31-1+rpt1 (2024-05-29) aarch64 GNU/Linux


# Turn off overclock

vi /boot/firmware/config.txt
arm_boost=0

# Increase swap on Raspberry pi

sudo dphys-swapfile swapoff

sudo vi /etc/dphys-swapfile

CONF_SWAPSIZE=1024

sudo dphys-swapfile setup

sudo dphys-swapfile swapon

sudo reboot

# Install GIT on Raspberry Pi

sudo apt update
sudo apt install git


# Clone the dmxlights repo

mkdir -p project/src/github.com/dhowlett99
git clone https://github.com/dhowlett99/dmxlights.git

## First make sure you have the latest version of golang 
cd ~/Downloads
wget https://go.dev/dl/go1.23.1.linux-arm64.tar.gz
sudo tar -C /usr/local -xzvf go1.23.1.linux-arm64.tar.gz 
go version


# Edit .profile

export PATH=$PATH:/usr/local/go/bin
GO111MODULE=on
GOPATH=/Users/<USER>/project

$ ./.profile
go version go1.23.1 linux/arm64


# Install prerequisites 

sudo apt-get install libusb-1.0-0-dev libasound2-dev portaudio19-dev libgl1-mesa-dev xorg-dev

cd project/src/github.com/dhowlett99/dmxlights
make build

./dmxlights


Configure ALSA 

vi /usr/share/alsa/alsa.conf

comment out unused interfaces.

# Install Visual Studio Code 

sudo apt update
sudo apt install code










