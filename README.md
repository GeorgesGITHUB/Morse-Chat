# Morse Chat (Work in Progress)
A Real Time Chat App with a morse code twist! 

Bidirectional communication between clients handled by a Go server, using Web Sockets. 

`Client1 <-ws-> Server <-ws-> Client2`

Dynamically renders content of the React component as Plain Text or Morse Code, using states.

Uses the Web Audio API (AudioContext, Oscillator, GainNode) to play Morse Code

# How to run it
1. Start the Go server in <code>/rtca/backend/</code> with <code>go run .</code>
2. Start the React frontend in <code>/rtca/</code> with <code>npm run dev</code>

# Screenshot(s)
Client apps communicating
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/1f6c377d-eedd-4fcf-b92e-a5e71ce6c434)


Run time logs
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/3233dc6d-6c61-4b62-a13e-0c950e793b1e)
