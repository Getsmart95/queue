import React, { useState, useEffect } from "react";
import {axios, getStorageItem} from "../../helpers"
import {USER_SESSION} from "../../constants"
// import Table from 'react-bootstrap-table'
import Table from "react-bootstrap/Table"
const Tablo = (props) => {
    const [data, setData] = useState({});  
    const [queues, setQueues] = useState([]);
    const [count, setCount] = useState({});
    useEffect(() => {
        axios
            .get('queues/getByDate/2020-12-31')
            .then(res => {
                setQueues(res);
                setCount(res[0])
            });
        // const interval = setInterval(()=>{
        //     axios
        //     .get('queues/getByDate/2020-12-31')
        //     .then(res => {
        //         setQueues(res);
        //     });
        // },10000)
            
        // return()=>clearInterval(interval)
    // setTimeout(() => {
    //     props.history.push(`/tablo`);
    // }, 30000)
       
},[]);

    function padLeadZeros(num, size) {
    var s = num + ""
    while (s.length < size) s = "0" + s;
    return s
}
    const handleChange = e => {
        const {name, value} = e.target;

        setData({
            [name]: value
        })
    };

    const handleSubmit = e => {
        e.preventDefault();
        const path = data.path;
        props.history.push(`/${path}`);            
        };


        return (
            <div className="auth-wrapper">
                 <nav className="navbar navbar-light bg-light justify-content-between">
                            <a className="navbar-brand">Navbar</a>
                            
                 </nav>
            <div className="auth-inner">
               
            <form onSubmit={handleSubmit}>
            <Table striped bordered hover>
                <thead>
                    <tr>
                    <th>№</th>
                    <th>Номер</th>
                    <th>Цель визита</th>
                    <th>Статус</th>
                    <th>Окно №</th>
                    <th>Действие</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                    <td>1</td>
                    <td><h3>{padLeadZeros(count.queue_code,3)}</h3></td>
                    <td><h3>{count.status}</h3></td>
                    <td>
                        <button className="btn btn-primary btn-block">Принять</button>
                        <button className="btn btn-primary btn-block">Следующий</button>
                    </td>
                    </tr>
                </tbody>
                <tbody>
                    <h2>В очереди</h2>
                {queues && queues.map(({queue_code, status}) => 
                    (queue_code !== count.queue_code ? ( 
                    <tr>
                    <td>1</td>
                    <td><h3>{padLeadZeros(queue_code,3)}</h3></td>
                    <td><h3>{status}</h3></td>
                    </tr>):("")
                )
                )}
                </tbody>
            </Table>
                {/* <h3>Список операции</h3>
                <div width="100%" heigth="300px">1</div>
                <button type="submit" name="path" value="account" onClick={handleChange} className="btn btn-primary">Посмотреть список счетов</button>
                <button type="submit" name="path" value="transfer" onClick={handleChange} className="btn btn-primary btn-block">Перевести деньги другому клиенту</button>
                <button type="submit" name="path" value="service" onClick={handleChange} className="btn btn-primary btn-block">Оплатить услугу</button>
                <button type="submit" name="path" value="" onClick={handleChange} className="btn btn-primary btn-block">Выйти из аккаунта</button> */}

            </form>
            </div>
            </div>
        )
    };

export default Tablo;