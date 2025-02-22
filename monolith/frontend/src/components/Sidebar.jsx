import { Stack, Button, Typography, Divider } from "@mui/joy"

function Sidebar({displayMorse, btnOnClickHandler}){
    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={1}
            sx={{minWidth: "15%", marginTop: "5% !important", marginBottom: "5% !important"}}
            >
            <Divider/>
            <Button onClick={btnOnClickHandler}>
                {
                    displayMorse ? "To Text" : "To Morse"
                }
            </Button>
            <Divider/>
        </Stack>
    )
}

export default Sidebar