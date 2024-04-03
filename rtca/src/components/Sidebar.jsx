import { Stack, Button } from "@mui/joy"
function Sidebar({displayMorse, btnOnClickHandler}){
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={1}
        >
            <Button onClick={btnOnClickHandler}>
                {
                    displayMorse ? "Text" : "Morse"
                }
            </Button>
        </Stack>
    )
}

export default Sidebar