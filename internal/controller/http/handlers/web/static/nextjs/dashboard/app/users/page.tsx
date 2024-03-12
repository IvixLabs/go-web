"use client"

import Image from "next/image";
import {useState} from "react";

async function getUsers() {
    const res = await fetch("http://localhost:8080/api/user/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default function Home() {

    const [users, setUsers] = useState<any[]>([])
    const [count, setCount] = useState(0)

    const userList = users.map(function (user) {
        return <li key={user.id}>
            <p>{user.id}</p>
            <p>{user.email}</p>
            <p>{user.adress}</p>
        </li>
    })

    return (
        <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <h1>{count}</h1>
            <button onClick={function () {
                setCount(count + 1)
            }}>Click
            </button>

            <button onClick={async function () {
                setUsers(await getUsers())
            }}>Fetch users
            </button>

            <ul>{userList}</ul>
        </main>
    );
}
