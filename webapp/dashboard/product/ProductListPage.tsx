import * as React from "react";
import {useContext, useEffect, useState} from "react";
import {Card} from "primereact/card";
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";
import {apiDeleteProduct, apiGetProducts, ProductListItem} from "./api";
import {Button} from "primereact/button";
import {Dialog} from "primereact/dialog";
import UpdateProductForm from "./UpdateProductForm";
import ProductContext from "./ProductContext";
import CreateProductForm from "./CreateProductForm";
import {ConfirmDialog} from "primereact/confirmdialog";

function ProductTable() {

    const {setProductId, products, loadProducts} = useContext(ProductContext)

    const [deleteProductId, setDeleteProductId] = useState<string | undefined>()

    useEffect(() => {
        loadProducts().then()
    }, [])

    const idClickHandler = function (e: React.MouseEvent<HTMLElement>, data: ProductListItem) {
        e.preventDefault()
        setProductId(data.id)
    }

    const deleteClickHandler = function (productId: string) {
        setDeleteProductId(productId)
    }

    const acceptDeletion = async function () {
        await apiDeleteProduct(deleteProductId)
        await loadProducts()
    }

    return <>
        <DataTable value={products} size="small">
            <Column field="id" header="Id"
                    body={(data: ProductListItem) => <a href="" onClick={e => idClickHandler(e, data)}> {data.id}</a>}/>
            <Column field="userEmail" header="User"/>
            <Column field="title" header="Title"/>
            <Column field="brand" header="Brand"/>
            <Column field="price" header="Price"/>
            <Column field="id"
                    body={(data: ProductListItem) => <Button severity="danger"
                                                             onClick={() => deleteClickHandler(data.id)}>Delete</Button>}/>
        </DataTable>
        <ConfirmDialog visible={deleteProductId !== undefined} onHide={() => setDeleteProductId(undefined)}
                       message="Are you sure you want to proceed?"
                       header="Confirmation" icon="pi pi-exclamation-triangle" accept={acceptDeletion}/>
    </>
}

export default function ProductListPage() {
    const [isNewProduct, setIsNewProduct] = useState<boolean>(false)
    const [products, setProducts] = useState<ProductListItem[]>([])
    const [productId, setProductId] = useState<string | undefined>()

    const loadProducts = async () => {
        setProducts(await apiGetProducts())
    }

    const hideCreateDialogHandler = function () {
        setIsNewProduct(false)
        loadProducts().then()
    }

    const hideDialogHandler = function () {
        setProductId(undefined)
        loadProducts().then()
    }

    return <ProductContext.Provider
        value={{isNewProduct, setIsNewProduct, loadProducts, productId, setProductId, products}}>
        <Card>
            <Button className="mb-2" onClick={() => setIsNewProduct(true)}>Create new product</Button>
            <ProductTable/>
        </Card>
        <Dialog header="User form" onHide={hideDialogHandler} visible={productId !== undefined}>
            <UpdateProductForm/>
        </Dialog>
        <Dialog header="Product form" onHide={hideCreateDialogHandler} visible={isNewProduct}>
            <CreateProductForm/>
        </Dialog>
    </ProductContext.Provider>
}