import axios from "axios";

export const API = axios.create({
    baseURL: '/api',
    timeout: 1000,
    headers: {'Authorization': `Bearer ${localStorage.getItem('token')}`}
});
