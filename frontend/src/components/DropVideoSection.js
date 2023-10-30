import './DropVideoSection.css';
import React, { useState } from 'react';
import { Upload, Button } from "antd";

export function DropVideoSection() {
    const [videoSrc , seVideoSrc] = useState("");

    const handleChange = ({file}) => {
        var url = URL.createObjectURL(file.originFileObj);
        seVideoSrc(url);
    };

    return (
        <div className="VideoSection">
            <div className="UploadButton">
                <Upload className="mt-3 mb-3"
                    accept=".mp4"
                    action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                    listType="picture"
                    maxCount={1}
                    onChange={handleChange}
                    >
                    <Button>
                    Upload Video
                    </Button>
                </Upload>
            </div>

            <div className="GridWrapper">
                <div className="VideoDisplay">
                    <video src={videoSrc} controls></video>
                </div>
                <div className='VideoUpload'>
                </div>
            </div>
        </div>
    )
}