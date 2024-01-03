import { JwtPayload, jwtDecode } from "jwt-decode"
import { Image, Text, View, ViewProps } from "react-native"
import { getStored } from "../../storage"
import { useEffect, useState } from "react"
import styles from "../../styles/styles"
import { getAvatar } from "../../avatars"

interface MessageRowProps extends ViewProps {
    message: Message,
    myHandle: string
}

const MessageRow = (props: MessageRowProps) => {
    const [avatarUri, setAvatarUri] = useState('')

    const isMe = props.myHandle === props.message.uhandle
    const meColor =  '#b5d2ad'
    const otherColor = '#f8d6b3'

    useEffect(() => {
        (async () => {
            const aUri = await getAvatar(props.message.uhandle)
            setAvatarUri(aUri)
            console.log(aUri);
            
        })()
    }, [])

    return <View style={[{flex: 1, margin: 2, alignItems: isMe ? 'flex-end' : 'flex-start', flexDirection: isMe ? 'row-reverse' : 'row'}]}>
        <Image
            source={{
                uri: avatarUri
            }}
            style={{
                width: 40,
                height: 40
            }}
        />
        <Text
            style={[styles.border, {paddingVertical: 2, paddingHorizontal: 4, backgroundColor: isMe ? meColor : otherColor}]}
        >
            {props.message.text}
        </Text>
    </View>
}

export default MessageRow