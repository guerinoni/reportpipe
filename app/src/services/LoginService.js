import {apiUrl} from "../constants/RoutesUrl";

async function login(user) {
    const body = {
        email: user?.email,
        password: user?.password
    }

    return await fetch(apiUrl + "/login",
        {
                body: body,
                method: "POST",
                headers: {
                    'Content-type': 'application/json'}
        })
}

async function register(user) {
    const body = {
        email: user?.email,
        password: user?.password,
        username: user?.username
    }
    return await fetch(apiUrl + "/signup",{
        body: body,
        method: "POST",
        headers: {
            'Content-type': 'application/json'}
    })
}

async function invite() {
    const body = {}
    return await fetch(apiUrl + "/invite",{
        body: body,
        method: "POST",
        headers: {
            'Content-type': 'application/json'}
    })
}

const LoginService = {
    login,
    register,
    invite
}

export default LoginService