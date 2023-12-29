import { useEffect, useState } from "react"
import ChatList from "./chatlist"
import ChatDisplay from "./chatdisplay"
import api from "../../api"

const Home = () => {
    const [chatID, setChatID] = useState('')

    const handleChatPress = (cid: string) => {
        setChatID(cid)
    }

    useEffect(() => {
        const connWs = async () => {
            if (chatID == '') return
            
            // const host = api.defaults.baseURL
            const host = 'localhost:8080'
            const socket = new WebSocket(`ws://${host}/api/chats/listen/${chatID}`)
            socket.onopen = () => {
                console.log('connected!');
                
            }
        }

        connWs()
    }, [chatID])

    return chatID.length === 0
        ? <ChatList onChatPress={handleChatPress}/>
        : <ChatDisplay chatID={chatID} onBack={() => setChatID('')} />
}

export default Home