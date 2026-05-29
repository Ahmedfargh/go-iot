import client from './client';

export const adminApi = {
  getAdmins: async (page = 1, pageSize = 10) => {
    const response = await client.get('/api/admins/', { params: { page, pageSize } });
    return response.data;
  },
  getAdmin: async (id) => {
    const response = await client.get(`/api/admins/${id}`);
    return response.data;
  },
  createAdmin: async (formData) => {
    const response = await client.post('/api/admins/', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    });
    return response.data;
  },
  updateAdmin: async (id, formData) => {
    // Note: Some backends require POST with _method=PUT for multipart updates, 
    // but the Postman says PUT. 
    const response = await client.put(`/api/admins/${id}`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    });
    return response.data;
  },
  deleteAdmin: async (id) => {
    const response = await client.delete(`/api/admins/${id}`);
    return response.data;
  },
};
