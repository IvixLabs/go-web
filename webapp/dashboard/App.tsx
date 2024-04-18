import * as React from "react";
import {createContext, ReactNode, useCallback, useContext, useEffect, useRef, useState} from "react";

import {Button} from 'primereact/button';
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";
import {confirmDialog} from "primereact/confirmdialog";
import {InputText} from "primereact/inputtext";
import {Card} from "primereact/card";
import {Message} from "primereact/message";
import {NavLink, Outlet, useLocation, useNavigate} from "react-router-dom";


async function apiGetProperties() {
    const res = await fetch("/api/property/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

interface ListProperty {
    id: string
    name: string
}

interface CreateProperty {
    name: string
}

class FormError {
    errors: any

    constructor(errors: any) {
        this.errors = errors;
    }
}

async function apiPutProperty(property: CreateProperty) {
    const res = await fetch("/api/property/create", {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(property)
    })
    if (!res.ok) {
        if (res.status === 400) {
            throw new FormError(await res.json())
        }

        throw new Error('Failed to fetch data')
    }

    return res.json()
}

async function apiDeleteProperty(id: string) {
    const res = await fetch("/api/property/delete?" + new URLSearchParams({id: id}), {
        method: "DELETE",
        headers: {"Content-Type": "application/json"},
    })
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}


function createEmptyCreateProperty(): CreateProperty {
    return {
        name: ""
    }
}

interface FormInputProps {
    name: string
    value: string
    errors: { [key: string]: string[] }
    disabled?: boolean

    updateForm(e: React.ChangeEvent<HTMLInputElement>): Promise<void>
}

function FormInput({errors, updateForm, name, value, disabled = undefined}: FormInputProps) {

    const [localValue, setLocalValue] = useState<string>(value)
    const lastEvent = useRef<React.ChangeEvent<HTMLInputElement>>()
    const lastTimeout = useRef<ReturnType<typeof setTimeout>>()

    useEffect(() => {
        if (localValue != value) {
            setLocalValue(value)
        }
    }, [value])

    console.log("FormInput " + name + " RENDER")
    return <div className="flex flex-column gap-2">
        <InputText value={localValue}
                   disabled={disabled}
                   className={(errors.name ? "p-invalid" : '')}
                   name={name}
                   onChange={e => {
                       setLocalValue(e.target.value);
                       lastEvent.current = e
                       if (lastTimeout.current) {
                           clearTimeout(lastTimeout.current)
                           lastTimeout.current = undefined
                       }
                       lastTimeout.current = setTimeout(async () => {
                           try {
                               if (!lastEvent.current) {
                                   return
                               }
                               const event = lastEvent.current
                               lastEvent.current = undefined
                               return updateForm(event)
                           } finally {
                               lastTimeout.current = undefined
                           }
                       }, 300)

                   }}></InputText>
        {errors.name && fieldErrors(errors.name)}
    </div>
}

const fieldErrors = function (errors: string[]) {
    return errors.map((error, index) => <Message severity="error" text={error} key={index}/>)
}

interface PropertyFormProps {
    onSuccess(): Promise<void>
}

function PropertyForm({onSuccess}: PropertyFormProps) {
    const [createPropertyForm, setCreatePropertyForm] = useState<CreateProperty>(createEmptyCreateProperty)
    const [errors, setErrors] = useState<{ [key: string]: string[] }>({})

    const updateForm = useCallback(async function (e: React.ChangeEvent<HTMLInputElement>) {
        const fieldName = e.target.name
        const fieldValue = e.target.value

        setCreatePropertyForm(curCreatePropertyForm =>
            ({...curCreatePropertyForm, ...{[fieldName]: fieldValue}}))

        setErrors(curErrors => {
            if (curErrors[fieldName]) {
                delete curErrors[fieldName]
            }

            return curErrors
        })

    }, [])

    const addProperty = useCallback(async function (e: React.FormEvent) {
        e.preventDefault()
        try {
            await apiPutProperty(createPropertyForm)
            setCreatePropertyForm(createEmptyCreateProperty())
            return onSuccess()
        } catch (e) {
            if (e instanceof FormError) {
                setErrors(e.errors)
            }
        }
    }, [onSuccess, createPropertyForm])

    console.log("Form RENDER")

    return <form onSubmit={addProperty}>
        <fieldset style={{border: "none", padding: 0, margin: 0}}
                  className="flex flex-column gap-2">
            <FormInput name="name" errors={errors} updateForm={updateForm} value={createPropertyForm.name}/>
            <div>
                <Button type="submit">Save</Button>
            </div>
        </fieldset>
    </form>
}

const PropertyFormMemo = React.memo(PropertyForm)

interface PropertyTableProps {
    // properties: ListProperty[]
    // deleteProperty(id: string): Promise<void>
}

function PropertyTable({}: PropertyTableProps) {

    const {properties, loadProperties} = useContext(PropertiesContext)

    useEffect(() => {
        loadProperties().then()
    }, [loadProperties])

    const deleteProperty = useCallback(async function (id: string) {
        await apiDeleteProperty(id)
        await loadProperties()
    }, [loadProperties])

    const deleteTemplate = function (value: ListProperty) {
        return <Button severity="danger"
                       onClick={async () => deleteProperty(value.id)}>Delete</Button>
    }

    return <DataTable value={properties} size="small">
        <Column field="id" header="Id"/>
        <Column field="name" header="Name"/>
        <Column field="id" body={deleteTemplate}/>
    </DataTable>
}

interface PropertiesContextValue {
    properties: ListProperty[]

    setProperties(properties: ListProperty[]): void

    loadProperties(): Promise<void>
}

const PropertiesContext = createContext<PropertiesContextValue>(undefined)


interface AppProvidersProps {
    children: ReactNode
}

function AppProviders({children}: AppProvidersProps) {
    const [properties, setProperties] = useState<ListProperty[]>([])

    const loadProperties = useCallback(async () => {
        setProperties(await apiGetProperties())
    }, [])

    return <PropertiesContext.Provider value={{properties, setProperties, loadProperties}}>
        {children}
    </PropertiesContext.Provider>
}

function PropertyPage() {
    const {loadProperties} = useContext(PropertiesContext)

    return <Card>
        <div className="flex flex-column gap-2">
            <PropertyTable/>

            <div className="flex flex-row gap-2">
                <Button onClick={loadProperties}>Load properties</Button>
            </div>

            <div>
                <PropertyFormMemo onSuccess={loadProperties}/>
            </div>
        </div>
    </Card>
}

export function RedirectToUsers() {
    const navigate = useNavigate();

    useEffect(() => {
        navigate("/dashboard/users")
    }, []);

    return <></>
}

export default function App() {

    const navigate = useNavigate()
    const location = useLocation()

    const confirm1 = function () {
        confirmDialog({
            message: 'Are you sure you want to proceed?',
            header: 'Confirmation',
            icon: 'pi pi-exclamation-triangle',
            defaultFocus: 'accept',
        })
    }

    const navLinkClass = function (props: { isActive: boolean }) {
        const {isActive} = props
        const baseClass = " no-underline line-height-3 cursor-pointer"

        if (isActive) {
            return "text-900 " + baseClass
        }
        return "text-500 " + baseClass
    }

    return (
        <div>
            <AppProviders>
                <div className="flex flex-column gap-2">
                    <ul className="list-none p-0 m-0 flex align-items-center font-medium mb-3">
                        <li className="p-2">
                            <NavLink to="users" className={navLinkClass}>Users</NavLink>
                        </li>
                        <li className="p-2">
                            <NavLink to="products" className={navLinkClass}>Products</NavLink>
                        </li>
                    </ul>
                    <Outlet/>
                </div>
            </AppProviders>

        </div>


    )
}