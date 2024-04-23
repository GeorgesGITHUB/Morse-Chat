# Morse Chat (Work in Progress)
A Real Time Chat App with a morse code twist! 

Bidirectional communication between React clients handled by a Go server, using Web Sockets. 

`Client1 <-ws-> Server <-ws-> Client2`

Uses the Web Audio API (AudioContext, Oscillator, GainNode) to play Morse Code

Saves conversation messages and user info to an AWS RDS PostgreSQL database.

Dynamically renders content of the React component as Plain Text or Morse Code, using states.


# Screenshot(s)
Client apps communicating
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/1f6c377d-eedd-4fcf-b92e-a5e71ce6c434)


Run time logs
![image](https://github.com/GeorgesGITHUB/Morse-Chat/assets/31967906/3233dc6d-6c61-4b62-a13e-0c950e793b1e)
