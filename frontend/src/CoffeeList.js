import React, { useEffect, useState} from "react";
import axios from 'axios';
import Table from 'react-bootstrap/Table'

const CoffeeList = () => {
    // const [loading, setLoading] = useState(true);
    const [data, setData] = useState([]);

    useEffect(() => {
        axios.get(window.global.api_location+'/products').then(function(response){
            console.log(response.data);
            setData(response.data)

        }).catch(function(error){
            console.log(error)
        });

    }, []);

    function getProducts() {
        let table = []

        for (let i=0; i < data.length; i++) {
            table.push(
                <tr key={i}>
                    <td>{data[i].name}</td>
                    <td>{data[i].price}</td>
                    <td>{data[i].sku}</td>
                </tr>
            );
        }
        return table
    }

    return (
        <div>
            <h1 style={{marginBottom: "40px"}}>Menu</h1>
            <Table>
                <thead>
                    <tr>
                        <th>
                            Name
                        </th>
                        <th>
                            Price
                        </th>
                        <th>
                            SKU
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {getProducts()}
                </tbody>
            </Table>
        </div>
    )
}

export default CoffeeList;