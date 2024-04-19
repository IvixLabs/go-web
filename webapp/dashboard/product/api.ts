const apiEndpoint = "/api/products"

export interface ProductListItem {
    id: string
    title: string
    brand: string
    userEmail: string
    price: number
}


export interface CreateProductDto {
    title: string
    brand: string
    price: number
    userId: string
}

export class FormError {
    errors: { [key: string]: string[] }

    constructor(errors: any) {
        this.errors = errors;
    }
}

export interface UpdateProductDto {
    id: string
    title: string
    brand: string
    price: number
    userEmail: string
}

export async function apiGetProducts(): Promise<ProductListItem[]> {
    const res = await fetch(apiEndpoint + "/")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export async function apiDeleteProduct(productId: string) {
    await fetch(apiEndpoint + "/" + productId, {
        method: "DELETE",
        headers: {"Content-Type": "application/json"},
    })
}

export async function apiCreateProduct(user: CreateProductDto) {
    const res = await fetch(apiEndpoint + "/", {
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

export async function apiGetProduct(userId: string): Promise<UpdateProductDto> {
    const res = await fetch(apiEndpoint + "/" + userId)
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export async function apiUpdateProduct(product: UpdateProductDto) {
    const res = await fetch(apiEndpoint + "/" + product.id, {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(product)
    })

    if (!res.ok) {
        if (res.status === 400) {
            throw new FormError(await res.json())
        }

        throw new Error('Failed to fetch data')
    }

    return res.json()
}