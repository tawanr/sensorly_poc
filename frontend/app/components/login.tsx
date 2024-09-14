"use client"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { loginUser } from "../lib/auth"

export default function LoginForm() {
  const loginUserBtn = async () => {
    const email = document.getElementById("email")!.value;
    const password = document.getElementById("password")!.value;
    if (!email || !password) {
      return;
    }
    let token = "";
    try {
      token = await loginUser(email, password);
    } catch (err) {
      console.error(err);
      return;
    }
    if (token) {
      localStorage.setItem("token", token);
    }
  };

  return (
    <div className="flex h-full w-full flex-col items-center justify-center gap-4 p-4 sm:px-6 sm:py-24 md:py-32 lg:px-8">
      <Card className="w-full max-w-sm">
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
          <CardDescription>
            Enter your email below to login to your account.
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <div className="grid gap-2">
            <Label htmlFor="email">Email</Label>
            <Input id="email" type="email" placeholder="m@example.com" required />
          </div>
          <div className="grid gap-2">
            <Label htmlFor="password">Password</Label>
            <Input id="password" type="password" required />
          </div>
        </CardContent>
        <CardFooter>
          <Button className="w-full" onClick={loginUserBtn}>Sign in</Button>
        </CardFooter>
      </Card>
    </div>
  )
}
