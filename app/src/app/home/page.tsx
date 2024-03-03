import * as React from 'react';
import Sheet from '@mui/joy/Sheet';
import Typography from '@mui/joy/Typography';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Input from '@mui/joy/Input';
import Button from '@mui/joy/Button';
import Link from '@mui/joy/Link';
import Dashboard from "@/app/dashboard/dashboard";

export default function Home() {
  return (

      //add checks if user is logged in session, if not, redirect to login - missing API to check user exist / logged

      <Dashboard />
  );
}
