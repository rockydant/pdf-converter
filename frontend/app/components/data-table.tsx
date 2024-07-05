import { DataTable } from 'primereact/datatable';
import { Column } from 'primereact/column';
import { useEffect, useState } from 'react';
import axios from 'axios';

export default function ContentPDF() {
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:12345/files');
                setData(response.data);
            } catch (err) {
                console.error('Error loading:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    if (loading) return <p>Loading...</p>;
    return (
        <div className="min-h-screen flex items-center justify-center">
            <DataTable value={data} tableStyle={{ minWidth: '50rem' }}>
                <Column field="id" header="ID"></Column>
                <Column field="title" header="Title"></Column>
                <Column field="content" header="Content"></Column>
            </DataTable>
        </div>
    );
}
