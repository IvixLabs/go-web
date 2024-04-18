import {createRoot} from 'react-dom/client';
import App, {RedirectToUsers} from "./dashboard/App";
import * as React from 'react';
import {PrimeReactProvider} from "primereact/api";

import 'primeflex/primeflex.css'; // flex
import 'primereact/resources/primereact.css'; //core css
import 'primereact/resources/themes/bootstrap4-light-blue/theme.css'; //theme
import 'primeicons/primeicons.css'; //icons
import './dashboard.css'
import {createBrowserRouter, redirect, RouterProvider} from "react-router-dom";
import UserListPage from "./dashboard/UserListPage";
import ProductListPage from "./dashboard/ProductListPage";


const router = createBrowserRouter([
    {
        path: "/dashboard",
        element: <App/>,
        children: [
            {
                path: "",
                element: <RedirectToUsers/>
            },
            {
                path: 'users',
                element: <UserListPage/>
            },
            {
                path: 'products',
                element: <ProductListPage/>
            }
        ]


    },
]);

const root = createRoot(document.getElementById('root') as HTMLElement);

root.render(
    <React.StrictMode>
        <PrimeReactProvider>
            <RouterProvider router={router}/>
        </PrimeReactProvider>
    </React.StrictMode>
);