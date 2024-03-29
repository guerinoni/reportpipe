"use client";
import * as React from 'react';
import Sheet from '@mui/joy/Sheet';
import Typography from '@mui/joy/Typography';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Input from '@mui/joy/Input';
import Button from '@mui/joy/Button';
import Link from '@mui/joy/Link';
import {useState} from "react";
import {useRouter} from "next/navigation";
import axios from "axios";
import { toast } from "react-hot-toast";

export default function SignInForm({type, userType}) {

    const [user, setUser] = useState({})
    const router = useRouter()

    const handleFormSubmit = async () => {
        if (type === 'sign-in') {
            try {
                const response = await axios.post("/api/users/login", user);
                console.log("Login success", response.data);
                toast.success("Login success");
                router.push("/home")
            } catch (error) {
                console.log("Login failed", error.message);
                toast.error(error.message);
            }
        } else {
            try {
                const response = await axios.post("/api/users/register", user);
                console.log("Sign up success", response.data);
                toast.success("Sign up success");
                router.push("/login")
            } catch (e) {
                console.log("E: ", e)
            }
        }
    }

    const handleEmail = (event) => {
        let tmpUser = {...user}
        tmpUser.email = event.target.value
        setUser(tmpUser)
    }

    const handlePswd = (event) => {
        let tmpUser = {...user}
        tmpUser.password = event.target.value
        setUser(tmpUser)
    }

    return (
        <Sheet
            sx={{
                display: 'flex',
                flexFlow: 'row nowrap',
                justifyContent: 'center',
                alignItems: 'center',
                minHeight: '100vh',
            }}
        >
            <Sheet
                sx={{
                    width: 300,
                    mx: 'auto',
                    my: 4,
                    py: 3,
                    px: 2,
                    display: 'flex',
                    flexDirection: 'column',
                    gap: 2,
                    borderRadius: 'sm',
                    boxShadow: 'md',
                }}
                variant="outlined"
            >
                <div>
                    <Typography level="h4" component="h1">
                        <strong>Welcome back 👋</strong>
                    </Typography>
                    <Typography level="body-sm">Sign in to continue.</Typography>
                </div>
                <FormControl id="email">
                    <FormLabel>Email</FormLabel>
                    <Input name="email" type="email" placeholder="johndoe@email.com" onChange={handleEmail}/>
                </FormControl>
                <FormControl id="password">
                    <FormLabel>Password</FormLabel>
                    <Input name="password" type="password" placeholder="password" onChange={handlePswd}/>
                </FormControl>
                <Button sx={{ mt: 1 }} onClick={handleFormSubmit}>Log in</Button>
                <Typography
                    endDecorator={<Link href="/sign-up">Sign up</Link>}
                    fontSize="sm"
                    sx={{ alignSelf: 'center' }}
                >
                    Don&apos;t have an account?
                </Typography>
            </Sheet>
        </Sheet>
    );
}
