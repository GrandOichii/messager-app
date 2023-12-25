import { Text, View, ViewProps } from "react-native"

interface ChatRowProps extends ViewProps {
    chatID: string
}

const ChatRow = (props: ChatRowProps) => {
    return <View>
        <Text>
            {props.chatID}
        </Text>
    </View>   
}

export default ChatRow