import { Box } from "@mui/joy"
import MessageBubble from "./MessageBubble"

function MessageBubbles({messages, user_id, displayMorse}){
    return(
        <Box sx={{ flex: 1, width: '80%'}}>
            {
                messages.map( (elem, index) => {
                    return (
                      <MessageBubble 
                        message={elem}
                        key={index}
                        isMe={ elem.sender_id === user_id}
                        displayMorse={displayMorse}
                      />
                    )
                })
            }
        </Box>
    )
}

export default MessageBubbles