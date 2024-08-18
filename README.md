# Portable EV3RT Compiler
A portable [EV3RT](https://ev3rt-git.github.io/) compiler. No need to install all those dependencies for compilation.

Supports the `TOPPERS/HRP2` and `TOPPERS/HRP3` kernel.

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
Pull the compiler: (Please choose one compiler, either the `hrp2` or `hrp3` one)
```bash
$ docker pull ghcr.io/binozo/portable-ev3rt-compiler:hrp2
$ docker pull ghcr.io/binozo/portable-ev3rt-compiler:hrp3
```
or build it yourself:
```bash
$ docker build -f compiler/Dockerfile -t ghcr.io/binozo/portable-ev3rt-compiler:hrp2 .
$ docker build -f compiler/Dockerfile -t ghcr.io/binozo/portable-ev3rt-compiler:hrp3 .
```

### Install EV3RT to the SD card
You can skip this if you already have EV3RT installed on your EV3.

Run the following command to generate the necessary `uImage` bootable system image: (Choose one)

#### TOPPERS/HRP2
```bash
$ git clone https://github.com/ev3rt-git/ev3rt-hrp2 compiler/build
$ cd compiler/build && git submodule init && git submodule update && cd ../..
$ docker run -v $PWD/compiler/build:/src/ ghcr.io/binozo/portable-ev3rt-compiler:hrp2 /bin/bash -c "cd cfg;make"
$ docker run -v $PWD/compiler/build:/src/ --rm --env APP=loader --env DIR=base-workspace --name portable-ev3rt-compiler ghcr.io/binozo/portable-ev3rt-compiler:hrp2
```
Your `uImage` will be in `compiler/build/base-workspace/`.

#### TOPPERS/HRP3
```bash
$ git clone https://github.com/ev3rt-git/ev3rt-hrp3 compiler/build
$ cd compiler/build && git submodule init && git submodule update && cd ../..
$ docker run -v $PWD/compiler/build:/src/ --rm --env IMG=loader --env DIR=sdk/firmware --name portable-ev3rt-compiler ghcr.io/binozo/portable-ev3rt-compiler:hrp3
```
Your `uImage` will be in `compiler/build/sdk/firmware/`.

#### Now continue:
Now copy the generated `uImage` in a newly fat32 flashed SD card in the root directory.
Now insert the SD card into your EV3 and try to boot.

## Compile your EV3RT App
As an example, we compile the `helloev3` sample project from the `compiler/build/sdk/workspace` directory.

Compile with HRP2:
```bash
$ docker run -v $PWD/compiler/build:/src/ --rm --env APP=helloev3 --name portable-ev3rt-compiler ghcr.io/binozo/portable-ev3rt-compiler:hrp2
```
Or compile with HRP3:
```bash
$ docker run -v $PWD/compiler/build:/src/ --rm --env APP=helloev3 --name portable-ev3rt-compiler ghcr.io/binozo/portable-ev3rt-compiler:hrp3
```

Now copy the built `compiler/build/sdk/workspace/app` file into the `ev3rt/apps` directory of your SD card. (Create the `apps` directory yourself if missing)
Or proceed with this way:

## Running the EV3RT App
[Official EV3RT Docs](https://ev3rt-git.github.io/get_started/#step-5-try-it-out) 
1. Connect the EV3RT running EV3 with your computer via usb. 
2. Place the EV3RT app into the `ev3rt/apps/` folder.
3. Disconnect the usb cable
4. Click `Load App` and `SD Card` on the EV3