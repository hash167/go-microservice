import React, { useEffect, useState} from "react";
import Toast from 'react-bootstrap/Toast';

const Toaster = (props) => {
    const {show, msg, updateToast} = props 
    const [localShow, setLocalShow] = useState(show);

    useEffect(() => {
        setLocalShow(show)
    }, [localShow, show]);

    return (
        <Toast onClose={() => updateToast(false)} show={localShow} delay={3000} autohide>
            <Toast.Header>
                <strong className="mr-auto">File Upload</strong>
            </Toast.Header>
            <Toast.Body>{msg}</Toast.Body>
        </Toast>
    )
}

export default Toaster;