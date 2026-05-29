import client from './client'

export const sectorApi = {
    getSectors: async (page = 1, pageSize = 10) => {
        const response = await client.get('/api/sectors/', {
            params: { page, pageSize }
        })
        return response.data
    },

    getSector: async (id) => {
        const response = await client.get(`/api/sectors/${id}`)
        return response.data
    },

    createSector: async (sectorData) => {
        const response = await client.post('/api/sectors/', sectorData)
        return response.data
    },

    updateSector: async (id, sectorData) => {
        const response = await client.put(`/api/sectors/${id}`, sectorData)
        return response.data
    },

    deleteSector: async (id) => {
        const response = await client.delete(`/api/sectors/${id}`)
        return response.data
    }
}
