import React, { Component, useEffect, useState } from "react";
import {axios, getStorageItem} from '../helpers'
const Login = (props) => {
    const [accounts, setAccount] = useState([]);
    const [cardNumber, setCardNumber] = useState('Счет не выбран');
    const [balance, setBalance] = useState('Счет не выбран');
    const [status, setStatus] = useState('Счет не выбран');
    useEffect(() => {
        const {ID} = getStorageItem();
            axios
                .get(`getAllAccounts/${ID}`)
                .then(res => {
                    setAccount(res);
                })
    }, []);


    const handleAccount = e => {
        const account = accounts.filter(({ID}) => ID === Number(e.target.value));
            const {CardNumber, Balance, Status} = account[0];
            setCardNumber(CardNumber);
            setBalance(Balance);
            setStatus(Status);
    }
    // Временно закоментил
    // const handleSubmit = e => {
    //     e.preventDefault();
        
    //     const _data = {
    //         ...transfer       
    //     };
    //     console.log(_data);
        
    //     const configs = {
    //         headers: {
    //             'Content-Type': 'multipart/form-data'
    //         }
    //     };
    //         axios
    //             .post(`transferTo/${_data.transferCardNumber}`, _data, configs)
    //             .then(res => {
    //                 if(res === false)
    //                     alert('Неверный номер карты')
    //                 else
    //                     alert('Перевод успешно выполнен')
    //                 setTimeout(() => {
    //                     props.history.push(`/operation`);
    //                 }, 3000)
    //             })
    //             .catch(err => {
    //                 console.error(err);
    //             })
    // };


        return (
            <form>
                <h4>Список счетов</h4>
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

                <div className="form-group">
                    <label>Статус</label>
                    <input type="text" className="form-control" name="Status" value={status} placeholder="Счет не выбран" />
                </div>
               
            </form>
        )
    };

export default Login;