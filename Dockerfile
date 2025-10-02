# Use official Go image for build
FROM docker.io/library/golang:1.22-bookworm

# Install dependencies for building and running Fyne apps
RUN apt-get update && apt-get install -y \
    libgl1 \
    libglx-mesa0 \
    libgl1-mesa-dri \
    libx11-6 \
    libxkbcommon0 \
    libxcursor1 \
    libxi6 \
    libxrandr2 \
    libxinerama1 \
    libxxf86vm1 \
    pkg-config \
    build-essential \
    libgl1-mesa-dev \
    xorg-dev \
    libxkbcommon-dev \
    && rm -rf /var/lib/apt/lists/*

# Create a normal user inside the container (avoid permission issues)
RUN useradd -m appuser

# Set working directory for the app
WORKDIR /home/appuser/app

# Copy source code into the container
COPY . .

# Build the app binary
RUN go build -o /home/appuser/go-timer .

# Switch to the new user (so Fyne can write to ~/.config, ~/.cache)
USER appuser
WORKDIR /home/appuser

# Run the app by default
CMD ["./go-timer"]
