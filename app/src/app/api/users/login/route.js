import {apiUrl} from "@/constants/RoutesUrl";
import {NextRequest, NextResponse} from "next/server";

export async function POST(request){
    const reqBody = await request.json()
    const {email, password} = reqBody;
    console.log(reqBody);

    const res = await fetch(apiUrl + "/login",
        {
            body: JSON.stringify({email, password}),
            method: "POST",
            headers: {
                'Content-type': 'application/json'}
        })

    const loginData = await res?.json()
    const token = loginData?.token

    const response = NextResponse.json({
        message: "Login successful",
        success: true,
    })

    response.cookies.set("token", token, {
        httpOnly: true,
    })
    return response;
}