import React, {useState} from "react";
import { useForm } from "react-hook-form";
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Toaster from './Toast.js'

import axios from 'axios';


const Admin = () => {
    // const [validated, setValidated] = useState(false)
    const validated = false
    const [id, setId] = useState('')
    const [file, setFile] = useState(null)
    const [buttonDisabled, setButtonDisabled] = useState(false)
    const [toastShow, setToastShow] = useState(false)
    const [toastText, setToastText] = useState('asd')
    const { handleSubmit } = useForm(); 

    const onSubmit = () => {
        console.log("Form was submitted, now the modal can be closed");                                                                     
        // create the data
        const data = new FormData()
        data.append('file', file);
        data.append('id', id);

        // upload the file
        axios.post(
            window.global.files_location, 
            data, 
            {'content-type': `multipart/form-data; boundary=${data._boundary}`})
        .then(res => {
            console.log(res);
            if(res.status === 200) {
                setToastText("Uploaded file");
            } else {
                setToastText("Unable to upload file. Error:" +res.statusText);
            }
            setButtonDisabled(false);
            setToastShow(true);
        }).catch(error => {
            console.log("Err" + error);
            setToastText("Unable to upload file. " + error);
            setButtonDisabled(false);
            setToastShow(true);
        });                                                                                                       
    }

    const updateToast = (show) => {
        setToastShow(show)
    }

    const handleIdChange = (event) => {
        setId(event.target.value)
        console.log(event.target)

    }
    const handleFileChange = (event) => {
        setFile(event.target.files[0])
        console.log(event.target)

    }
    return (
        <div>
          <h1 style={{marginBottom: "40px"}}>Admin</h1>
          <Container className="text-left">
            <Form noValidate validated={validated} onSubmit={handleSubmit(onSubmit)}>
                  <Form.Group as={Row} controlId="productID">
                      <Form.Label column sm="2">Product ID:</Form.Label>
                      <Col sm="6">
                          <Form.Control type="text" name="id" placeholder="" required style={{width: "80px"}} value={id} onChange={handleIdChange}/>
                          <Form.Text className="text-muted">Enter the product id to upload an image for</Form.Text>
                          <Form.Control.Feedback type="invalid">Please provide a product ID.</Form.Control.Feedback>
                      </Col>
                  <Col sm="4">
                      <Toaster show={toastShow} msg={toastText} updateToast={updateToast}/>
                  </Col>
                  </Form.Group>
                  <Form.Group as={Row}>
                      <Form.Label column sm="2">File:</Form.Label>
                      <Col sm="10">
                          <Form.Control type="file" name="file" placeholder="" required onChange={handleFileChange}/>
                          <Form.Text className="text-muted">Image to associate with the product</Form.Text>
                          <Form.Control.Feedback type="invalid">Please select a file to upload.</Form.Control.Feedback>
                      </Col>
                  </Form.Group>
                  <Button  type="submit" disabled={buttonDisabled}>Submit form</Button>
            </Form>
          </Container>
          </div>
    )
}

export default Admin;