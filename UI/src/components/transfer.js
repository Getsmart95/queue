import React, { Component, useEffect, useState } from "react";
import {axios, getStorageItem} from '../helpers'
const Transfer = (props) => {
    const [accounts, setAccount] = useState([]);
    const [cardNumber, setCardNumber] = useState('Счет не выбран');
    const [balance, setBalance] = useState('Счет не выбран');
    const [transfer, setTransfer] = useState({});
    useEffect(() => {
        const {ID} = getStorageItem();
            axios
                .get(`getAllAccounts/${ID}`)
                .then(res => {
                    setAccount(res);
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
            
            console.log(accounts.filter(({ID}) => ID === Number(e.target.value)));
            console.log(Number(e.target.value));
            console.log(account);
            const {CardNumber, Balance, AccountNumber} = account[0];
            setCardNumber(CardNumber);
            setBalance(Balance);
            setTransfer({
                ...transfer,
                AccountNumber
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
                .post(`transferTo/${_data.transferCardNumber}`, _data, configs)
                .then(res => {
                    if(res === false)
                        alert('Неверный номер карты')
                    else
                        alert('Перевод успешно выполнен')
                    setTimeout(() => {
                        props.history.push(`/operation`);
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
                    <label>Перевести на карту</label>
                    <input type="text" className="form-control" name="transferCardNumber" onChange={handleChange} placeholder="Введите номер карты" />
                </div>

                <div className="form-group">
                    <label>Сумма перевода</label>
                    <input type="text" className="form-control" name="transferAmmount" onChange={handleChange} placeholder="Введите сумму перевода" />
                </div>

                <div className="form-group">
                    <label>Введите сообщение получателю</label>
                    <input type="text" className="form-control" name="message" onChange={handleChange} placeholder="Введите сообщение получателю" />
                </div>
                <button type="submit" className="btn btn-primary btn-block">Перевести</button>
               
            </form>
        )
    };

export default Transfer;