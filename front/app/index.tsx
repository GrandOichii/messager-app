import { SafeAreaView, Text, View } from 'react-native'
import Login from './components/login'
import styles from '../styles/styles'
import { useEffect, useState } from 'react'
import { getStored } from '../storage'
import api from '../api'
import ChatList from './components/chatlist'
import Home from './components/home'

const Index = () => {

    const [loggedIn, setLoggedIn] = useState(false)
    
    const checkLogin = async () => {        
        const token = await getStored('jwt_token')
        if (token) {
            api.defaults.headers.common = { 'Authorization': `Bearer ${token}` }
        }
        setLoggedIn(!!token)
        
    }

    useEffect(() => { checkLogin() }, [])

    return <SafeAreaView style={{flex: 1}}>
        { loggedIn 
           ? <Home /> 
           : <Login onLogin={checkLogin} />
        }
    </SafeAreaView>
}

export default Index