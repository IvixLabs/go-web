import * as React from "react";
import {useContext, useEffect, useState} from "react";
import {Card} from "primereact/card";
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";
import UpdateUserForm from "./UpdateUserForm";
import UserContext, {UserListItem} from "./UserContext";
import {Dialog} from "primereact/dialog";
import CreateUserForm from "./CreateUserForm";
import {Button} from "primereact/button";
import {ConfirmDialog} from "primereact/confirmdialog";

async function apiGetUsers(): Promise<UserListItem[]> {
    const res = await fetch("/api/user/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

async function apiDeleteUser(userId: string) {
    const res = await fetch("/api/user?userId=" + userId, {
        method: "DELETE",
        headers: {"Content-Type": "application/json"},
    })
}


function UserTable() {

    const {setUserId, users, loadUsers, apiDeleteUser} = useContext(UserContext)

    const [deleteUserId, setDeleteUserId] = useState<string | undefined>()


    useEffect(() => {
        loadUsers().then()
    }, [])

    const idClickHandler = function (e: React.MouseEvent<HTMLElement>, data: UserListItem) {
        e.preventDefault()
        setUserId(data.id)
    }

    const deleteClickHandler = function (userId: string) {
        setDeleteUserId(userId)
        console.log(userId)
    }

    const acceptDeletion = async function () {
        await apiDeleteUser(deleteUserId)
        await loadUsers()
    }

    return <>
        <DataTable value={users} size="small">
            <Column field="id" header="Id"
                    body={(data: UserListItem) => <a href="" onClick={e => idClickHandler(e, data)}> {data.id}</a>}/>
            <Column field="email" header="Email"/>
            <Column field="address" header="Address"/>
            <Column field="id"
                    body={(data: UserListItem) => <Button severity="danger" onClick={() => deleteClickHandler(data.id)}>Delete</Button>}/>
        </DataTable>
        <ConfirmDialog visible={deleteUserId !== undefined} onHide={() => setDeleteUserId(undefined)}
                       message="Are you sure you want to proceed?"
                       header="Confirmation" icon="pi pi-exclamation-triangle" accept={acceptDeletion}/>
    </>
}


export default function UserListPage() {

    const [isNewUser, setIsNewUser] = useState<boolean>(false)
    const [userId, setUserId] = useState<string | undefined>()
    const [users, setUsers] = useState<UserListItem[]>([])
    const loadUsers = async () => {
        setUsers(await apiGetUsers())
    }

    const hideDialogHandler = function () {
        setUserId(undefined)
        loadUsers().then()
    }

    const hideCreateDialogHandler = function () {
        setIsNewUser(false)
        loadUsers().then()
    }

    return <>
        <UserContext.Provider value={{userId, setUserId, loadUsers, users, isNewUser, setIsNewUser, apiDeleteUser}}>
            <Card>
                <Button onClick={() => setIsNewUser(true)}>Create new user</Button>
                <UserTable/>
            </Card>
            <Dialog header="User form" onHide={hideDialogHandler} visible={userId !== undefined}>
                <UpdateUserForm/>
            </Dialog>

            <Dialog header="User form" onHide={hideCreateDialogHandler} visible={isNewUser}>
                <CreateUserForm/>
            </Dialog>
        </UserContext.Provider>
    </>
}