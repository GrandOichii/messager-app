import { ScrollView, Text, View, ViewProps } from "react-native"
import { SafeAreaView, SafeAreaViewProps } from "react-native-safe-area-context"
import styles from "../../styles/styles"
import { useEffect, useState } from "react"
import api from "../../api"
import ChatRow from "./chatrow"
import { TouchableOpacity } from "react-native-gesture-handler"
import CreateChat from "./createchat"

interface ChatListProps extends ViewProps {
    onChatPress: (chatID: string) => void    
}

const ChatList = (props: ChatListProps) => {

    const [chatIDs, setChatIDs] = useState<string[]>([])

    useEffect(() => {
        fetchChats()        
    }, [])

    const fetchChats = async () => {
        const req = await api.get('/api/chats')
        setChatIDs(req.data)        
    }
    
    return <View style={{flex: 1, padding: 4}}>
        <CreateChat onCreate={chat => setChatIDs([...chatIDs, chat.id])}/>
        {
            chatIDs.length === 0 
            ? <View style={{alignItems: 'center', flex: 1, justifyContent: 'center'}}>
                <Text>
                    No chats yet!
                </Text>
            </View>
            : <ScrollView>
                {chatIDs.map(chatID => (
                    <ChatRow chatID={chatID} key={chatID} onPress={props.onChatPress} />
                ))}
            </ScrollView>
        }
    </View>
}

export default ChatList