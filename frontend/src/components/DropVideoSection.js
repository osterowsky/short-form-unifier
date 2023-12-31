import './DropVideoSection.css';
import { React, useState } from "react";

export function DropVideoSection() {

    const [selectedFile, setSelectedFile] = useState(null);
    const [fileURL, setFileURL] = useState(null);
    const [title, setTitle] = useState("");
    const [PrivacyLevel, setPrivacyLevel] = useState("public");
    const [youtubeDescription, setYoutubeDescription] = useState("");
    const [youtubeTags, setYoutubeTags] = useState("");

    const handleFileUpload = (event) => {
        const file = event.target.files[0];
        setSelectedFile(file);
        setFileURL(URL.createObjectURL(file));
    };

    const sendVideo = () => {
        // Check if a file is selected
        if (!selectedFile) {
            alert("Please select a file first.");
            return;
        }

        const videoData = {
            title: title,
            privacy_level: PrivacyLevel,
            tiktok: {
                disable_duet: false,
            },
            youtube: {
                description: youtubeDescription,
                tags: youtubeTags.split(',').map(tag => tag.trim()), // Split tags and trim whitespace
            },
        };

        const formData = new FormData();
        formData.append('video', selectedFile);
        formData.append('data', JSON.stringify(videoData)); // 'data' is the key for the JSON payload

        fetch('http://localhost:8080/upload', {
            method: 'POST',
            mode: 'no-cors', // no-cors, *cors, same-origin
            body: formData
        }).then(function (res) {
            if (res.ok) {
              alert("Perfect!");
            } else if (res.status === 401) {
              alert("Oops!");
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
                    <div className="inputContainer">
                        <label className="inputLabel">
                            Title
                            <i className="fab fa-youtube"></i>
                            <i className="fab fa-tiktok"></i>
                            <i className="fab fa-instagram"></i>
                        </label>
                        <input
                            type="text"
                            placeholder="Best video in the world"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                            className="inputTextField"
                        />
                    </div>
                    <div className="inputContainer">
                        <label className="inputLabel">
                            Description
                            <i className="fab fa-youtube"></i>
                            </label>
                        <textarea
                            rows={4}
                            placeholder="This video is sick"
                            value={youtubeDescription}
                            onChange={(e) => setYoutubeDescription(e.target.value)}
                        />
                    </div>
                    <div className="inputContainer">
                        <label className="inputLabel">
                            Privacy Level
                            <i className="fab fa-youtube"></i>
                            <i className="fab fa-tiktok"></i>
                        </label>
                            <div>
                                <input id="public" type="radio" name="privacyLevel" value="public" onChange={(e) => setPrivacyLevel(e.target.value)} />
                                <label class="inputLabelRadio" htmlFor="public">Public</label>
                            </div>
                            <div>
                                <input id="private" type="radio" name="privacyLevel" value="private" onChange={(e) => setPrivacyLevel(e.target.value)} />
                                <label class="inputLabelRadio" for="private">Private</label>
                            </div>
                    </div>
                    <div className="inputContainer">
                        <label className="inputLabel">
                            Tags
                            <i className="fab fa-youtube"></i>
                        </label>
                        <input
                            type="text"
                            placeholder="your-mother,thatswhatshesaid,420"
                            value={youtubeTags}
                            onChange={(e) => setYoutubeTags(e.target.value)}
                        />
                    </div>
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