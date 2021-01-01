import React, { useState, useEffect } from "react";
import {axios} from "../../helpers"
import Table from "react-bootstrap/Table"
import moment from "moment";
const Tablo = (props) => {
    const [queues, setQueues] = useState([]);

    useEffect(() => {
        let TodayDate = moment().format("YYYY-MM-DD");
        axios
            .get(`queues/getByDate/${TodayDate}`)
            .then(res => {
                setQueues(res);
            });
            const interval = setInterval(()=>{
                axios
                .get(`queues/getByDate/${TodayDate}`)
                .then(res => {
                    setQueues(res);
                });
            },10000)
        //moment().format("DD-MM-YYYY hh:mm:ss"));
            return()=>clearInterval(interval)
    },[]);
 
    function padLeadZeros(num, size) {
        var s = num + ""
        while (s.length < size) s = "0" + s;
        return s
    }

    return (
        <div className="auth-wrapper">
            <nav className="navbar navbar-light bg-light justify-content-between">
                    <a className="navbar-brand">Navbar</a>
            </nav>
            <div className="auth-inner">
            
            <Table striped bordered hover>
                <thead>
                    <tr>
                    <th>#</th>
                    <th>Номер</th>
                    <th>Статус</th>
                    </tr>
                </thead>
                {queues && queues.map(({queue_code, status},index) =>
                    <tbody>
                    <tr>
                    <td>{index+1}</td>
                    <td><h3>{padLeadZeros(queue_code,3)}</h3></td>
                    <td><h3>{status}</h3></td>
                    </tr>
                </tbody>
                )}
            </Table>
            </div>
        </div>
    )
    };

export default Tablo;