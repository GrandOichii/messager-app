import { useState } from "react"
import ChatList from "./chatlist"
import ChatDisplay from "./chatdisplay"

const Home = () => {
    const [chatID, setChatID] = useState('')

    const handleChatPress = (cid: string) => {
        setChatID(cid)
    }

    return chatID.length === 0
        ? <ChatList onChatPress={handleChatPress}/>
        : <ChatDisplay chatID={chatID} onBack={() => setChatID('')} />
}

export default Home