import { useState } from 'react'
import { Button, Stack, Typography, Input, Divider } from '@mui/joy'
import { useNavigate } from 'react-router-dom';


function Login() {
    const [username, setUsername] = useState('')
    const navigate = useNavigate()

    function handleSubmit(){
        // Add authentication logic here
        console.log('authentication logic missing')
        
        navigate('/app')
    }

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
                onClick={handleSubmit}
                >
                Submit
            </Button>
            <Divider/>
        </Stack>
    )
}

export default Login