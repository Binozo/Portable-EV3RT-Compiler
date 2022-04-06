FROM ubuntu:22.04

LABEL maintainer=binozoworks
LABEL org.opencontainers.image.source="https://github.com/Binozo/Portable-EV3RT-Compiler"
LABEL org.opencontainers.image.description="A portable EV3RT compiler"

RUN mkdir app
COPY . app
WORKDIR app

# Update system and install all necessary packages
RUN apt update
RUN apt install wget -y
RUN apt-get install -y bzip2
RUN apt-get -y install git
RUN apt-get -y install make

# Install Compiler
RUN wget http://ev3rt-git.github.io/public/ev3rt-prepare-ubuntu.sh
RUN bash ev3rt-prepare-ubuntu.sh
ENV PATH=$PATH:/opt/gcc-arm-none-eabi-6-2017-q1-update/bin

# Prepare Workspace
RUN git clone https://github.com/ev3rt-git/ev3rt-hrp2
WORKDIR ev3rt-hrp2
RUN git submodule init
RUN git submodule update

# Prepare the configurator
WORKDIR /app/ev3rt-hrp2/cfg
RUN make

# Install the server-side app
WORKDIR /app/ev3rt-hrp2/sdk/
COPY cmd/main/main .

# Change directory
WORKDIR /app/ev3rt-hrp2/sdk/workspace
RUN mkdir project

# Run App
EXPOSE 5321
CMD ./main