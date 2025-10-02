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


Build the container image:

podman build -t go-timer .


Run the app (Linux Mint, Ubuntu, Fedora, etc.):

xhost +local:
./run.sh

� Windows (native build, no container needed)

Install Go → https://go.dev/dl/

Clone the repo (PowerShell):

git clone https://github.com/MarciaSatie/go_timer.git
cd go_timer


Fetch dependencies:

go mod tidy


Build and run:

go build -o go-timer.exe .
.\go-timer.exe


This will open the Go Timer window directly as a native Windows application �

� Files

main.go → Go source code for the timer.

Dockerfile → Container build instructions for Linux.

run.sh → Convenience script to run in Podman (Linux).