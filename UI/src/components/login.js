import React, { Component, useEffect, useState } from "react";
import {axios, setStorageItem} from '../helpers'
import {USER_SESSION} from "../constants"
const Login = (props) => {
    const [auth, setAuth] = useState({});
    const [alert, setAlert] = useState({show: false, message: ''});
    
    const handleChange = e => {
        const {name, value} = e.target;

        setAuth({
            ...auth,
            [name]: value
        })
    };

    const handleSubmit = e => {
        e.preventDefault();
        const data = {
            ...auth       
        };

        
        const configs = {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        };
        // console.log(data)
            axios
                .post('login', data, configs)
                    .then(res => {
                        console.log(res);
                    setStorageItem(USER_SESSION, res);
                    setStorageItem("user", res);
                    setAlert({
                        show: true,
                        message: 'Объявление успешно создано!',
                        color: 'success'
                    });
                    setTimeout(() => {
                        props.history.push(`/operation`);
                    }, 3000)
                });
                
            
        };


        return (
            <form onSubmit={handleSubmit}>
                <h3>Авторизация</h3>
                <div className="form-group">
                    <label>Логин</label>
                    <input type="text" className="form-control" name="login" onChange={handleChange} placeholder="Введите логин" />
                </div>

                <div className="form-group">
                    <label>Пароль</label>
                    <input type="password" className="form-control" name="password" onChange={handleChange} placeholder="Введите пароль" />
                </div>

                <div className="form-group">
                {/* <label >{alert.message}</label> */}
                    <div className="custom-control custom-checkbox">
                        <input type="checkbox" className="custom-control-input" id="customCheck1" />
                        <label className="custom-control-label" htmlFor="customCheck1">Запомнить</label>
                    </div>
                </div>

                <button type="submit" className="btn btn-primary btn-block">Submit</button>
                <p className="forgot-password text-right">
                    Забыли <a href="#">пароль?</a>
                </p>
            </form>
        )
    };

export default Login;