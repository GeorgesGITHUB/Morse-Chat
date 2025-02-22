import { Stack, Textarea ,Button } from "@mui/joy"

function InputArea({handleSend, msg, setMsg}){
    return (
        <Stack
        direction="row"
        justifyContent="center"
        alignItems="flex-start"
        spacing={2}
        width="85%"
        >
            <Textarea
            disabled={false}
            minRows={2}
            placeholder="Type Something..."
            variant="outlined"
            onChange={e => setMsg(e.target.value)}
            value={msg}
            sx={{flexGrow: 5}}
            />
            <Button
                onClick={handleSend}
                disabled= {msg===""}
            >
                Send
            </Button>
        </Stack>
    )
}

export default InputArea