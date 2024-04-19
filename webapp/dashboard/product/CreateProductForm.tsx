import {InputText} from "primereact/inputtext";
import {Button} from "primereact/button";
import * as React from "react";
import {ChangeEvent, FormEvent, useContext, useEffect, useState} from "react";
import {Message} from "primereact/message";
import {apiCreateProduct, CreateProductDto, FormError, UpdateProductDto} from "./api";
import ProductContext from "./ProductContext";
import {Dropdown} from "primereact/dropdown";
import {apiGetUsers, UserListItem} from "../userApi";
import {InputNumber, InputNumberValueChangeEvent} from "primereact/inputnumber";


function NewCreateProductDto(): CreateProductDto {
    return {title: "", brand: "", price: 0, userId: ""}
}

export default function CreateProductForm() {

    const {setIsNewProduct, loadProducts} = useContext(ProductContext)
    const [createProduct, setCreateProduct] = useState<CreateProductDto>(NewCreateProductDto())
    const [formError, setFormError] = useState<FormError>({errors: {}})
    const [users, setUsers] = useState<UserListItem[]>([])


    useEffect(function () {
        apiGetUsers().then(res => setUsers(res))
    }, [])

    const saveUser = async function (event: FormEvent<HTMLFormElement>) {
        event.preventDefault()
        try {
            await apiCreateProduct(createProduct)
            setIsNewProduct(false)
            await loadProducts()
        } catch (e) {
            setFormError(e)
        }

    }

    const updateField = function (e: ChangeEvent<HTMLInputElement>) {
        setCreateProduct((prev: CreateProductDto) => ({...prev, ...{[e.target.name]: e.target.value}}))
    }

    const updateNumberField = function (e: InputNumberValueChangeEvent) {
        const origEvent = e.originalEvent as ChangeEvent<HTMLInputElement>
        setCreateProduct((prevUser: CreateProductDto) => ({...prevUser, ...{[origEvent.target.name]: e.value}}))
    }


    const getFieldError = function (name: string) {
        return formError.errors[name] && formError.errors[name].map((v,i) => <Message key={i} severity="error" text={v}/>)
    }

    return <form onSubmit={saveUser}>
        <div className="p-fluid">
            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="title">User</label>
                <Dropdown value={createProduct.userId}
                          name="userId"
                          onChange={(e) => setCreateProduct(prev => ({...prev, ...{userId: e.value}}))} options={users}
                          optionLabel="email"
                          optionValue="id"
                          placeholder="Select a User" className="w-full md:w-14rem"/>
                {getFieldError("userId")}
            </div>

            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="title">Title</label>
                <InputText id="title" name="title" type="text"
                           value={createProduct.title}
                           onChange={updateField}/>
                {getFieldError("title")}
            </div>

            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="brand">Brand</label>
                <InputText id="brand" name="brand" type="text"
                           value={createProduct.brand}
                           onChange={updateField}/>
                {getFieldError("brand")}
            </div>

            <div className="field flex flex-column align-items-start gap-2">
                <label htmlFor="price">Price</label>
                <InputNumber id="price" name="price"
                             useGrouping={false}
                           value={createProduct.price}
                           onValueChange={updateNumberField}/>
                {getFieldError("price")}
            </div>

        </div>
        <Button type="submit">Save</Button>
    </form>
}