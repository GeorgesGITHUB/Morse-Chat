import { useState } from 'react'
import { Button, Stack, Typography, Input, Divider } from '@mui/joy'
import { useNavigate } from 'react-router-dom';
import postUser from './services/postUser';
import fetchUserID from './services/fetchUserID';

function CreateProfile(){
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [inputErr, setInputErr] = useState(false)
    const navigate = useNavigate()

    async function handleCreateProfile() {
        try {
            await fetchUserID(username,password)
            setInputErr(true)
            setUsername('')
            setPassword('')
            return
        } catch (error) {
            console.log(`${username} does not exist`)
        }

        try {
            console.log(`creating user ${username}`)
            await postUser(username,password)
            console.log('fetching user_id')
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
            
        } catch (error) { // unlikely to catch
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
            <Typography level="h1">Create Profile</Typography>
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
                disabled={username===''||password===''}
                onClick={handleCreateProfile}
            >
                Submit
            </Button>
            <Divider/>
            { inputErr &&
            <Typography
                level='body-sm'
                color='danger'
            >
                That username already exists
            </Typography>}
        </Stack>
    )
}

export default CreateProfile