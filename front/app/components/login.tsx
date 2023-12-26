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

    const [email, setEmail] = useState('mymail@email.com')
    const [emailError, setEmailError] = useState('')
    
    const [handle, setHandle] = useState('coolhandle')
    const [handleError, setHandleError] = useState('')
    
    const [password, setPassword] = useState('pass')
    const [passwordError, setPasswordError] = useState('')

    const [failed, setFailed] = useState(false)
    const [handleButtonEnabled, setHandleButtonEnabled] = useState(false)

    const resetForm = () => {
        // TODO add back
        // setEmail('')
        // setHandle('')
        // setPassword('')
        setFailed(false)
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
        // TODO add exception checking
        setFailed(false)
        try {
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
        } catch (e) {
            setFailed(true)            
        }

    }

    const mapInputColor = (v: string) => v.length == 0 ? '#daf5f0' : '#ffa07a'

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
                <TextInput value={email} onChangeText={setEmail} style={[styles.formTextInput, {backgroundColor: mapInputColor(emailError)}]} />
            </View>
            {
                !isLogin &&
                <View style={styles.row}>
                    <View style={{justifyContent: 'center', flex: 1}}>
                        <Text style={styles.formLabel}>Handle: </Text>
                    </View>
                    <TextInput value={handle} onChangeText={setHandle} style={[styles.formTextInput, {backgroundColor: mapInputColor(handleError)}]} />
                </View>
            }
            <View style={styles.row}>
                <View style={{justifyContent: 'center', flex: 1}}>
                    <Text style={styles.formLabel}>Password: </Text>
                </View>
                <TextInput value={password} onChangeText={setPassword} style={[styles.formTextInput, {backgroundColor: mapInputColor(passwordError)}]} secureTextEntry={true} />
            </View>
            {/* // TODO figure out the disabled submit button color */}
            <TouchableOpacity onPress={handleSubmit} style={[{flex: 1, alignItems: 'center', margin: 3}, styles.submit, {backgroundColor: handleButtonEnabled ? '#ff69b4' : '#ff69b4' }]} disabled={!handleButtonEnabled}>
                <Text style={styles.submitText}>
                    {isLogin ? 'Login' : 'Register'}
                </Text>
            </TouchableOpacity>

            {
                failed &&
                <View style={{flex:1, margin: 4, alignItems: 'center'}}>
                    <Text style={{color: 'red'}}>Failed to { isLogin ? 'login' : 'register'}</Text>
                </View>
            }
        </View>
            
     </SafeAreaView>
}

export default Login