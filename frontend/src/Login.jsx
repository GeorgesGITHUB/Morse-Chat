import { useState } from 'react'
import { Button, Stack, Typography, Input, Divider } from '@mui/joy'


function Login() {
    const [username, setUsername] = useState('')

    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={3}
            height="50vh" // % of the viewport

        >
            <Divider>Morse Chat</Divider>
            <Typography level="h1">Login</Typography>
            <Input
                placeholder='Username'
                onChange={e => setUsername(e.target.value)}
                />
            <Button
                disabled={username===''}
                //onClick={}
                >
                Submit
            </Button>
            <Divider/>
        </Stack>
    )
}

export default Login