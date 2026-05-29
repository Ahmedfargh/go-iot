import client from './client';

export const deviceApi = {
  getDevices: async (page = 1, pageSize = 10) => {
    const response = await client.get('/api/devices/', { params: { page, pageSize } });
    return response.data;
  },
  getDevice: async (id) => {
    const response = await client.get(`/api/devices/${id}`);
    return response.data;
  },
  createDevice: async (deviceData) => {
    const response = await client.post('/api/devices/', deviceData);
    return response.data;
  },
  updateDevice: async (id, deviceData) => {
    const response = await client.put(`/api/devices/${id}`, deviceData);
    return response.data;
  },
  deleteDevice: async (id) => {
    const response = await client.delete(`/api/devices/${id}`);
    return response.data;
  },
};
