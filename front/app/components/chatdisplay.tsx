import { Pressable, Text, TouchableOpacity, View, ViewProps } from "react-native"
import styles from "../../styles/styles"
import { ScrollView, TextInput } from "react-native-gesture-handler"
import { useEffect, useState } from "react"
import api from "../../api"
import { JwtPayload, jwtDecode } from "jwt-decode"
import { getStored } from "../../storage"
import MessageRow from "./messagerow"

interface ChatDisplayProps extends ViewProps {
    chatID: string,
    onBack: () => void
}

const ChatDisplay = (props: ChatDisplayProps) => {
    const [socket, setSocket] = useState<WebSocket | null>(null)

    const [message, setMessage] = useState('')
    const [messages, setMessages] = useState<Message[]>([])

    const handleSend = async () => {
        // TODO check for errors
        const req = api.post('/api/chats/addmessage', {
            chat_id: props.chatID,
            text: message
        })
        setMessage('')
    }

    useEffect(() => {

        const connWs = async () => {
            if (props.chatID == '') return

            const token = await getStored('jwt_token')

            // TODO add back
            // const host = api.defaults.baseURL
            const host = 'localhost:8080'
            const jwtData = jwtDecode(token!) as JwtPayload & {
                handle: string
            }
            const handle = jwtData.handle
            const sock = new WebSocket(`ws://${host}/api/chats/listen?chatid=${props.chatID}&handle=${handle}`)
            sock.onopen = async () => {
                sock.send(token!)
            }

            sock.onclose = async (e) => {
                console.error(e);
            }

            sock.onmessage = async (e) => {
                const m = JSON.parse(e.data) as Message
                setMessages([
                    ...messages,
                    m
                ])
            }

            setSocket(sock)
        }

        connWs()
        fetchMessages()
    }, [])

    const fetchMessages = async () => {
        // TODO
    }
    
    const handleBack = async () => {        
        socket!.close()
        props.onBack()
    }

    return <View style={{flex: 1, margin: 5}}>
        <View style={[styles.border, styles.row, {alignItems: 'center'}]}>
            <Pressable onPress={handleBack} style={[{padding: 4}]}>
                <Text>
                    {"<"}
                </Text>
            </Pressable>
            <Text style={{marginRight: 5, flex: 1, textAlign: 'right'}}>
                {props.chatID}
            </Text>
        </View>
        <View style={{flex: 1}}>
            <ScrollView>
                {messages.map((m, index) => (
                    <MessageRow key={index} message={m} />
                ))}
            </ScrollView>
        </View>
        <View style={[{flexDirection: 'row'}]}>
            <TextInput
                placeholder="Enter message"
                style={[styles.formTextInput, {padding: 4, marginRight: 3}]}
                value={message}
                onChangeText={setMessage}
            />
            <Pressable onPress={handleSend} style={[styles.submit, {justifyContent: 'center', alignItems: 'center'}]}>
                <Text>
                    Send
                </Text>
            </Pressable>
        </View>
    </View>
}

export default ChatDisplay