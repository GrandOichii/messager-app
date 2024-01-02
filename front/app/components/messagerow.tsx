import { JwtPayload, jwtDecode } from "jwt-decode"
import { Text, View, ViewProps } from "react-native"
import { getStored } from "../../storage"
import { useEffect } from "react"
import styles from "../../styles/styles"

interface MessageRowProps extends ViewProps {
    message: Message,
    myHandle: string
}

const MessageRow = (props: MessageRowProps) => {

    const isMe = props.myHandle === props.message.uhandle
    const meColor =  '#b5d2ad'
    const otherColor = '#f8d6b3'

    return <View style={[{flex: 1, margin: 2, alignItems: isMe ? 'flex-end' : 'flex-start'}]}>
        <Text
            style={[styles.border, {paddingVertical: 2, paddingHorizontal: 4, backgroundColor: isMe ? meColor : otherColor}]}
        >
            {props.message.text}
        </Text>
    </View>
}

export default MessageRow