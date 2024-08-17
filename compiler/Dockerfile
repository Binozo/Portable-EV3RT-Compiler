FROM ubuntu:24.04

LABEL maintainer=binozoworks
LABEL org.opencontainers.image.source="https://github.com/Binozo/Portable-EV3RT-Compiler"
LABEL org.opencontainers.image.description="A portable EV3RT compiler"

# Compilation options
ENV DIR="sdk/workspace"
ENV APP="helloev3"

WORKDIR /src

# Update system and install all necessary packages
RUN apt-get update
RUN apt-get install wget git bzip2 make -y

# Install Compiler
RUN wget http://ev3rt-git.github.io/public/ev3rt-prepare-ubuntu.sh
RUN bash ev3rt-prepare-ubuntu.sh
ENV PATH=$PATH:/opt/gcc-arm-none-eabi-6-2017-q1-update/bin

# Prepare Workspace
RUN git clone https://github.com/ev3rt-git/ev3rt-hrp2
WORKDIR /src/ev3rt-hrp2
RUN git submodule init
RUN git submodule update

# Prepare the configurator
WORKDIR /src/ev3rt-hrp2/cfg
RUN make

# Install the server-side app
WORKDIR /src/ev3rt-hrp2
CMD cd /src/ev3rt-hrp2/$DIR ; make app=$APP