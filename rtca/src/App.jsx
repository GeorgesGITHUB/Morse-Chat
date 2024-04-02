import './App.css'
import { useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { 
  Box, Stack, Grid, Typography
  } from '@mui/joy'
import MessageBubble from './components/MessageBubble'
import MessageBubbles from './components/MessageBubbles'
import InputArea from './components/InputArea'

function App() {
  const WS_URL = 'ws://localhost:8080/ws'
  const {
    sendJsonMessage, 
    lastJsonMessage, 
    readyState
  } = useWebSocket(WS_URL, { share: true })

  // placeholder username generation. Replace when login,auth completed
  const [username, setusername] = useState(Date.now().toString())
  const [msg, setMsg] = useState('')
  const [msgHistory, setMsgHistory] = useState([])

  useEffect(() => {
    if (lastJsonMessage !== null) {
      setMsgHistory( (prevHistory) => [...prevHistory, lastJsonMessage] )
    }
  }, [lastJsonMessage])

  function handleSend() {
    sendJsonMessage({
      sender: username,
      contentRaw: msg,
      contentText: '',
      contentMorse: ''
    })
    //Reset
    setMsg('')
  }

  return (
    <>
      <Stack
        direction="column"
        justifyContent="center"
        alignItems="center"
        spacing={2}
        mx="25px"
      >
        <Typography
          level="h1"
        >
          Joy UI-fication of Morse Chat
        </Typography>
        <MessageBubbles
          messages={msgHistory}
          username={username}
        >
        </MessageBubbles>
        <InputArea
          handleSend={handleSend}
          msg={msg}
          setMsg={setMsg}
        >
        </InputArea>
      </Stack>
    </>
  )
}

export default App
