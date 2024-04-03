import { Box } from "@mui/joy"
import MessageBubble from "./MessageBubble"

function MessageBubbles({messages, username, displayMorse}){
    return(
        <Box sx={{ flex: 1, width: '80%'}}>
            {
                messages.map( (elem, index) => {
                    return (
                      <MessageBubble 
                        message={elem}
                        key={index}
                        isMe={ elem.sender === username}
                        displayMorse={displayMorse}
                      />
                    )
                })
            }
        </Box>
    )
}

export default MessageBubbles