import { Box, Typography, Stack } from "@mui/joy";

function MessageBubble({message, id, isMe}){
    return (
        <Stack
        direction="row"
        justifyContent={
            isMe ? "flex-end" : "flex-start"
        }
        alignItems="center"
        spacing={0}
        >
            <Box key={id}>
                <Typography
                    level="body-md"
                    variant={
                        isMe ? "soft" : "solid"
                    }
                >
                    {message.contentRaw}
                </Typography>
            </Box>
        </Stack>

    )
}

export default MessageBubble