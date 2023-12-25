import { Text, View } from "react-native"
import { SafeAreaView } from "react-native-safe-area-context"
import styles from "../../styles/styles"
import { useEffect, useState } from "react"
import api from "../../api"
import ChatRow from "./chatrow"

const ChatList = () => {

    const [chatIDs, setChatIDs] = useState([])

    useEffect(() => {
        fetchChats()
    }, [])

    const fetchChats = async () => {
        const req = await api.get('/api/chats')
        setChatIDs(req.data)
        console.log(chatIDs);
        
    }
    
    return <SafeAreaView style={{flex: 1, padding: 4}}>
        {chatIDs.map(chatID => (
            <ChatRow chatID={chatID} key={chatID} />
        ))}
    </SafeAreaView>
}

export default ChatList