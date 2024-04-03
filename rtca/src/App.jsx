import './App.css'
import { useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { 
  Box, Stack, Grid, Typography, Divider
  } from '@mui/joy'
import MessageBubbles from './components/MessageBubbles'
import InputArea from './components/InputArea'
import Sidebar from './components/Sidebar'

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
  const [displayMorse, setDisplayMorse] = useState(false)

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

  function displayMorseToggler() {
    setDisplayMorse( prev => !prev)
  }

  return (
    <>
      <Stack
        direction="row"
        justifyContent="flex-start"
        alignItems="flex-start"
        spacing={2}
        divider={<Divider orientation='vertical'></Divider>}
      >
        <Sidebar 
          displayMorse={displayMorse}
          btnOnClickHandler={displayMorseToggler}
        ></Sidebar>
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
            Morse Chat
          </Typography>
          <MessageBubbles
            messages={msgHistory}
            username={username}
            displayMorse={displayMorse}
          >
          </MessageBubbles>
          <InputArea
            handleSend={handleSend}
            msg={msg}
            setMsg={setMsg}
          >
          </InputArea>
        </Stack>
        </Stack>
    </>
  )
}

export default App
