import axios from "axios";

export const API = axios.create({
    baseURL: '/api',
    timeout: 5000,
    headers: {'Authorization': `Bearer ${localStorage.getItem('token')}`}
});
