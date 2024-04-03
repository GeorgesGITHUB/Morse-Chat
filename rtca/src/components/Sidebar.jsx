import { Stack, Button } from "@mui/joy"

function Sidebar({displayMorse, btnOnClickHandler}){
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={1}
            sx={{minWidth: "15%", marginTop: "5% !important", marginBottom: "5% !important"}}
        >
            <Button onClick={btnOnClickHandler}>
                {
                    displayMorse ? "To Text" : "To Morse"
                }
            </Button>
        </Stack>
    )
}

export default Sidebar