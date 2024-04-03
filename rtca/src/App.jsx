import './App.css'
import { useEffect, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { Stack, Divider } from '@mui/joy'
import Sidebar from './components/Sidebar'
import ChatArea from './components/ChatArea'

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

  const displayMorseToggler=()=>setDisplayMorse( prev => !prev)

  return (
    <>
      <Stack
        direction="row"
        justifyContent="flex-start" // X Axis
        alignItems="flex-start"     // Y Axis
        spacing={2}
        divider={<Divider orientation='vertical'></Divider>}
        sx={{margin: "3%"}}
      >
        <Sidebar 
          displayMorse={displayMorse}
          btnOnClickHandler={displayMorseToggler}
        />
        <ChatArea
          msgHistory={msgHistory}
          username={username}
          displayMorse={displayMorse}
          handleSend={handleSend}
          msg={msg}
          setMsg={setMsg}
        />
        </Stack>
    </>
  )
}

export default App
