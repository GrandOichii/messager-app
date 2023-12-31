import { Text, View, ViewProps } from "react-native"

interface MessageRowProps extends ViewProps {
    message: Message
}

const MessageRow = (props: MessageRowProps) => {


    return <View>
        <Text>{`${props.message.uhandle}: ${props.message.text}`}</Text>
    </View>
}

export default MessageRow