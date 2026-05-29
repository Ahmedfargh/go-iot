import { defineStore } from 'pinia';
import { authApi } from '../api/auth';

export const useAuthStore = defineStore('auth', {
  state: () => {
    let user = null;
    try {
      const storedUser = localStorage.getItem('user');
      user = storedUser ? JSON.parse(storedUser) : null;
    } catch (e) {
      console.warn('Failed to parse user from localStorage', e);
      localStorage.removeItem('user');
    }
    return {
      token: localStorage.getItem('token') || null,
      user
    };
  },
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    async login(email, password) {
      try {
        const data = await authApi.login(email, password);
        this.token = data.token;
        this.user = data.user || { email }; // Fallback user info
        
        localStorage.setItem('token', data.token);
        localStorage.setItem('user', JSON.stringify(this.user));
        
        return data;
      } catch (error) {
        throw error;
      }
    },
    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    },
  },
});
