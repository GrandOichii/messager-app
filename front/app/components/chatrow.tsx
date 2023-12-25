import { Text, View, ViewProps } from "react-native"
import styles from "../../styles/styles"
import { TouchableOpacity } from "react-native-gesture-handler"

interface ChatRowProps extends ViewProps {
    chatID: string
}

// TODO change background color based on whether the chat is read or not

const ChatRow = (props: ChatRowProps) => {
    const onPress = async () => {
        console.log('press');
    }

    return <View style={[styles.chatRow, styles.border]}>
        <TouchableOpacity onPress={onPress}>
            <Text>
                {props.chatID}
            </Text>
        </TouchableOpacity>
    </View>   
}

export default ChatRow