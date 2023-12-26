import { StyleSheet } from "react-native";

const styles = StyleSheet.create({
    border: {
        borderRadius: 7,
        borderColor: 'black',
        borderWidth: 2,
        borderRightWidth: 4,
        borderBottomWidth: 4
    },
    container: {
        justifyContent: 'center',
        alignItems: 'center',
        flex: 1,
        backgroundColor: "#FFFF00"
    },
    loginForm: {
        backgroundColor: "#7df9ff",
        padding: 10,
    },
    row: {
        margin: 3,
        flexDirection: 'row',
    },
    formLabel: {
        fontWeight: 'bold',
    },
    formTextInput: {
        padding: 2,
        flex: 1,
        borderRadius: 4,
        borderColor: 'black',
        borderWidth: 2,
        borderRightWidth: 3,
        borderBottomWidth: 3
    },
    submit: {
        padding: 3,
        borderRadius: 4,
        borderColor: 'black',
        borderWidth: 2,
        borderRightWidth: 3,
        borderBottomWidth: 3,
    },
    submitText: {
        fontWeight: 'bold',
        fontSize: 16
    },
    selectedLoginType: {
        backgroundColor: '#ffb2ef',
        padding: 3,
        borderRadius: 4,
        borderColor: 'black',
        borderWidth: 2,
        borderRightWidth: 3,
        borderBottomWidth: 3,
        
    },
    chatRow: {
        margin: 1,
        padding: 2,
        // flex: 1,
    }
})

export default styles