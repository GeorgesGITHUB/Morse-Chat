import './App.css'
import { useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { 
  Box, Stack, Grid, Typography, Textarea, Button, Input
  } from '@mui/joy'
import MessageBubble from './components/MessageBubble'
import MessageBubbles from './components/MessageBubbles'

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
        <Stack
          direction="row"
          justifyContent="center"
          alignItems="flex-start"
          spacing={2}
          width="80%"
          sx={{alignItems: 'stretch'}}
        >
          <Textarea
            disabled={false}
            minRows={2}
            placeholder="Type Something..."
            variant="outlined"
            onChange={e => setMsg(e.target.value)}
            value={msg}
            sx={{flexGrow: 5}}
          />
          <Button
              sx={{flexGrow: 1}}
              onClick={handleSend}
          >
              Send
          </Button>
        </Stack>
      </Stack>
    </>
  )
}

export default App
