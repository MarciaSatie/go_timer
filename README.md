# Go Timer ⏱️

A simple **timer application** written in [Go](https://go.dev) using the [Fyne](https://fyne.io) GUI toolkit.  
It runs inside a **Podman container** so that it can be executed on any Linux machine without needing Go or dependencies installed.

---

## � Project Goals
- Learn how to build desktop GUI apps in Go.
- Practice containerization with Podman.
- Show a working, portable example for my portfolio.

---

## �‍� How I Built This
I developed this app while learning Go and Fyne step by step.  
I used **AI (ChatGPT)** as a coding tutor:  
- to explain Go syntax and concepts in “baby steps,”  
- to help debug errors,  
- and to learn how to containerize the app with Podman.  

This way, I am testing the **learning process** guided interactively by AI.  
The goal is to improve my **technical skills** and my **ability to learn new tools independently**.

---

## �️ How to Build

Clone the repo:
```bash
git clone https://github.com/MarciaSatie/go_timer.git
cd go_timer

Build the image:

podman build -t go-timer .

▶️ How to Run

First, allow containers to access your display:

xhost +local:


Then run the app:

./run.sh


Or run directly with:

podman run --rm -it \
  --userns keep-id \
  -e DISPLAY \
  -e XDG_CONFIG_HOME=/tmp \
  -e XDG_CACHE_HOME=/tmp \
  -v /tmp/.X11-unix:/tmp/.X11-unix:ro \
  localhost/go-timer


A window titled "Go Timer" will appear.
Type the number of seconds, press Start Timer, and the countdown will run.

� Files

main.go → the Go source code for the timer.

Dockerfile → instructions to build the containerized app.

run.sh → convenience script to run the app without typing the long Podman command.