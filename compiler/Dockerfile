FROM debian:trixie-slim

# Build options
ARG HRP_KERNEL_VERSION="2"

LABEL maintainer=Binozoworks
LABEL org.opencontainers.image.source="https://github.com/Binozo/Portable-EV3RT-Compiler"
LABEL org.opencontainers.image.description="A portable EV3RT TOPPERS/HRP${HRP_KERNEL_VERSION} Kernel compiler"

# Compilation options
ENV DIR="sdk/workspace"
ENV APP=""
ENV IMG=""

WORKDIR /src

# Update system and install all necessary packages
RUN apt-get update
RUN apt-get install wget git bzip2 make ruby -y

# Install Compiler
RUN wget http://ev3rt-git.github.io/public/ev3rt-prepare-ubuntu.sh
RUN bash ev3rt-prepare-ubuntu.sh
RUN rm ev3rt-prepare-ubuntu.sh
ENV PATH=$PATH:/opt/gcc-arm-none-eabi-6-2017-q1-update/bin

# Prepare Workspace
ENV rtVersion="ev3rt-hrp${HRP_KERNEL_VERSION}"
RUN git clone https://github.com/ev3rt-git/${rtVersion} .
RUN git submodule init
RUN git submodule update

# Prepare the configurator
WORKDIR /src/cfg
RUN cd cfg && file Mafefile && make || true
# To make HRP2 and HRP3 support possible in this single Dockerfile we have to check if HRP2's cfg directory exist and then run make on it.
# HRP3 does not have that directory
#RUN make

# Install the server-side app
WORKDIR /src/
# Added img=$IMG for HRP3 support
CMD cd /src/$DIR ; make app=$APP img=$IMG