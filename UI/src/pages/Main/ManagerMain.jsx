import React, { useState, useEffect } from "react";
import {axios} from "../../helpers"
import Table from "react-bootstrap/Table"
import moment from "moment"
const Tablo = (props) => {
    const [queues, setQueues] = useState([]);
    const [count, setCount] = useState([]);
    
    useEffect(() => {
        let TodayDate = moment().format("YYYY-MM-DD");
        axios
            .get(`queues/getByDate/${TodayDate}`)
            .then(res => {
                setQueues(res);
                setCount(res[0]) 
        });

        const interval = setInterval(()=>{
            axios
            .get(`queues/getByDate/${TodayDate}`)
            .then(res => {
                setQueues(res);
                setCount(res[0]) 
            });
        },10000)
            
        return()=>clearInterval(interval)

       
},[]);

    function padLeadZeros(num, size) {
    var s = num + ""
    while (s.length < size) s = "0" + s;
    return s
}

    const handleSubmit = e => {
        e.preventDefault();
    

        const _data = {
            status: "Approved"
        }
        axios
            .put(`queues/changeStatus/${count.id}`, _data)
                .then(res => {
                    window.location.reload(false)
                })
                .catch(err => {
                    console.error(err);
                })
        };
        
    
    const handleNext = e => {
        e.preventDefault();
        const _data = {
            status: "Completed"
        }
        
        axios
            .put(`queues/changeStatus/${count.id}`, _data)
                .then(res => {
                    window.location.reload(false)
                })
                .catch(err => {
                    console.error(err);
                })
        };


        return (
            <div className="auth-wrapper">
                 <nav className="navbar navbar-light bg-light justify-content-between">
                            <a className="navbar-brand">Queue</a>
                            
                 </nav>
            <div className="auth-inner">
               
            <Table striped bordered hover>
                <thead>
                    <tr>
                    <th>№</th>
                    <th>Номер</th>
                    <th>Статус</th>
                    <th>Окно №</th>
                    <th>Действие</th>
                    </tr>
                </thead>
                <tbody>
                {queues && queues.map(({queue_code, status, purpose_id}, index) => 
                    (queue_code === count.queue_code ? ( 
                    <tr>
                    <td>{index+1}</td>
                    <td><h3>{padLeadZeros(queue_code,3)}</h3></td>
                    <td><h3>{status}</h3></td>
                    <td><h3>{purpose_id.Int64}</h3></td>
                    <td>
                        <button onClick={handleSubmit} className="btn btn-primary btn-block">Принять</button>
                        <button onClick={handleNext} className="btn btn-primary btn-block">Следующий</button>
                    </td>
                    </tr>):("")
                )
                )}
                </tbody>
                <tbody>
                    <h2>В очереди</h2>
                {queues && queues.map(({queue_code, status},index) => 
                    (queue_code !== count.queue_code ? ( 
                    <tr>
                    <td>{index+1}</td>
                    <td><h3>{padLeadZeros(queue_code,3)}</h3></td>
                    <td><h3>{status}</h3></td>
                    </tr>):("")
                )
                )}
                </tbody>
            </Table>
               
            </div>
            </div>
        )
    };

export default Tablo;