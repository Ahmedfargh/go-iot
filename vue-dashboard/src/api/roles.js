import client from './client';

export const roleApi = {
  getRoles: async () => {
    const response = await client.get('/api/roles/');
    return response.data;
  },
  createRole: async (data) => {
    const response = await client.post('/api/roles/', data);
    return response.data;
  },
  syncPermissions: async (id, permissions) => {
    const response = await client.post(`/api/roles/${id}/permissions`, { permissions });
    return response.data;
  },
};
