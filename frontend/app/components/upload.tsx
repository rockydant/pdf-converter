"use client";
import {useState } from "react";

export default function Upload() {
  const [file, setFile] = useState<File | null>(null);
  const [data, setData] = useState(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      if (e.target.files && e.target.files.length > 0) {
          setFile(e.target.files[0]);
      }
  };

  const handleUpload = async () => {
      if (file) {
          const formData = new FormData();
          formData.append('file', file);

          try {
              const response = await fetch('http://localhost:12345/upload', {
                  method: 'POST',
                  body: formData,
              });

              if (!response.ok) {
                throw new Error(`HTTP error: Status ${response.status}`);
              }
              console.log(response)
            //   let postsData = await response.json();
            //   setData(postsData);
            //   console.log(data)
          } catch (error) {
              console.error('Error uploading file:', error);
          }
      }
  };

  return (
    <div className="min-h-screen flex items-center justify-center">
    <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-lg">
        <h1 className="text-3xl font-semibold mb-4">Upload PDF File</h1>
        <div className="mb-4">
            <input type="file" onChange={handleFileChange} className="border-gray-300 border p-2 w-full" />
        </div>
        <button onClick={handleUpload} className="bg-blue-500 hover:bg-blue-600 text-white py-2 px-4 rounded-md">
            Upload
        </button>
    </div>
</div>
  );
}
