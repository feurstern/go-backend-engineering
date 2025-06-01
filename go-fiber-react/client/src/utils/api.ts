import axios from 'axios';

const api = axios.create({
  baseURL: 'http://127.0.0.1:5353/',
  headers: {
    'Content-Type': 'apllication/json',
  },
});

export default api;
