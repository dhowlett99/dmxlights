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

wget https://go.dev/dl/go1.23.1.linux-arm64.tar.gz
sudo tar -C /usr/local -xzvf go1.23.1.linux-arm64.tar.gz 
go version
go version go1.23.1 linux/arm64

# Marks Notes 

wget https://go.dev/dl/go1.23.1.linux-arm64.tar.gz
tar xzf go1.23.1.linux-arm64.tar.gz
export PATH=$PATH:$PWD/go/bin
sudo apt-get install libusb-1.0-0-dev libasound2-dev portaudio19-dev libgl1-mesa-dev xorg-dev
git clone https://github.com/dhowlett99/dmxlights
cd dmxlights/
go version
make 
./dmxlights


## Install PortAudio

 sudo apt-get install libusb-1.0-0-dev libasound2-dev portaudio19-dev libgl1-mesa-dev  xorg-dev

# Install Lib USB

sudo apt-get install libusb-1.0-0-dev

sudo apt-get install libx11-dev

sudo apt install libxcursor-dev

sudo apt install libxrandr-dev

sudo apt-get install libx11-dev libxinerama-dev

sudo apt-get install libxi-dev

# Environment Variables.
export GOARCH=arm
export GOHOSTARCH=arm
export LDFLAGS="-L /usr/lib/aarch64-linux-gnu"

# Clean up default ALSA configuration file.







