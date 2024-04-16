import * as React from "react";
import {createContext, ReactNode, useCallback, useContext, useEffect, useRef, useState} from "react";

import {Button} from 'primereact/button';
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";
import {confirmDialog, ConfirmDialog} from "primereact/confirmdialog";
import {InputText} from "primereact/inputtext";
import {Card} from "primereact/card";
import {Message} from "primereact/message";
import {Dropdown} from "primereact/dropdown";


async function apiPutEntity(entity: ListEntity) {
    const res = await fetch("/api/entity/save", {
        method: "PUT",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify(entity)
    })

    if (!res.ok) {
        if (res.status === 400) {
            throw new FormError(await res.json())
        }

        throw new Error('Failed to fetch data')
    }

    return res.json()
}

async function apiGetProperties() {
    const res = await fetch("/api/property/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

async function apiGetEntities(req: {
    id?: string,
    first?: string,
    rows?: string,
    sortField?: string,
    sortFieldType?: string,
    sortOrder?: string
}): Promise<EntitiesRes> {

    const params = new URLSearchParams()

    Object.keys(req).forEach(k => {
        if ((req as { [Key: string]: any })[k]) {
            params.set(k, (req as { [Key: string]: any })[k])
        }
    })

    const res = await fetch("/api/entity/list?" + params)
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

interface EntitiesRes {
    entities: ListEntity[]
    properties: { name: string, type: string }[]
    total: number
}


interface EntityProperty {
    name: string
    value: string
    disabled: boolean
    type: string
}

interface ListEntity {
    id?: string
    properties: { [key: string]: EntityProperty }
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

    console.log("PropertyTable RENDER")
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

interface EntityFormProps {
    entity: ListEntity

    updateEntity(entity: ListEntity): void

    loadEntities(): Promise<void>
}

function EntityForm({entity, updateEntity, loadEntities}: EntityFormProps) {

    const [newPropName, setNewPropName] = useState<string>("")

    const propertyTypes = [
        {name: 'String', code: 'string'},
        {name: 'Integer', code: 'int'},
        {name: 'Float', code: 'float'},
    ];

    const addProperty = useCallback(function () {
        if (newPropName.length > 0) {
            entity.properties[newPropName] = {
                name: newPropName, value: "", disabled: false, type: "string"
            }
            updateEntity(entity)
            setNewPropName("")
        }
    }, [entity, newPropName, setNewPropName])

    const saveEntity = useCallback(async function (e: React.FormEvent) {
        e.preventDefault()
        await apiPutEntity(entity)
        return loadEntities()
    }, [entity])

    console.log("EntityForm RENDER")
    return <form onSubmit={saveEntity}>
        <p>{entity.id}</p>
        <fieldset style={{border: "none", padding: 0, margin: 0}} className="flex flex-column gap-2">
            {Object.keys(entity.properties).map((prop: string) =>
                <div key={prop} className="flex flex-column gap-2">
                    <label htmlFor="username">{prop}</label>
                    <FormInput name={prop} errors={{}}
                               disabled={entity.properties[prop].disabled}
                               updateForm={async (e: React.ChangeEvent<HTMLInputElement>) => {
                                   entity.properties[prop].value = e.target.value
                                   updateEntity(entity)
                               }}
                               value={entity.properties[prop].value ?? ""}/>
                    <Dropdown value={entity.properties[prop].type ?? "string"}
                              onChange={(e) => {
                                  entity.properties[prop].type = e.value
                                  updateEntity(entity)
                              }}
                              options={propertyTypes}
                              optionLabel="name"
                              optionValue="code"
                              placeholder="Select a Type" className="w-full md:w-14rem"/>

                    {!entity.properties[prop].disabled && <Button onClick={() => {
                        entity.properties[prop].disabled = true;
                        updateEntity(entity)
                    }}>disable</Button>}
                    {entity.properties[prop].disabled && <Button onClick={() => {
                        entity.properties[prop].disabled = false;
                        updateEntity(entity)
                    }}>enable</Button>}
                </div>)
            }
            <div>
                <FormInput name="newPropName" errors={{}}
                           updateForm={async (e: React.ChangeEvent<HTMLInputElement>) => {
                               setNewPropName(e.target.value)
                           }}
                           value={newPropName}/>
                <Button type="button" onClick={addProperty}>add</Button>
            </div>
            <div>
                <Button type="submit">Save</Button>
            </div>
        </fieldset>
    </form>
}

const EntityFormMemo = React.memo(EntityForm)


interface EntityTableProps {
    entitiesRes: EntitiesRes

    setEntityRes(res: EntitiesRes): void

    firstRows: number
    rows: number
    sortField: string
    sortFieldType: string
    sortOrder: number

    updatePage(firstRows: number, rows: number, sortField?: string, sortFieldType?: string, sortOrder?: number): Promise<void>
}

function EntityTable({
                         entitiesRes,
                         setEntityRes,
                         firstRows,
                         rows,
                         sortField,
                         sortFieldType,
                         sortOrder,
                         updatePage
                     }: EntityTableProps) {

    const editTemplate = function (value: ListEntity) {
        return <Button severity="warning"
                       onClick={async () => setEntityRes(await apiGetEntities({id: value.id}))}>Edit</Button>
    }

    useEffect(function () {
        updatePage(firstRows, rows, sortField, sortFieldType, sortOrder).then()
    }, [updatePage])

    console.log("EntityTable RENDER")
    return <DataTable value={entitiesRes.entities}
                      first={firstRows}
                      rows={rows}
                      lazy={true}
                      sortField={sortField}
                      totalRecords={entitiesRes.total}
                      onPage={async (e) => {

                          const foundProp = entitiesRes.properties
                              .find(p => p.name == sortField)
                          let sortFieldType = undefined
                          if (foundProp) {
                              sortFieldType = foundProp.type
                          }

                          return updatePage(e.first, e.rows, sortField, sortFieldType, sortOrder)
                      }}
                      onSort={async (e) => {

                          const sortFieldName = e.sortField.split(".")[1]

                          const foundProp = entitiesRes.properties
                              .find(p => p.name == sortFieldName)
                          let sortFieldType = undefined
                          if (foundProp) {
                              sortFieldType = foundProp.type
                          }

                          return updatePage(0, rows, sortFieldName, sortFieldType, e.sortOrder)
                      }}
                      paginator
                      rowsPerPageOptions={[10, 20, 30, 40, 50]} size="small">
        <Column field="id" header="Id"/>
        {entitiesRes.properties.map(prop => <Column key={prop.name} sortable

                                                    field={"properties." + prop.name + ".value"} header={prop.name}/>)}
        <Column field="id" body={editTemplate}/>
    </DataTable>

}

const EntityTableMemo = React.memo(EntityTable)

function EntityPage() {

    const [entitiesRes, setEntitiesRes] = useState<EntitiesRes>({entities: [], properties: [], total: 0})
    const [entityRes, setEntityRes] = useState<EntitiesRes>(undefined)

    const [firstRows, setFirstRows] = useState<number>(0)
    const [rows, setRows] = useState<number>(10)
    const [sortField, setSortField] = useState<string>(undefined)
    const [sortFieldType, setSortFieldType] = useState<string>(undefined)
    const [sortOrder, setSortOrder] = useState<number>(1)

    const loadEntities = useCallback(async function () {
        setEntitiesRes(await apiGetEntities({
            first: firstRows.toString(),
            rows: rows.toString(),
            sortField: sortField,
            sortFieldType: sortFieldType,
            sortOrder: sortOrder.toString()
        }))

    }, [setEntitiesRes, firstRows, rows, sortField, sortFieldType, sortOrder])

    const updatePage = useCallback(async function (firstRows: number, rows: number, sortField: string, sortFieldType: string, sortOrder: number) {
        setFirstRows(firstRows)
        setRows(rows)
        setSortField(sortField)
        setSortOrder(sortOrder)
        setSortFieldType(sortFieldType)
        // const res = await apiGetEntities({first: firstRows.toString(), rows: rows.toString()})
        // setEntitiesRes(res)
        await loadEntities()

    }, [setFirstRows, setRows, setEntitiesRes, loadEntities])

    return <Card>
        <EntityTableMemo entitiesRes={entitiesRes} setEntityRes={setEntityRes}
                         sortOrder={sortOrder}
                         sortField={sortField}
                         sortFieldType={sortFieldType}
                         rows={rows} firstRows={firstRows} updatePage={updatePage}/>
        <Button onClick={() => {
            setEntityRes({entities: [{properties: {}}], properties: [], total: 0})
        }}>New</Button>
        {entityRes && <EntityFormMemo entity={entityRes.entities[0]}
                                      loadEntities={loadEntities}
                                      updateEntity={e => setEntityRes({...entityRes, ...{entities: [e]}})}/>}
    </Card>
}


export default function App() {

    const confirm1 = function () {
        confirmDialog({
            message: 'Are you sure you want to proceed?',
            header: 'Confirmation',
            icon: 'pi pi-exclamation-triangle',
            defaultFocus: 'accept',
        })
    }

    console.log("APP render")
    return (
        <div>
            <AppProviders>
                <PropertyPage/>
                <Button onClick={confirm1} icon="pi pi-check" label="Confirm"></Button>
                <ConfirmDialog/>
                <EntityPage/>
            </AppProviders>

        </div>
    )
}