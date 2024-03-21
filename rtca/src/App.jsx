import { useEffect, useState } from 'react'
import './App.css'
import useWebSocket from 'react-use-websocket'

function App() {
  const WS_URL = 'ws://localhost:8080/ws'
  const {
    sendJsonMessage, 
    lastJsonMessage, 
    readyState
  } = useWebSocket(WS_URL, { share: true })

  const [user, setUser] = useState('')
  const [msg, setMsg] = useState('')
  const [msgHistory, setMsgHistory] = useState([])

  useEffect(() => {
    if (lastJsonMessage !== null) {
      setMsgHistory( (prevHistory) => [...prevHistory, lastJsonMessage] )
    }
  }, [lastJsonMessage])

  function handleSend() {
    sendJsonMessage({
      Sender: user,
      Content: msg
    })
    //Reset field of inputs
    setUser('')
    setMsg('')
  }

  return (
    <>
      <h1>MVP prototype</h1>
      <p>Message Log</p>
      <div className='msgLog'>
        { msgHistory.map( (elem, id) => {
          return <li key={id}>{elem.content}</li>
        })}
      </div>
      <p>Enter your name</p>
      <input 
        type="text"
        value={user}
        onChange={e => setUser(e.target.value)}
      />
      <p>Enter a Message</p>
      <textarea
        value={msg}
        onChange={e => setMsg(e.target.value)}
      ></textarea>
      <div>
        <button onClick={handleSend}>Send Message</button>
      </div>
      
    </>
  )
}

export default App
