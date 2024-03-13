import {apiUrl} from "../constants/RoutesUrl";

//import { serialize } from 'cookie'
//import {encrypt} from "next/dist/server/app-render/action-encryption-utils";
//import {handleLoginCookies} from "@/app/actions";
//import {cookies} from "next/headers";

/*
export function cookieHandler(sessionData) {
    const encryptedSessionData = encrypt(sessionData)

    cookies().set('session', encryptedSessionData, {
        httpOnly: true,
        secure: process.env.NODE_ENV === 'production',
        maxAge: 60 * 60 * 24 * 7, // One week
        path: '/',
    })
}
*/
async function login(user, userType) {
    const body = {
        email: user?.email,
        password: user?.password
    }

    const response = await fetch(apiUrl + "/login",
        {
                body: JSON.stringify(body),
                method: "POST",
                headers: {
                    'Content-type': 'application/json'}
        })

    //  await cookieHandler(response)
}

async function register(user) {
    const body = {
        email: user?.email,
        password: user?.password,
        username: user?.username
    }
    return await fetch(apiUrl + "/signup",{
        body: JSON.stringify(body),
        method: "POST",
        headers: {
            'Content-type': 'application/json'}
    })
}

async function invite() {
    const body = {}
    return await fetch(apiUrl + "/invite",{
        body: JSON.stringify(body),
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