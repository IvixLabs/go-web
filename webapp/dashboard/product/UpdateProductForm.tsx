import {InputText} from "primereact/inputtext";
import {Button} from "primereact/button";
import * as React from "react";
import {ChangeEvent, FormEvent, useContext, useEffect, useState} from "react";
import {Message} from "primereact/message";
import ProductContext from "./ProductContext";
import {apiGetProduct, apiUpdateProduct, FormError, UpdateProductDto} from "./api";
import {InputNumber, InputNumberChangeEvent, InputNumberValueChangeEvent} from "primereact/inputnumber";


function NewUpdateProductDto(): UpdateProductDto {
    return {id: "", brand: "", title: "", price: 0, userEmail: ""}
}

export default function UpdateProductForm() {

    const {productId, setProductId, loadProducts} = useContext(ProductContext)
    const [updateProduct, setUpdateProduct] = useState<UpdateProductDto>(NewUpdateProductDto())
    const [formError, setFormError] = useState<FormError>({errors: {}})

    useEffect(function () {
        if (productId === undefined) {
            setUpdateProduct(NewUpdateProductDto())
        } else {
            apiGetProduct(productId).then(function (obj) {
                setUpdateProduct(obj)
            })
        }
    }, [productId])

    const save = async function (event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        try {
            await apiUpdateProduct(updateProduct)
            setProductId(undefined)
            await loadProducts()
        } catch (e) {
            setFormError(e)
        }
    }

    const updateField = function (e: ChangeEvent<HTMLInputElement>) {
        setUpdateProduct((prevUser: UpdateProductDto) => ({...prevUser, ...{[e.target.name]: e.target.value}}))
    }

    const updateNumberField = function (e: InputNumberValueChangeEvent) {
        const origEvent = e.originalEvent as ChangeEvent<HTMLInputElement>
        setUpdateProduct((prevUser: UpdateProductDto) => ({...prevUser, ...{[origEvent.target.name]: e.value}}))
    }

    const getFieldError = function (name: string) {
        return formError.errors[name] && formError.errors[name].map((v, i) => <Message key={i} severity="error" text={v}/>)
    }

    return <form onSubmit={save}>
        <div className="p-fluid">
            <div className="field flex flex-column align-items-start gap-2">
                <span>{updateProduct.userEmail}</span>
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="title">Title</label>
                <InputText id="title" name="title" type="text" value={updateProduct.title} onChange={updateField}/>
                {getFieldError("title")}
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="brand">Brand</label>
                <InputText id="brand" name="brand" type="text" value={updateProduct.brand} onChange={updateField}/>
                {getFieldError("brand")}
            </div>
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="price">Price</label>
                <InputNumber id="price" name="price" value={updateProduct.price} useGrouping={false}
                           onValueChange={updateNumberField}/>
                {getFieldError("price")}
            </div>
        </div>
        <Button type="submit">Save</Button>
    </form>
}