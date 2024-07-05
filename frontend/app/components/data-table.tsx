import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';

export default function ContentPDF() {
    return (
        <div className="min-h-screen flex items-center justify-center">
            <DataTable tableStyle={{ minWidth: '50rem' }}>
                <Column field="code" header="Code"></Column>
                <Column field="name" header="Name"></Column>
                <Column field="category" header="Category"></Column>
                <Column field="quantity" header="Quantity"></Column>
            </DataTable>
        </div>
    );
}
