import { Typography, Stack } from "@mui/joy"
import MessageBubbles from "./MessageBubbles"
import InputArea from "./InputArea"

function ChatArea({
    msgHistory, username, displayMorse, handleSend, msg, setMsg
}) {
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={2}
            sx={{minWidth: "100%"}}
        >
            <Typography level="h1">Morse Chat</Typography>
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
    )
}

export default ChatArea