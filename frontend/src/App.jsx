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
  } = useWebSocket(WS_URL, { share: true })

  const user_id = sessionStorage.getItem('user_id')
  const username = sessionStorage.getItem('username')
  const password = sessionStorage.getItem('password')

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
      sender_id: user_id,
      username: username,
      contentRaw: msg,
    })
    //Reset
    setMsg('')
  }

  const displayMorseToggler=()=>setDisplayMorse( prev => !prev)

  return (
    <Stack
      direction="row"
      justifyContent="flex-start" // X Axis
      alignItems="flex-start"     // Y Axis
      spacing={1}
      divider={<Divider orientation='vertical'></Divider>}
    >
      <Sidebar 
        displayMorse={displayMorse}
        btnOnClickHandler={displayMorseToggler}
      />
      <ChatArea
        msgHistory={msgHistory}
        user_id={user_id}
        displayMorse={displayMorse}
        handleSend={handleSend}
        msg={msg}
        setMsg={setMsg}
      />
    </Stack>
  )
}

export default App
