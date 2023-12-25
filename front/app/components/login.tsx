import { Alert, SafeAreaView, Text, TextInput, TouchableOpacity, View, ViewProps } from "react-native"
import styles from "../../styles/styles"
import { ComponentProps, useEffect, useState } from "react"
import api from "../../api"
import { setStored } from "../../storage"

interface LoginProps extends ViewProps {
    onLogin: () => void
}

const Login = (props: LoginProps) => {
    // TODO add validation checking
    // TODO add disabling button clicking when processing requests 

    const [isLogin, setIsLogin] = useState(true)

    const [email, setEmail] = useState('mymail@mail.com')
    const [handle, setHandle] = useState('coolhandle')
    const [password, setPassword] = useState('pass')

    const resetForm = () => {
        // TODO add back
        // setEmail('')
        // setHandle('')
        // setPassword('')
    }

    useEffect(resetForm, [isLogin])

    const handleSubmit = async () => {        
        // TODO add exception checking
        if (!isLogin) {
            const req = await api.post('/api/users/register', {
                email: email,
                handle: handle,
                password: password,
            })        
        }

        const req = await api.post('/api/users/login', {
            email: email,
            password: password,
        })

        await setStored('jwt_token', req.data.token)
        resetForm()
        props.onLogin()
    }

    return <SafeAreaView style={styles.container}>
        <View style={[styles.loginForm, styles.border]}>
            <View style={styles.row}>
                <TouchableOpacity onPress={() => setIsLogin(true)} style={[{flex:1, justifyContent: 'center', alignItems: 'center'}, isLogin ? styles.selectedLoginType : null]}>
                    <Text>Login</Text>
                </TouchableOpacity>
                <TouchableOpacity onPress={() => setIsLogin(false)} style={[{flex:1, justifyContent: 'center', alignItems: 'center'}, !isLogin ? styles.selectedLoginType : null]}>
                    <Text>Register</Text>
                </TouchableOpacity>
            </View>
            <View style={styles.row}>
                <View style={{justifyContent: 'center', flex: 1}}>
                    <Text style={styles.formLabel}>Email: </Text>
                </View>
                <TextInput value={email} onChangeText={setEmail} style={styles.formTextInput} />
            </View>
            {
                !isLogin &&
                <View style={styles.row}>
                    <View style={{justifyContent: 'center', flex: 1}}>
                        <Text style={styles.formLabel}>Handle: </Text>
                    </View>
                    <TextInput value={handle} onChangeText={setHandle} style={styles.formTextInput} />
                </View>
            }
            <View style={styles.row}>
                <View style={{justifyContent: 'center', flex: 1}}>
                    <Text style={styles.formLabel}>Password: </Text>
                </View>
                <TextInput value={password} onChangeText={setPassword} style={styles.formTextInput} secureTextEntry={true} />
            </View>
            <TouchableOpacity onPress={handleSubmit} style={[{flex: 1, alignItems: 'center', margin: 3}, styles.submit]}>
                <Text style={styles.submitText}>
                    {isLogin ? 'Login' : 'Register'}
                </Text>
            </TouchableOpacity>
        </View>
            
     </SafeAreaView>
}

export default Login