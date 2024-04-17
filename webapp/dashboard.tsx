import { createRoot } from 'react-dom/client';
import App from "./dashboard/App";
import * as React from 'react';
import {PrimeReactProvider} from "primereact/api";

import 'primeflex/primeflex.css'; // flex
import 'primereact/resources/primereact.css'; //core css
import 'primereact/resources/themes/bootstrap4-light-blue/theme.css'; //theme
import 'primeicons/primeicons.css'; //icons
import './dashboard.css'

const root = createRoot(document.getElementById('root') as HTMLElement);
root.render(
    <React.StrictMode>
        <PrimeReactProvider>
            <App />
        </PrimeReactProvider>
    </React.StrictMode>
);