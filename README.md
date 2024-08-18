# Portable EV3RT Compiler
A portable [EV3RT](https://ev3rt-git.github.io/) compiler. No need to install all those dependencies for compilation.

## Requirements
You need the following software to be installed:
- Docker
- git

## Setup
Clone this repository:
```bash
$ git clone https://github.com/Binozo/Portable-EV3RT-Compiler
$ cd Portable-EV3RT-Compiler
```
Pull the compiler
```bash
$ docker pull ghcr.io/binozo/portable-ev3-compiler:latest
```
or build it yourself:
```bash
$ docker build -f compiler/Dockerfile -t ghcr.io/binozo/portable-ev3-compiler:latest .
```

### Install EV3RT to the SD card
You can skip this if you already have EV3RT installed on your EV3.

Run the following command to generate the necessary `uImage` bootable system image:
```bash
$ docker run -v $PWD/compiler/build:/src/ev3rt-hrp2/ --rm --env APP=loader --env DIR=base-workspace --name portable-ev3-compiler ghcr.io/binozo/portable-ev3-compiler:latest 
```

Now copy the generated `compiler/build/base-workspace/uImage` in a newly fat32 flashed SD card in the root directory.
Now insert the SD card into your EV3 and try to boot.

## Compile your EV3RT App
As an example, we compile the `linetrace` sample project from the `compiler/build/sdk/workspace` directory.

Run:
```bash
$ docker run -v $PWD/compiler/build:/src/ev3rt-hrp2/ --rm --env APP=linetrace --name portable-ev3-compiler ghcr.io/binozo/portable-ev3-compiler:latest 
```

Now copy the built `compiler/build/sdk/workspace/app` file into the `ev3rt/apps` directory of your SD card. (Create the `apps` directory yourself if missing)

## Running the EV3RT App
[Official EV3RT Docs](https://ev3rt-git.github.io/get_started/#step-5-try-it-out) 
1. Connect the EV3RT running EV3 with your computer via usb. 
2. Place the EV3RT app into the `ev3rt/apps/` folder.
3. Disconnect the usb cable
4. Click `Load App` and `SD Card` on the EV3