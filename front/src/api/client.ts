import axios from 'axios';

export const ApiClient = axios.create({
    baseURL: 'http://localhost:8080/api/v1', // Replace with your API base URL
    headers: {
        'Content-Type': 'application/json',
    },
});
