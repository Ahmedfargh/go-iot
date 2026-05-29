import client from './client';

export const authApi = {
  login: async (email, password) => {
    const response = await client.post('/login', { email, password });
    return response.data;
  },
};
