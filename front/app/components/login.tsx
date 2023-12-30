import { Alert, Pressable, SafeAreaView, Text, TextInput, TouchableOpacity, View, ViewProps } from "react-native"
import styles from "../../styles/styles"
import { ComponentProps, useEffect, useState } from "react"
import api from "../../api"
import { setStored } from "../../storage"

interface LoginProps extends ViewProps {
    onLogin: () => void
}

const Login = (props: LoginProps) => {
    const [isLogin, setIsLogin] = useState(true)

    const [email, setEmail] = useState('mymail@email.com')
    const [emailError, setEmailError] = useState('')
    
    const [handle, setHandle] = useState('coolhandle')
    const [handleError, setHandleError] = useState('')
    
    const [password, setPassword] = useState('pass')
    const [passwordError, setPasswordError] = useState('')

    const [failedLabel, setFailedLabel] = useState('')
    const [handleButtonEnabled, setHandleButtonEnabled] = useState(false)

    const [processing, setProcessing] = useState(false)

    const resetForm = () => {
        // TODO add back
        // setEmail('')
        // setHandle('')
        // setPassword('')
        setFailedLabel('')
    }

    new Array<[string, React.Dispatch<React.SetStateAction<string>>, RegExp]>(

        [handle, setHandleError, /^.{4,}$/],
        [password, setPasswordError, /^.{4,}$/],
        [email, setEmailError, /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/],
    
        ).forEach(([state, setErrorFunc, pattern]) => {
        useEffect(() => {
            const cond = !pattern.test(state)
            
            setErrorFunc(cond ? 'error' : '')
            setHandleButtonEnabled(!cond)
        }, [state])
    })

    useEffect(resetForm, [isLogin])

    const handleSubmit = async () => {        
        setFailedLabel('')
        setProcessing(true)
        
        // TODO split register and login errors?
        try {
            if (!isLogin) {
                const req = await api.post('/api/users/register', {
                    email: email,
                    handle: handle,
                    password: password,
                })        
            }
        } catch (e) {
            setFailedLabel('register')
            setProcessing(false)
            return
        }
            
        try {
            const req = await api.post('/api/users/login', {
                email: email,
                password: password,
            })
            await setStored('jwt_token', req.data.token)
            resetForm()
            props.onLogin()
        } catch (e) {
            setFailedLabel('login')
            setProcessing(false)
            return
        }
        
        setProcessing(false)
    }

    const canSubmit = () => !processing && handleButtonEnabled

    const mapInputColor = (v: string) => v.length === 0 ? '#daf5f0' : '#ffa07a'

    return <SafeAreaView style={styles.container}>
        <View style={[styles.loginForm, styles.border]}>
            <View style={styles.row}>
                <Pressable disabled={processing} onPress={() => setIsLogin(true)} style={[{flex:1, justifyContent: 'center', alignItems: 'center'}, isLogin ? styles.selectedLoginType : null]}>
                    <Text>Login</Text>
                </Pressable>
                <Pressable disabled={processing} onPress={() => setIsLogin(false)} style={[{flex:1, justifyContent: 'center', alignItems: 'center'}, !isLogin ? styles.selectedLoginType : null]}>
                    <Text>Register</Text>
                </Pressable>
            </View>
            <View style={styles.row}>
                <View style={{justifyContent: 'center', flex: 1}}>
                    <Text style={styles.formLabel}>Email: </Text>
                </View>
                <TextInput editable={!processing} value={email} onChangeText={setEmail} style={[styles.formTextInput, {backgroundColor: mapInputColor(emailError)}]} />
            </View>
            {
                !isLogin &&
                <View style={styles.row}>
                    <View style={{justifyContent: 'center', flex: 1}}>
                        <Text style={styles.formLabel}>Handle: </Text>
                    </View>
                    <TextInput editable={!processing} value={handle} onChangeText={setHandle} style={[styles.formTextInput, {backgroundColor: mapInputColor(handleError)}]} />
                </View>
            }
            <View style={styles.row}>
                <View style={{justifyContent: 'center', flex: 1}}>
                    <Text style={styles.formLabel}>Password: </Text>
                </View>
                <TextInput editable={!processing} value={password} onChangeText={setPassword} style={[styles.formTextInput, {backgroundColor: mapInputColor(passwordError)}]} secureTextEntry={true} />
            </View>
            {/* // TODO figure out the disabled submit button color */}
            <Pressable onPress={handleSubmit} style={[{flex: 1, alignItems: 'center', margin: 3}, styles.submit, {backgroundColor: canSubmit() ? '#ff69b4' : '#ff69b4' }]} disabled={!canSubmit()}>
                <Text style={styles.submitText}>
                    {isLogin ? 'Login' : 'Register'}
                </Text>
            </Pressable>

            {
                failedLabel != '' &&
                <View style={{flex:1, margin: 4, alignItems: 'center'}}>
                    <Text style={{color: 'red'}}>Failed to {failedLabel}</Text>
                </View>
            }
        </View>
            
     </SafeAreaView>
}

export default Login