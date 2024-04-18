import {InputText} from "primereact/inputtext";
import {Password} from "primereact/password";
import {Button} from "primereact/button";
import * as React from "react";
import {ChangeEvent, FormEvent, useContext, useState} from "react";
import UserContext from "./UserContext";
import {Message} from "primereact/message";

async function apiPostEntity(user: CreateUserDto) {
    const res = await fetch("/api/user", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(user)
    })

    if (!res.ok) {
        if (res.status === 400) {
            throw new FormError(await res.json())
        }

        throw new Error('Failed to fetch data')
    }

    return res.json()
}

class FormError {
    errors: {[key:string]:string[]}

    constructor(errors: any) {
        this.errors = errors;
    }
}

interface CreateUserDto {
    email: string
    password: string
    address: string
}

function NewCreateUserDto(): CreateUserDto {
    return {email: "", password: "", address: ""}
}

export default function CreateUserForm() {

    const {setIsNewUser, loadUsers} = useContext(UserContext)
    const [createUser, setCreateUser] = useState<CreateUserDto>(NewCreateUserDto())
    const [formError, setFormError] = useState<FormError>({errors:{}})

    const saveUser = async function (event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        try {
            await apiPostEntity(createUser)
            setIsNewUser(false)
            await loadUsers()
        } catch (e) {
            setFormError(e)
        }

    }

    const updateUserField = function (e: ChangeEvent<HTMLInputElement>) {
        setCreateUser((prevUser: CreateUserDto) => ({...prevUser, ...{[e.target.name]: e.target.value}}))
    }

    const getFieldError = function(name:string) {
        return formError.errors[name] && formError.errors[name].map(v=><Message severity="error" text={v} />)
    }

    return <form onSubmit={saveUser}>
        <div className="p-fluid">
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="email">Email</label>
                <InputText id="email" name="email" type="email"
                           value={createUser.email}
                           onChange={updateUserField}/>
                {getFieldError("email")}
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="password">Password</label>
                <Password id="password"  name="password" type="passsword" autoComplete="new-password"
                          value={createUser.password}
                          onChange={updateUserField}/>
                {getFieldError("password")}
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="address">Address</label>
                <InputText id="address" name="address" type="text" value={createUser.address}
                           onChange={updateUserField}/>
                {getFieldError("address")}
            </div>
        </div>
        <Button type="submit">Save</Button>
    </form>
}