import './DropVideoSection.css';
import { React, useState } from "react";

export function DropVideoSection() {

    const [selectedFile, setSelectedFile] = useState(null);
    const [fileURL, setFileURL] = useState(null);

    const handleFileUpload = (event) => {
        const file = event.target.files[0];
        setSelectedFile(file);
        setFileURL(URL.createObjectURL(file));
    };

    const sendVideo = () => {
        const formData = new FormData();
        console.log(selectedFile)
        formData.append('video', selectedFile);
        fetch('http://localhost:8080/upload', {
            method: 'POST',
            mode: no-cors,
            body: formData
        }).then(function (res) {
            if (res.ok) {
              alert("Perfect! ");
            } else if (res.status == 401) {
              alert("Oops! ");
            }
          }, function (e) {
            alert("Error submitting form!");
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
                    <video src={fileURL} controls></video>
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