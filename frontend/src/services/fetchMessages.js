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
        // { message: [ {'message_id', 1}, ... ,{'message_id', n} ]}
        const jsonPromise = await resPromise.json()
        
        if (jsonPromise.messages===null){
            return []
        }
        
        // returning promise
        // [ {'message_id', 1}, ... ,{'message_id', n} ]
        return jsonPromise.messages

    } catch (error) {
        // bubbled up error to be handled by service caller
        throw new Error(error)
    }


}

export default fetchMessages