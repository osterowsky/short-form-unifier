import './DropVideoSection.css';
import { React, useState } from "react";

export function DropVideoSection() {

    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileUpload = (event) => {
        const file = event.target.files[0];
        setSelectedFile(URL.createObjectURL(file));
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
        </div>
    )
}