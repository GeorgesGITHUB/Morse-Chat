import { Box, Typography, Stack } from "@mui/joy";
import { useEffect, useState } from "react";
import MessageOptions from "./MessageOptions";

function MessageBubble({message, id, isMe, displayMorse}){
    const [localDisplayMorse, setLocalDisplayMorse] = useState(displayMorse)
    const [hovering, setHovering] = useState(false)

    // global display morse overrides the local, on update
    useEffect(()=>{
        setLocalDisplayMorse(displayMorse)
    }, [displayMorse])

    const localDisplayMorseToggler=()=>setLocalDisplayMorse(prev => !prev)
    const handleMouseEnter=()=>setHovering(true)
    const handleMouseLeave=()=>setHovering(false)

    return (
        <Stack
            sx={{margin: "10px"}}
            direction="row"
            justifyContent={isMe ? "flex-end" : "flex-start"}
            alignItems="center"
            spacing={2}
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
            
        >
            { 
                hovering && isMe && 
                <MessageOptions
                    handleTranslateToggle={localDisplayMorseToggler}
                    morseString={message.contentMorse}
                />
            }
            <Box key={id} >
                <Typography
                    level="body-md"
                    variant={isMe ? "solid" : "soft"}
                >
                    {
                        localDisplayMorse ? message.contentMorse: message.contentText
                    }
                </Typography>
            </Box>
            { 
                hovering && !isMe && 
                <MessageOptions
                    handleTranslateToggle={localDisplayMorseToggler}
                    morseString={message.contentMorse}
                />
            }
        </Stack>

    )
}

export default MessageBubble