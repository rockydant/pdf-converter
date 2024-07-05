"use client";

import { SetStateAction, useState } from "react";
import Upload from "./components/upload";
import { TabView, TabPanel } from 'primereact/tabview';
import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';
import ContentPDF from "./components/data-table";


export default function Home() {
    return (
        <TabView>
            <TabPanel header="Upload">
                <Upload />
            </TabPanel>
            <TabPanel header="Data">
               <ContentPDF />
            </TabPanel>
            
        </TabView>
    );
}
