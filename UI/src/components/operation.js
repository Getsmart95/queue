import React, { useState } from "react";
import {axios, getStorageItem} from "../helpers"
import {USER_SESSION} from "../constants"
const Operation = (props) => {
    const [data, setData] = useState({});    

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
            <form onSubmit={handleSubmit}>
                <h3>Список операции</h3>
             
                <button type="submit" name="path" value="account" onClick={handleChange} className="btn btn-primary btn-block">Посмотреть список счетов</button>
                <button type="submit" name="path" value="transfer" onClick={handleChange} className="btn btn-primary btn-block">Перевести деньги другому клиенту</button>
                <button type="submit" name="path" value="service" onClick={handleChange} className="btn btn-primary btn-block">Оплатить услугу</button>
                <button type="submit" name="path" value="" onClick={handleChange} className="btn btn-primary btn-block">Выйти из аккаунта</button>

            </form>
        )
    };

export default Operation;