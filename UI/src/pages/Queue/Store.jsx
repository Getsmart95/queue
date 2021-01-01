import React, { useState, useEffect } from "react";
import {axios} from "../../helpers"
import moment from "moment"

const Store = (props) => {
    const [services, setServices] = useState([]);
    const [purpose, setPurpose] = useState({});
    useEffect(() => {
        axios
            .get('purposes')
            .then(res => {
                setServices(res);
            })
    }, []);

    const handleChange = e => {
        const service = services.filter(({id}) => id === Number(e.target.value));
        const {id} = service[0]
        setPurpose(id)
       
    };

    const handleSubmit = e => {
        e.preventDefault();
        let TodayDate = moment().format("YYYY-MM-DD");  
        const _data = {
            purpose_id: purpose ,
            date: TodayDate     
        };

        axios
            .post(`/queue/addQueue`, _data)
            .then(res => {
                    alert('Очередь успешно добавлена')
                setTimeout(() => {
                    props.history.push(`/store`);
                }, 3000)
            })
            .catch(err => {
                console.error(err);
            })
        };

    return (
        <div className="auth-wrapper">
            <div className="purpose-list">
                <form onSubmit={handleSubmit}>
                    <h3>Выберите услугу</h3>
                    {services && services.map(({id, name}) =>
                                    <button type="submit" name="service" value={id} onClick={handleChange} className="btn btn-primary btn-block">{name}</button>
                                )}
                </form>
            </div>
        </div>
    )
};

export default Store;