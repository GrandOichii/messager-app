import { Pressable, Text, TouchableOpacity, View, ViewProps } from "react-native"
import styles from "../../styles/styles"
import { TextInput } from "react-native-gesture-handler"
import { useEffect, useState } from "react"

interface ChatDisplayProps extends ViewProps {
    chatID: string,
    onBack: () => void
}

const ChatDisplay = (props: ChatDisplayProps) => {

    const [message, setMessage] = useState('')

    const handleSend = async () => {
        // setMessage('')
        console.log(message);
        
    }

    useEffect(() => {
        fetchMessages()
    }, [])

    const fetchMessages = async () => {
        
    }

    return <View style={{flex: 1, margin: 5}}>
        <View style={[styles.border, styles.row, {alignItems: 'center'}]}>
            <Pressable onPress={props.onBack} style={[{padding: 4}]}>
                <Text>
                    {"<"}
                </Text>
            </Pressable>
            <Text style={{marginRight: 5, flex: 1, textAlign: 'right'}}>
                {props.chatID}
            </Text>
        </View>
        <View style={{flex: 1}}>

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