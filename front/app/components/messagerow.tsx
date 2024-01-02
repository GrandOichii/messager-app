import { JwtPayload, jwtDecode } from "jwt-decode"
import { Text, View, ViewProps } from "react-native"
import { getStored } from "../../storage"
import { useEffect } from "react"

interface MessageRowProps extends ViewProps {
    message: Message,
    myHandle: string
}

const MessageRow = (props: MessageRowProps) => {
    
    
    // useEffect(() => {
    //     const token = await getStored('jwt_token')
    
    //     const jwtData = jwtDecode(token!) as JwtPayload & {
    //         handle: string
    //     }
    //     const handle = jwtData.handle

    // })


    return <View style={[{flex: 1, backgroundColor: 'cyan'}]}>
        <Text
            style={{textAlign: props.myHandle === props.message.uhandle ? 'left' : 'right'}}
        >
            {`${props.message.uhandle}: ${props.message.text}`}
        </Text>
    </View>
}

export default MessageRow