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
import LoginService from "../services/LoginService";
import {redirect} from "next/navigation";


export default function InviteForm() {

    const [user, setUser] = useState({})

    const handleFormSubmit = async () => {
        try {
            await LoginService.invite(user)
            //redirect("/home")
        } catch (e) {
            console.log("E: ", e)
        }
    }

    const handleEmail = (event) => {
        let tmpUser = {...user}
        tmpUser.email = event.target.value
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
                        <strong>Invite user</strong>
                    </Typography>
                    <Typography level="body-sm">Send a login link</Typography>
                </div>
                <FormControl id="email">
                    <FormLabel>Email</FormLabel>
                    <Input name="email" type="email" placeholder="johndoe@email.com" onChange={handleEmail}/>
                </FormControl>
                <Button sx={{ mt: 1 }} onClick={handleFormSubmit}>Invite</Button>
            </Sheet>
        </Sheet>
    );
}
