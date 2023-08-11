FROM ghcr.io/goreleaser/goreleaser-cross:v1.20.0

RUN apt-get update -y
RUN apt-get install -y libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev # raylib deps
