const BASE_URL = "http://localhost:8080"

async function fetchUserID(username, password) {
    try {
        const resPromise = await fetch(`
            ${BASE_URL}/api/users/id?username=${username}&password=${password}
        `)
        
        if (!resPromise.ok) {
            throw new Error('Failed fetching user_id')
        }

        // converting into a promise of json
        const jsonPromise = await resPromise.json()
        
        // returning promise for the value of key: user_id 
        return jsonPromise.user_id

    } catch (error) {
        // bubbled up error to be handled by service caller
        throw new Error(error)
    }


}

export default fetchUserID