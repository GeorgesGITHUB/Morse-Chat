const BASE_URL = "http://localhost:8080"

async function postUser(username, password) {
    try {
        const resPromise = await fetch(`${BASE_URL}/api/users`,{
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({username, password})
        })
        
        if (!resPromise.ok) {
            throw new Error('Failed posting User')
        }

    } catch (error) {
        // bubbled up error to be handled by service caller
        throw new Error(error)
    }


}

export default postUser