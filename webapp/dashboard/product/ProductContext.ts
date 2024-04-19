import {createContext} from "react";
import {ProductListItem} from "./api";



export interface ProductContextValue {
    productId?: string,

    setProductId(userId: string): void

    loadProducts(): Promise<void>

    products: ProductListItem[]

    isNewProduct: boolean

    setIsNewProduct(isNewUser: boolean): void
}


const ProductContext =
    createContext<ProductContextValue>({
        setProductId(userId: string): void {
        },
        async loadProducts(): Promise<void> {
        },
        products: [],
        isNewProduct: false,
        setIsNewProduct(isNewUser: boolean) {
        },
    })

export default ProductContext
