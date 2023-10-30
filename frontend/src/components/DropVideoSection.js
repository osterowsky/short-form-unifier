import './DropVideoSection.css';
import { React, useState } from "react";

export function DropVideoSection() {

    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileUpload = (event) => {
        const file = event.target.files[0];
        setSelectedFile(URL.createObjectURL(file));
    };

    const sendVideo = () => {
        const formData = new FormData();
        formData.append('file', selectedFile);
        fetch('http://localhost:8080/upload', {
            method: 'POST',
            body: formData
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
        })
        .catch(error => {
            console.error(error);
        });
    };

    return (
        <div className="VideoSection">

            <div className="UploadButton">
                <button>
                    <label htmlFor="fileUpload">Upload Video</label>
                    <input
                    type="file"
                    id="fileUpload"
                    accept="video/*"  // Accept only video files
                    onChange={handleFileUpload}
                    style={{ display: 'none' }}
                    />
                </button>
            </div>

            <div className="GridWrapper">
                <div className="VideoDisplay">
                    <video src={selectedFile} controls></video>
                </div>
                <div className='VideoUpload'>
                </div>
            </div>

            <div className="UploadButton">
                <button onClick={sendVideo}>
                    <label>Upload</label>
                </button>
            </div>
        </div>
    )
}