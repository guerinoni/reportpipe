"use client";

import * as React from 'react';
import Sheet from '@mui/joy/Sheet';
import Typography from '@mui/joy/Typography';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Input from '@mui/joy/Input';
import Button from '@mui/joy/Button';
import Link from '@mui/joy/Link';
import SignInForm from "@/components/signInForm";

export default function SignIn() {

    return (
        <>
            <SignInForm type={"sign-in"}/>
        </>
    );
}
