import { Typography, Stack } from "@mui/joy"
import MessageBubbles from "./MessageBubbles"
import InputArea from "./InputArea"

function ChatArea({
    msgHistory, user_id, displayMorse, handleSend, msg, setMsg}) {
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={2}
            sx={{minWidth: "85%"}}
        >
            <Typography level="h1">Morse Chat</Typography>            
            <MessageBubbles
            messages={msgHistory}
            user_id={user_id}
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
    )
}

export default ChatArea