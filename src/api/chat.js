// api.js

import axios from 'axios';

const baseURL = 'http://localhost:8080'; // Update with the correct protocol and host
const api = axios.create({
  baseURL: baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export default api;
