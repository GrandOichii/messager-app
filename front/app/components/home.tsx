import { useEffect, useState } from "react"
import ChatList from "./chatlist"
import ChatDisplay from "./chatdisplay"
import api from "../../api"
import { getStored } from "../../storage"
import { JwtPayload, jwtDecode } from "jwt-decode"

const Home = () => {
    const [chatID, setChatID] = useState('')

    const handleChatPress = async (cid: string) => {
        setChatID(cid)
    }

    const handleBack = async () => {
        setChatID('')
    }

    return chatID.length === 0
        ? <ChatList onChatPress={handleChatPress}/>
        : <ChatDisplay chatID={chatID} onBack={handleBack}  />
}

export default Home