const BASE_URL = "http://localhost:8080"

async function fetchMessages() {
    try {
        const resPromise = await fetch(`
            ${BASE_URL}/api/messages
        `)
        
        if (!resPromise.ok) {
            throw new Error('Bad response fetching all Messages')
        }

        // converting into a promise of json
        const jsonPromise = await resPromise.json()

        // returning promise
        return jsonPromise.messages

    } catch (error) {
        // bubbled up error to be handled by service caller
        throw new Error(error)
    }


}

export default fetchMessages