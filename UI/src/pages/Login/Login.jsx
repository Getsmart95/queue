import React, { Component, useEffect, useState } from "react";
import {axios, setStorageItem} from '../../helpers'
import {USER_SESSION} from "../../constants"
const Login = (props) => {
    const [auth, setAuth] = useState({});
    
    const handleChange = e => {
        const {name, value} = e.target;
        
        setAuth({
            ...auth,
            [name]: value
        })
        console.log(auth)
    };

    const handleSubmit = e => {
        e.preventDefault();
        const data = {
            ...auth       
        };
        console.log(data)
        
        const configs = {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        };
        // console.log(data)
            axios
                .post('login', data, configs)
                    .then(res => {
                        console.log(res)
                    const {token} = res
                    setStorageItem(USER_SESSION, token);
                    setStorageItem("user", res);
                    setTimeout(() => {
                        props.history.push(`/main`);
                    }, 3000)
                });
                
            
        };


        return (
            <div className="auth-wrapper">
            <div className="auth-login">
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
            </div>
            </div>
        )
    };

export default Login;