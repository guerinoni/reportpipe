import {apiUrl} from "@/constants/RoutesUrl";
import {NextResponse} from "next/server";

export async function POST(request){
    const reqBody = await request.json()
    const {email, password} = reqBody;
    console.log(reqBody);

    const res = await fetch(apiUrl + "/signup",
        {
            body: JSON.stringify({username: email, email, password}),
            method: "POST",
            headers: {
                'Content-type': 'application/json'}
        })

    return NextResponse.json({
        message: "Sign up successful",
        success: true,
    });
}