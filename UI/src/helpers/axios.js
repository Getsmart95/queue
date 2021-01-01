import axs from 'axios';
import {URL, USER_SESSION} from '../constants';
export {CancelToken} from 'axios';

export const axios = axs.create({
    baseURL: URL,
});

axios.defaults.headers.common['Accept-Language'] = 'ru';

axios.interceptors.request.use(function (config) {
    const user_session = localStorage.getItem(USER_SESSION);
    // const div = document.createElement('div');
    // const loader = document.getElementById('loader');
    // if (!loader) {
    //     div.id = 'loader';
    //     div.className = 'loader-container loader-absolute';
    //     document.body.appendChild(div);
    //     const divInner = document.createElement('div');
    //     divInner.className = 'loader';
    //     div.appendChild(divInner);
    // }

    if (user_session) {
        const token = JSON.parse(user_session);
        config.headers['Authorization'] = `Bearer ${token}`;
    } else {
        // config.headers['Authorization'] = '';
        // config.headers['Access-Control-Allow-Origin'] = '*';
    }
    return config;
}, function (error) {
    return Promise.reject(error);
});

axios.interceptors.response.use(function (response) {
    const div = document.getElementById('loader');
    if (div) {
        div.remove();
    }
    return response.data;
}, function (error) {
    return Promise.reject(error);
});
