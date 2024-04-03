import { Box, Typography, Stack } from "@mui/joy";
import { Button } from '@mui/base/Button';
import { useEffect, useState } from "react";

function MessageBubble({message, id, isMe, displayMorse}){
    const [localDisplayMorse, setLocalDisplayMorse] = useState(displayMorse)

    useEffect(()=>{
        setLocalDisplayMorse(displayMorse)
    }, [displayMorse])

    function localDisplayMorseToggler() {
        setLocalDisplayMorse(prev => !prev)
    }

    return (
        <Stack
            sx={{margin: "10px"}}
            direction="row"
            justifyContent={
                isMe ? "flex-end" : "flex-start"
            }
            alignItems="center"
        >
            <Box key={id} >
                <Button onClick={localDisplayMorseToggler}>
                    <Typography
                        level="body-md"
                        variant={
                            isMe ? "solid" : "soft"
                        }
                    >
                        {
                            localDisplayMorse ? message.contentMorse: message.contentText
                        }
                    </Typography>
                </Button>
            </Box>
        </Stack>

    )
}

export default MessageBubble