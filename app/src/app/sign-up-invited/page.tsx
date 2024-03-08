import * as React from 'react';
import SignInForm from "@/components/SignInForm";

export default function SignUpInvited() {
  return (
      <>
          <SignInForm type={"sign-up-invited"} userType={"invited"}/>
      </>
  );
}
