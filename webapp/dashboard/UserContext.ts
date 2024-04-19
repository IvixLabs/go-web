import {createContext} from "react";
import {UserListItem} from "./userApi";



export interface UserContextValue {
    userId?: string,

    setUserId(userId: string): void

    loadUsers(): Promise<void>

    users: UserListItem[]

    isNewUser: boolean

    setIsNewUser(isNewUser: boolean): void

    apiDeleteUser(userId: string): Promise<void>
}


const UserContext =
    createContext<UserContextValue>({
        setUserId(userId: string): void {
        },
        async loadUsers(): Promise<void> {
        },
        users: [],
        isNewUser: false,
        setIsNewUser(isNewUser: boolean) {
        },
        async apiDeleteUser(userId: string): Promise<void> {
        }
    })

export default UserContext
