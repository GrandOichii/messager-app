import { Text, View, ViewProps } from "react-native"
import { TextInput, TouchableOpacity } from "react-native-gesture-handler"
import styles from "../../styles/styles"
import { useState } from "react"
import api from "../../api"

interface CreateChatProps extends ViewProps {
    onCreate: (c: Chat) => void
}

const CreateChat = (props: CreateChatProps) => {
    // TODO change color based on whether the user handle exists or not

    const [newChatHandle, setNewChatHandle] = useState('')
    
    const handleCreate = async () => {
        try {
            
            const req = await api.post('/api/chats/create', {
                with: newChatHandle
            })
            const newChat = req.data as Chat
            props.onCreate(newChat)
            setNewChatHandle('')

        } catch (e) {
            console.log((e as Error).message);
        }
    }
    
    return <View style={styles.row}>
        <View style={[styles.formTextInput, {marginRight: 2, flexDirection: 'row', alignItems: 'center'}]} >
            <Text style={{paddingHorizontal: 2}}>
                @
            </Text>
            <TextInput 
                style={{flex: 1}}
                placeholderTextColor='#b5d2ad'
                placeholder="Enter user handle"
                onChangeText={setNewChatHandle}
                value={newChatHandle}
                spellCheck={false}
            />
        </View>
        
        <TouchableOpacity onPress={handleCreate} style={[{justifyContent: 'center', flex: 1, padding: 5}, styles.button]}>
            <Text style={{fontWeight: 'bold'}}>Create</Text>
        </TouchableOpacity>
    </View>
}

export default CreateChat