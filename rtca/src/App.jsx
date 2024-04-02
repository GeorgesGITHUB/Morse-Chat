import './App.css'
import { useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { 
  Box, Stack, Grid, Typography, Textarea, Button, Input
  } from '@mui/joy'
import MessageBubble from './components/MessageBubble'

function App() {
  const WS_URL = 'ws://localhost:8080/ws'
  const {
    sendJsonMessage, 
    lastJsonMessage, 
    readyState
  } = useWebSocket(WS_URL, { share: true })

  // replace initial value of user when login is implemented
  const [user, setUser] = useState(Date.now().toString())
  const [msg, setMsg] = useState('')
  const [msgHistory, setMsgHistory] = useState([])

  useEffect(() => {
    if (lastJsonMessage !== null) {
      setMsgHistory( (prevHistory) => [...prevHistory, lastJsonMessage] )
    }
  }, [lastJsonMessage])

  function handleSend() {
    sendJsonMessage({
      sender: user,
      contentRaw: msg,
      contentText: '',
      contentMorse: ''
    })
    //Reset
    setMsg('')
  }

  function generateMessageBubbles(){
    const list = msgHistory.map( (elem, index) => {
      return (
        <MessageBubble 
          message={elem}
          key={index}
          myMessage={ elem.sender === user}
        />
      )
    })

    return list
    
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
        <Box sx={{ flex: 1, width: '80%'}}>
          {generateMessageBubbles()}
        </Box>
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
