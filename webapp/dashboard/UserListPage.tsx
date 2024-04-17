import * as React from "react";
import {useEffect, useState} from "react";
import {Card} from "primereact/card";
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";

async function apiGetUsers(): Promise<UserListItem[]> {
    const res = await fetch("/api/user/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

interface UserListItem {
    id: string
    email: string
    address: string
}

function UserTable() {

    const [users, setUsers] = useState<UserListItem[]>([])

    const loadUsers = async () => {
        setUsers(await apiGetUsers())
    }

    useEffect(() => {
        loadUsers().then()
    }, [])


    return <DataTable value={users} size="small">
        <Column field="id" header="Id"/>
        <Column field="email" header="Email"/>
        <Column field="address" header="Address"/>
    </DataTable>
}

export default function UserListPage() {
    return <Card>
        <div className="flex flex-column gap-2">
            <UserTable/>
        </div>
    </Card>
}