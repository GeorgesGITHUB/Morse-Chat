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

  // placeholder user_id until login,auth complete
  const [user_id, setUser_id] = useState(0)
  // placeholder username until login,auth complete
  const [username, setUsername] = useState('')

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

  // placeholder until user_id and username can be set by user
  useEffect(() => {
    console.log('user_id',user_id,'username',username);
  }, [user_id, username]);
  // placeholder until user_id and username can be set by user
  const loadProfilePreset1 = ()=> {
    setUser_id("613")
    setUsername("Georges")
    console.log("Loaded Preset 1")
  }
  // placeholder until user_id and username can be set by user
  const loadProfilePreset2 = ()=> {
    setUser_id("961")
    setUsername("Elias")
    console.log("Loaded Preset 2")
  }
  // placeholder until user_id and username can be set by user
  const loadProfilePreset3 = ()=> {
    setUser_id("627")
    setUsername("John")
    console.log("Loaded Preset 3")
  }

  return (
    <>
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
          loadProfilePreset1={loadProfilePreset1}
          loadProfilePreset2={loadProfilePreset2}
          loadProfilePreset3={loadProfilePreset3}
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
    </>
  )
}

export default App
