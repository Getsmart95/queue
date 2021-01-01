import React, { Component, useEffect, useState } from "react";
import {axios, getStorageItem} from '../helpers'
const Service = (props) => {
    const [accounts, setAccount] = useState([]);
    const [cardNumber, setCardNumber] = useState('Счет не выбран');
    const [balance, setBalance] = useState('Счет не выбран');
    const [transfer, setTransfer] = useState({});
    const [services, setService] = useState([]);
    useEffect(() => {
        const {ID} = getStorageItem();
            axios
                .get(`getAllAccounts/${ID}`)
                .then(res => {
                    setAccount(res);
                })

            axios
            .get(`getAllServices`)
            .then(res => {
                console.log(res);
                setService(res);
            })
    }, []);

    const handleChange = e => {
        const {name, value} = e.target;

        setTransfer({
            ...transfer,
            [name]: value
        })
        console.log(transfer);
    };

    const handleAccount = e => {
        const account = accounts.filter(({ID}) => ID === Number(e.target.value));
            const {CardNumber, Balance, AccountNumber} = account[0];
            setCardNumber(CardNumber);
            setBalance(Balance);
            setTransfer({
                ...transfer,
                AccountNumber
            })
    }

    const handleService = e => {
        const service = services.filter(({ID}) => ID === Number(e.target.value));
            const {ID, Name, ServiceAccountNumber} = service[0];
            setTransfer({
                ...transfer,
                ID,
                Name,
                ServiceAccountNumber
            })
    }
    const handleSubmit = e => {
        e.preventDefault();
        
        const _data = {
            ...transfer       
        };
        console.log(_data);
        
        const configs = {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        };
            axios
                .post(`payService`, _data, configs)
                .then(res => {
                    if(res === false)
                        alert('Неверный номер карты')
                    else
                        alert('Перевод успешно выполнен')
                    setTimeout(() => {
                        // props.history.push(`/operation`);
                    }, 3000)
                })
                .catch(err => {
                    console.error(err);
                })
        };


        return (
            <form onSubmit={handleSubmit}>
                
                <div className="form-group">
                    <label>Выберите счет</label>
                    <select className="form-control" onChange={handleAccount}>
                        <option value="Выбрать">Выбрать</option>
                            {accounts && accounts.map(({ID, AccountNumber}) =>
                                        <option value={ID} key={ID}>{AccountNumber}</option>
                                    )}
                    </select>
                </div>

                <div className="form-group">
                    <label>Номер карты</label>
                    <input type="text" className="form-control" name="cardNumber" value={cardNumber} placeholder="Счет не выбран" />
                </div>

                <div className="form-group">
                    <label>Баланс</label>
                    <input type="text" className="form-control" name="balance" value={balance} placeholder="Счет не выбран" />
                </div>

                <hr class="mb-4"/>

                <div className="form-group">
                    <label>Выберите услугу</label>
                    <select className="form-control" onChange={handleService}>
                        <option value="Выбрать">Выбрать</option>
                            {services && services.map(({ID, Name}) =>
                                        <option value={ID} key={ID}>{Name}</option>
                                    )}
                    </select>                
                </div>

                <div className="form-group">
                    <label>Сумма оплаты</label>
                    <input type="text" className="form-control" name="ammount" onChange={handleChange} placeholder="Введите сумму оплаты" />
                </div>

                <button type="submit" className="btn btn-primary btn-block">Оплатить</button>
               
            </form>
        )
    };

export default Service;