const userApiEndpoint = "/api/users"

export interface UserListItem {
    id: string
    email: string
    address: string
}


export interface CreateUserDto {
    email: string
    password: string
    address: string
}

export class FormError {
    errors: { [key: string]: string[] }

    constructor(errors: any) {
        this.errors = errors;
    }
}

export interface UpdateUserDto {
    id: string
    email: string
    password: string
    address: string
}

export async function apiGetUsers(): Promise<UserListItem[]> {
    const res = await fetch(userApiEndpoint + "/")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export async function apiDeleteUser(userId: string) {
    await fetch(userApiEndpoint + "/" + userId, {
        method: "DELETE",
        headers: {"Content-Type": "application/json"},
    })
}

export async function apiCreateEntity(user: CreateUserDto) {
    const res = await fetch(userApiEndpoint + "/", {
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

export async function apiGetUser(userId: string): Promise<UpdateUserDto> {
    const res = await fetch(userApiEndpoint + "/" + userId)
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export async function apiUpdateEntity(user: UpdateUserDto) {
    const res = await fetch(userApiEndpoint + "/" + user.id, {
        method: "PUT",
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