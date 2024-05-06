import { useState } from 'react'
import { Button, Stack, Typography, Input, Divider } from '@mui/joy'
import { useNavigate } from 'react-router-dom';
import fetchUsername from './services/validateUser';


function Login() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [renderLogin, setRenderLogin] = useState(true)
    const navigate = useNavigate()

    const toggleRenderLogin = () => setRenderLogin(prev => !prev)

    async function handleLoginSubmit(){
        // try {
        //     const response = await fetchUsername(username)
        //     console.log(`fetched response: ${response}`)
        //     navigate('/app')
        // } catch (error) {
        //     console.error(error)
        // }        
    }

    return (
        <>
            { renderLogin &&
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
                <Input
                    placeholder='Password'
                    onChange={e => setPassword(e.target.value)}
                />
                <Button
                    disabled={username===''||password===''}
                    onClick={handleLoginSubmit}
                    >
                    Submit
                </Button>
                <Button
                    onClick={toggleRenderLogin}
                    >
                    Create Profile
                </Button>
                <Divider/>
            </Stack>}
            { !renderLogin &&
            <Stack
                direction="column"
                justifyContent="center"
                alignItems="center"
                spacing={3}
                height="50vh" // % of the viewport

            >
                <Divider>Morse Chat</Divider>
                <Typography level="h1">Create Profile</Typography>
                <Input
                    placeholder='Username'
                    onChange={e => setUsername(e.target.value)}
                />
                <Input
                    placeholder='Password'
                    onChange={e => setPassword(e.target.value)}
                />
                <Button
                    disabled={username===''||password===''}
                    //onClick={}
                    >
                    Submit
                </Button>
                <Divider/>
            </Stack>}
        </>
    )
}

export default Login