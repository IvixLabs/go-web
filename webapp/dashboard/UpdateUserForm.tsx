import {InputText} from "primereact/inputtext";
import {Password} from "primereact/password";
import {Button} from "primereact/button";
import * as React from "react";
import {ChangeEvent, FormEvent, useContext, useEffect, useState} from "react";
import UserContext from "./UserContext";
import {Message} from "primereact/message";
import {apiGetUser, apiUpdateEntity, FormError, UpdateUserDto} from "./userApi";


function NewUpdateUserDto(): UpdateUserDto {
    return {id: "", password: "", address: "", email: ""}
}

export default function UpdateUserForm() {

    const {userId, setUserId, loadUsers} = useContext(UserContext)
    const [updateUser, setUpdateUser] = useState<UpdateUserDto>(NewUpdateUserDto())
    const [formError, setFormError] = useState<FormError>({errors: {}})

    useEffect(function () {
        if (userId === undefined) {
            setUpdateUser(NewUpdateUserDto())
        } else {
            apiGetUser(userId).then(function (obj) {
                setUpdateUser({...obj, ...{password: ""}})
            })
        }
    }, [userId])

    const saveUser = async function (event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        try {
            await apiUpdateEntity(updateUser)
            setUserId(undefined)
            await loadUsers()
        } catch (e) {
            setFormError(e)
        }
    }

    const updateUserField = function (e: ChangeEvent<HTMLInputElement>) {
        setUpdateUser((prevUser: UpdateUserDto) => ({...prevUser, ...{[e.target.name]: e.target.value}}))
    }

    const getFieldError = function (name: string) {
        return formError.errors[name] && formError.errors[name].map(v => <Message severity="error" text={v}/>)
    }

    return <form onSubmit={saveUser}>
        <div className="p-fluid">
            <div className="field flex flex-column align-items-start gap-2">
                <span>{updateUser.email}</span>
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="password">Password</label>
                <Password id="password" name="password" type="passsword" autoComplete="new-password"
                          value={updateUser.password}
                          onChange={updateUserField}/>
                {getFieldError("password")}
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="address">Address</label>
                <InputText id="address" name="address" type="text" value={updateUser.address}
                           onChange={updateUserField}/>
                {getFieldError("address")}
            </div>
        </div>
        <Button type="submit">Save</Button>
    </form>
}