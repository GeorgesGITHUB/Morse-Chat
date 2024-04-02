import { Box } from "@mui/joy"
import MessageBubble from "./MessageBubble"

function MessageBubbles({messages, username}){
    return(
        <Box sx={{ flex: 1, width: '80%'}}>
            {
                messages.map( (elem, index) => {
                    return (
                      <MessageBubble 
                        message={elem}
                        key={index}
                        isMe={ elem.sender === username}
                      />
                    )
                })
            }
        </Box>
    )
}

export default MessageBubbles