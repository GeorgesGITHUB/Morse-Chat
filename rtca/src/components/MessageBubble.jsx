import { Box, Typography, Stack } from "@mui/joy";

function MessageBubble({message, id, myMessage}){
    console.log('As',message.sender,'myMessage is',myMessage)
    return (
        <Stack
        direction="row"
        justifyContent={
            myMessage ? "flex-end" : "flex-start"
        }
        alignItems="center"
        spacing={0}
        >
            <Box key={id}>
                <Typography
                    level="body-md"
                    variant={
                        myMessage ? "soft" : "solid"
                    }
                >
                    {message.contentRaw}
                </Typography>
            </Box>
        </Stack>

    )
}

export default MessageBubble