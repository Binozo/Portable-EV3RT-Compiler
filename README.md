# Portable EV3RT Compiler
A portable [EV3RT](https://ev3rt-git.github.io/) compiler. No need to install all those dependencies for compilation.

## Setup
1. Pull the docker image 
```dockerfile
docker pull ghcr.io/binozo/portable-ev3-compiler:latest
```

2. Run the container
```dockerfile
docker run -d -p 5321:5321 --name portable-ev3rt-compiler ghcr.io/binozo/portable-ev3-compiler:latest
```
(The docker-compose file can be found [here](https://github.com/Binozo/Portable-EV3RT-Compiler/blob/master/server/docker-compose.yaml))

3. Download the Client
   - [Windows](https://github.com/Binozo/Portable-EV3RT-Compiler/tree/master/client/bin/main.exe)
   - [Linux](https://github.com/Binozo/Portable-EV3RT-Compiler/tree/master/client/bin/main)
   - Or compile it yourself => `cd client/cmd/main/ && go build`

## Usage
You can use the client in two ways:
1. Place the executable right into your project folder and execute it. \
Example:
Copy the executable file into the [helloev3 example project](https://github.com/ev3rt-git/ev3rt-hrp2-sdk/tree/d33726dd7d8000519ba2c97cb15cff382cff2dab/workspace/helloev3) and execute it. \
The executable file for ev3rt will be in the same directory and will be named "app".

2. Run the executable with two arguments: \
`main.exe <project-path> <output-path>`

## Running the ev3rt app
[Official EV3RT Docs](https://ev3rt-git.github.io/get_started/#step-5-try-it-out) 
1. Connect the EV3RT running EV3 with your computer via usb. 
2. Place the EV3RT app into the `ev3rt/apps/` folder.
3. Disconnect the usb cable
4. Click `Load App` and `SD Card` on the EV3