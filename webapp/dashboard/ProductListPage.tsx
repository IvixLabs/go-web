import * as React from "react";
import {useEffect, useState} from "react";
import {Card} from "primereact/card";
import {DataTable} from "primereact/datatable";
import {Column} from "primereact/column";

async function apiGetProducts(): Promise<ProductListItem[]> {
    const res = await fetch("/api/product/list")
    if (!res.ok) {
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

interface ProductListItem {
    id: string
    title: string
    brand: string
}

function ProductTable() {

    const [products, setProducts] = useState<ProductListItem[]>([])

    const loadProducts = async () => {
        setProducts(await apiGetProducts())
    }

    useEffect(() => {
        loadProducts().then()
    }, [])


    return <DataTable value={products} size="small">
        <Column field="id" header="Id"/>
        <Column field="title" header="Title"/>
        <Column field="brand" header="Brand"/>
    </DataTable>
}

export default function ProductListPage() {
    return <Card>
        <div className="flex flex-column gap-2">
            <ProductTable/>
        </div>
    </Card>
}