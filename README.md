# Morse Chat (See "How to run" section below)
A Real Time Chat App with a morse code twist!

Communication between **React** clients, managed by a **Go HTTP server**, using **concurrency**, and connected to clients via **Web Sockets**.

Frontend and Backend are **Docker** containers built and served using a **multi-stage process**, with **ports** exposed to each other.

**Plays Morse Code** audio using the **Web Audio API** (AudioContext, Oscillator, GainNode).

Saves conversation messages and user info to an **AWS RDS PostgreSQL** database.

Dynamically (**stateful**) renders content of the React component as Plain Text or Morse Code.


# Screenshot(s)
Client apps communicating
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/1f6c377d-eedd-4fcf-b92e-a5e71ce6c434)


Run time logs
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/3233dc6d-6c61-4b62-a13e-0c950e793b1e)

# Prerequisites to run
- Have `git` installed
- Have `Docker` installed

# How to run
1. Open a terminal instance and download the repo by pasting `git clone https://github.com/GeorgesGITHUB/Morse-Chat/`
2. Change your current directory to the project's root folder, using `cd Morse-Chat`
3. Create and run the Docker containers by pasting and running `docker-compose up`
4. Have many chat clients by opening multiple browser tabs of `http://localhost:5173/`
5. For each client, select a different Profile Preset by clicking on one of the side bar buttons
6. Message away!
