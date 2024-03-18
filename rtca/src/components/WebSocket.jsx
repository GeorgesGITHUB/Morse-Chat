import React, { useState, useEffect } from 'react';

const WebSocketComponent = () => {
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    //debugger
    // Establish WebSocket connection when component mounts
    const newSocket = new WebSocket('ws://localhost:8080/ws');

    // Event listener for connection open
    newSocket.onopen = () => {
      console.log('WebSocket connection established');
    };

    // Event listener for received messages
    newSocket.onmessage = (event) => {
      console.log('Message received:', event.data);
    };

    // Event listener for connection close
    newSocket.onclose = () => {
      console.log('WebSocket connection closed');
    };

    // Save the WebSocket instance to state
    setSocket(newSocket);

    // Clean up function to close WebSocket when component unmounts
    return () => {
      // Check if WebSocket instance exists before closing
      if (socket && socket.readyState === WebSocket.OPEN) {
        socket.close();
        }
    };
  }, []); // Run effect only once when component mounts

  const sendMessage = () => {
    if (socket) {
      // Example: Sending a message to the server
      socket.send(JSON.stringify(
        { 
            Sender: 'Person A',
            Content: 'Hello, how are you'
        }
      ));
    }
  };

  return (
    <div>
      <button onClick={sendMessage}>Send Message</button>
    </div>
  );
};

export default WebSocketComponent;
