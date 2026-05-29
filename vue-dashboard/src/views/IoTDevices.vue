<template>
  <div class="px-margin-mobile md:px-0">
    <header class="flex flex-col md:flex-row justify-between md:items-end mb-lg gap-4">
      <div>
        <h1 class="font-headline-md text-headline-md text-on-surface mb-xs">Active Devices</h1>
        <p class="text-on-surface-variant font-body-md">Managing <span class="text-primary font-bold">{{ devices.length }} nodes</span> across the industrial network</p>
      </div>
      <div class="flex gap-md w-full md:w-auto">
        <router-link to="/devices/provision" class="bg-primary hover:bg-primary-fixed-dim text-on-primary font-label-caps text-label-caps px-lg py-sm rounded-lg flex items-center gap-sm transition-all active:scale-95 ml-auto md:ml-0">
          <span class="material-symbols-outlined text-[18px]">add</span>
          PROVISION NODE
        </router-link>
      </div>
    </header>

    <!-- Device Registry Table -->
    <div class="glass-panel rounded-xl overflow-hidden">
      <div class="px-md py-md bg-white/5 border-b border-white/10 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
        <div class="flex flex-col sm:flex-row items-start sm:items-center gap-md w-full sm:w-auto">
          <span class="font-title-sm text-title-sm">Technical Registry</span>
          <div class="flex items-center bg-surface-container-lowest border border-white/10 px-sm py-1 rounded w-full sm:w-auto">
            <span class="material-symbols-outlined text-[18px] text-on-surface-variant mr-xs">search</span>
            <input v-model="searchQuery" class="bg-transparent border-none focus:ring-0 font-body-md text-body-md text-on-surface py-0 w-full sm:w-48 outline-none" placeholder="Search identifiers..." type="text"/>
          </div>
        </div>
      </div>
      <div class="overflow-x-auto custom-scrollbar">
        <table class="w-full border-collapse min-w-[800px]">
          <thead>
            <tr class="bg-surface-container-high/50 text-left border-b border-white/10">
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Node Identifier</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Status</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Temperature</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-center">Humidity</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase">Sector</th>
              <th class="px-md py-3 font-label-caps text-label-caps text-on-surface-variant uppercase text-right">Actions</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-white/5">
            <tr v-if="isLoading">
              <td colspan="6" class="px-md py-12 text-center">
                <span class="material-symbols-outlined animate-spin text-primary text-3xl">refresh</span>
                <p class="font-label-caps text-on-surface-variant mt-sm">Synchronizing Node Matrix...</p>
              </td>
            </tr>
            <tr v-else-if="error">
              <td colspan="6" class="px-md py-12 text-center">
                <span class="material-symbols-outlined text-error text-3xl">report_problem</span>
                <p class="font-label-caps text-error mt-sm">{{ error }}</p>
                <button @click="fetchDevices" class="mt-md text-primary font-title-sm hover:underline">Retry Handshake</button>
              </td>
            </tr>
            <tr v-else-if="filteredDevices.length === 0">
              <td colspan="6" class="px-md py-12 text-center text-on-surface-variant font-body-md">
                {{ devices.length === 0 ? 'No hardware nodes registered in the array.' : 'No nodes match the search criteria.' }}
              </td>
            </tr>
            <tr v-for="device in filteredDevices" :key="device.id" class="hover:bg-white/5 transition-colors group">
              <td class="px-md py-4">
                <div class="flex items-center gap-sm">
                  <span class="material-symbols-outlined text-on-surface-variant text-[20px]">precision_manufacturing</span>
                  <span class="font-data-md text-on-surface tracking-wider font-bold">{{ device.id }}</span>
                </div>
              </td>
              <td class="px-md py-4">
                <div class="flex justify-center">
                  <div class="flex items-center gap-xs bg-surface-container-high px-sm py-[2px] rounded-full border border-white/5">
                    <span class="w-1.5 h-1.5 rounded-full shadow-[0_0_8px]" 
                          :class="[
                            statusColor(device.status, 'bg'),
                            device.status !== 'offline' ? 'status-pulse' : ''
                          ]"></span>
                    <span class="font-label-caps text-[10px] uppercase text-on-surface-variant">{{ device.status }}</span>
                  </div>
                </div>
              </td>
              <td class="px-md py-4 text-center">
                <span class="font-data-md" :class="device.status === 'offline' ? 'text-on-surface-variant' : 'text-on-surface'">{{ device.temp }}°C</span>
              </td>
              <td class="px-md py-4 text-center">
                <span class="font-data-md" :class="device.status === 'offline' ? 'text-on-surface-variant' : 'text-on-surface'">{{ device.hum }}%</span>
              </td>
              <td class="px-md py-4">
                <div class="flex items-center gap-xs text-on-surface-variant">
                  <span class="material-symbols-outlined text-[16px]">location_on</span>
                  <span class="font-label-caps text-[10px]">{{ device.sector }}</span>
                </div>
              </td>
              <td class="px-md py-4 text-right">
                <div class="flex justify-end items-center gap-sm">
                  <button @click="togglePower(device)" 
                          class="p-2 rounded-lg border transition-all text-on-surface-variant hover:border-primary/50 hover:text-primary active:scale-90"
                          :title="device.status === 'offline' ? 'Start Node' : 'Kill Switch'">
                    <span class="material-symbols-outlined text-[20px]">{{ device.status === 'offline' ? 'power_settings_new' : 'power_off' }}</span>
                  </button>
                  <router-link :to="`/devices/configure/${device.dbId}`" class="p-2 hover:bg-primary/20 hover:text-primary rounded-lg transition-all text-on-surface-variant active:scale-90" title="Configure Node">
                    <span class="material-symbols-outlined text-[20px]">settings</span>
                  </router-link>
                  <button @click="deleteDevice(device.dbId)" class="p-2 hover:bg-error/20 hover:text-error rounded-lg transition-all text-on-surface-variant active:scale-90" title="Decommission Node">
                    <span class="material-symbols-outlined text-[20px]">delete_forever</span>
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Pagination Controls -->
    <div class="mt-md flex items-center justify-between bg-surface-container-low/30 p-md rounded-xl border border-white/5">
      <div class="text-on-surface-variant font-label-caps text-[11px]">
        Showing {{ (currentPage - 1) * pageSize + 1 }} to {{ Math.min(currentPage * pageSize, totalItems) }} of {{ totalItems }} nodes
      </div>
      <div class="flex items-center gap-xs">
        <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1" class="p-2 hover:bg-white/5 rounded-lg disabled:opacity-30 disabled:hover:bg-transparent transition-all">
          <span class="material-symbols-outlined text-[20px]">chevron_left</span>
        </button>
        <div class="flex items-center gap-xs">
          <button v-for="page in totalPages" :key="page" @click="changePage(page)" class="w-8 h-8 rounded-lg font-data-md text-[12px] transition-all" :class="currentPage === page ? 'bg-primary text-on-primary' : 'hover:bg-white/5 text-on-surface-variant'">
            {{ page }}
          </button>
        </div>
        <button @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages" class="p-2 hover:bg-white/5 rounded-lg disabled:opacity-30 disabled:hover:bg-transparent transition-all">
          <span class="material-symbols-outlined text-[20px]">chevron_right</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { deviceApi } from '../api/devices';

const devices = ref([]);
const searchQuery = ref('');
const isLoading = ref(false);
const error = ref('');

// Pagination
const currentPage = ref(1);
const pageSize = ref(10);
const totalItems = ref(0);
const totalPages = computed(() => Math.ceil(totalItems.value / pageSize.value));

const changePage = (page) => {
    if (page < 1 || page > totalPages.value) return;
    currentPage.value = page;
    fetchDevices();
};

const filteredDevices = computed(() => {
    if (!searchQuery.value) return devices.value;
    const query = searchQuery.value.toLowerCase();
    return devices.value.filter(d => 
        d.id.toLowerCase().includes(query) || 
        d.sector.toLowerCase().includes(query) ||
        d.status.toLowerCase().includes(query)
    );
});

const fetchDevices = async () => {
    isLoading.value = true;
    error.value = '';
    try {
        const response = await deviceApi.getDevices(currentPage.value, pageSize.value);
        totalItems.value = response.total;
        const data = response.data;
        devices.value = data.map(d => ({
            id: d.identifier || d.Identifier,
            dbId: d.ID || d.id,
            status: (d.status || d.Status || 'offline').toLowerCase(),
            temp: d.temperature || d.Temperature || '--.-',
            hum: d.humidity || d.Humidity || '--',
            sector: d.Sector?.name || d.sector?.name || 'Unassigned'
        }));
    } catch (err) {
        console.error('Failed to fetch devices:', err);
        error.value = 'Failed to synchronize with industrial nodes.';
    } finally {
        isLoading.value = false;
    }
};

const statusColor = (status, prop) => {
  if (status === 'online') return prop === 'bg' ? 'bg-primary' : 'text-primary';
  if (status === 'warning') return prop === 'bg' ? 'bg-tertiary' : 'text-tertiary';
  return prop === 'bg' ? 'bg-error' : 'text-error';
}

const togglePower = async (device) => {
  const newStatus = device.status === 'offline' ? 'online' : 'offline';
  try {
      await deviceApi.updateDevice(device.dbId, {
          status: newStatus,
          temperature: newStatus === 'online' ? parseFloat((20 + Math.random() * 10).toFixed(1)) : 0,
          humidity: newStatus === 'online' ? parseFloat((30 + Math.random() * 30).toFixed(1)) : 0
      });
      fetchDevices();
  } catch (err) {
      console.error('Failed to toggle power:', err);
      alert('Node communication timeout: Check master bridge.');
  }
}

const deleteDevice = async (id) => {
    if (!confirm('Are you sure you want to decommission this hardware node?')) return
    try {
        await deviceApi.deleteDevice(id)
        devices.value = devices.value.filter(d => d.dbId !== id)
    } catch (err) {
        console.error('Failed to delete device:', err)
        alert('Decommission protocol failed: Node remains active.')
    }
}

onMounted(() => {
    fetchDevices();
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
.status-pulse {
    animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}
@keyframes pulse {
    0%, 100% { opacity: 1; }
    50% { opacity: .4; }
}
.custom-scrollbar::-webkit-scrollbar {
  height: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.02);
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
}
</style>
