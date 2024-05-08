import { useState } from 'react'
import { Button, Stack, Typography, Input, Divider, Box } from '@mui/joy'
import { Link, useNavigate } from 'react-router-dom';
import fetchUserID from './services/fetchUserID';

function Login() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [inputErr, setInputErr] = useState(false)
    const navigate = useNavigate()

    async function handleLoginProfile(){
        try {
            const user_id = await fetchUserID(username,password)
            // Here until authentication is implemented
            // *******************************************
            sessionStorage.removeItem('user_id')
            sessionStorage.removeItem('username')
            sessionStorage.removeItem('password')
            sessionStorage.setItem('user_id',user_id)
            sessionStorage.setItem('username',username)
            sessionStorage.setItem('password',password)
            // *******************************************
            navigate('/app')
            
        } catch (error) {
            console.error(error)
            setInputErr(true)
            setUsername('')
            setPassword('')
        }
    }

    return (
        <Stack
            direction="column"
            justifyContent="center"
            alignItems="center"
            spacing={2}
            height="50vh" // % of the viewport
        >
            <Divider>Morse Chat</Divider>
            <Typography level="h1">Login</Typography>
            <Input
                placeholder='Username'
                value={username}
                onChange={e => {
                    setUsername(e.target.value)
                    setInputErr(false)
                }}
                error={inputErr}
            />
            <Input
                placeholder='Password'
                value={password}
                onChange={e => {
                    setPassword(e.target.value)
                    setInputErr(false)
                }}
                error={inputErr}
            />
            <Button
                //disabled={username===''||password===''}
                onClick={handleLoginProfile}
                >
                Submit
            </Button>
            <Link to="/CreateProfile">
                <Button>Create Profile</Button>
            </Link>
            <Divider/>
        </Stack>
    )
}

export default Login