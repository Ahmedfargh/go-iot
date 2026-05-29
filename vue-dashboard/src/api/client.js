import axios from 'axios';

const client = axios.create({
  baseURL: '', // Use relative URL so Vite proxy can handle it
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
});

// Request interceptor to add the auth token
client.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

// Response interceptor to handle errors and logging
client.interceptors.response.use((response) => {
  console.log(`[API Response] ${response.config.method.toUpperCase()} ${response.config.url}:`, response.data);
  return response;
}, (error) => {
  if (error.response) {
    console.error(`[API Error] ${error.config.method.toUpperCase()} ${error.config.url}:`, error.response.status, error.response.data);
  } else {
    console.error(`[API Network Error]:`, error.message);
  }
  
  if (error.response && error.response.status === 401) {
    const isLoginRequest = error.config.url.endsWith('/login');
    if (!isLoginRequest) {
      // Handle unauthorized error (e.g., redirect to login)
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
  }
  return Promise.reject(error);
});

export default client;
