<template>
  <div class="px-margin-mobile md:px-0 max-w-2xl mx-auto">
    <!-- Header Section -->
    <div class="mb-lg">
      <div class="flex items-center gap-sm mb-md">
        <router-link to="/devices" class="p-2 -ml-2 text-on-surface-variant transition-all duration-200 active:scale-95 hover:bg-surface-container-high rounded-lg">
          <span class="material-symbols-outlined">arrow_back</span>
        </router-link>
        <span class="font-label-caps text-label-caps text-on-surface-variant">BACK TO NODES</span>
      </div>
      <h2 class="font-display-lg text-display-lg text-on-surface">{{ isDeviceEdit ? 'Configure Node' : 'Provision New Node' }}</h2>
      <p class="text-on-surface-variant font-body-md mt-base">
        {{ isDeviceEdit ? 'Update hardware parameters and operational status for this industrial node.' : 'Register a new hardware node into the industrial network matrix.' }}
      </p>
    </div>

    <!-- Form Content -->
    <form class="space-y-gutter" @submit.prevent="submitForm">
      <!-- Identity Fields -->
      <section class="glass-panel p-md rounded-xl space-y-gutter">
        <div class="space-y-xs">
          <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Node Identifier</label>
          <div class="relative">
            <input v-model="form.identifier" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary focus:ring-1 focus:ring-primary outline-none" placeholder="e.g., NODE-2024-X" type="text" required/>
            <span class="material-symbols-outlined absolute right-3 top-3 text-on-surface-variant text-sm">precision_manufacturing</span>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-gutter">
          <div class="space-y-xs">
            <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Sector Assignment</label>
            <select v-model="form.sector_id" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary outline-none appearance-none">
              <option :value="0" disabled>Select Sector</option>
              <option v-for="sector in sectors" :key="sector.id" :value="sector.id">
                {{ sector.name }}
              </option>
            </select>
          </div>
          <div class="space-y-xs">
            <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Operational Status</label>
            <select v-model="form.status" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md transition-all focus:bg-surface-container focus:border-primary outline-none">
              <option value="online">Online</option>
              <option value="warning">Warning</option>
              <option value="offline">Offline</option>
            </select>
          </div>
        </div>
      </section>

      <!-- Telemetry Defaults (Only for simulated data/init) -->
      <section class="glass-panel p-md rounded-xl space-y-gutter" v-if="!isDeviceEdit">
        <p class="font-title-sm text-on-surface mb-sm">Initial Telemetry Calibration</p>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-gutter">
          <div class="space-y-xs">
            <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Temp (°C)</label>
            <input v-model.number="form.temperature" step="0.1" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md outline-none" type="number"/>
          </div>
          <div class="space-y-xs">
            <label class="font-label-caps text-label-caps text-on-surface-variant block uppercase tracking-wider">Humidity (%)</label>
            <input v-model.number="form.humidity" step="0.1" class="w-full bg-surface-container-lowest border border-outline-variant rounded-lg px-md py-3 text-on-surface font-data-md outline-none" type="number"/>
          </div>
        </div>
      </section>

      <!-- Action Buttons -->
      <div class="flex flex-col sm:flex-row gap-md pt-lg">
        <button class="flex-1 bg-primary text-on-primary font-title-sm text-title-sm py-4 rounded-xl flex items-center justify-center gap-sm transition-all duration-200 hover:brightness-110 active:scale-[0.98] shadow-lg shadow-primary/10" type="submit" :disabled="isSubmitting">
          <span v-if="isSubmitting" class="material-symbols-outlined animate-spin">refresh</span>
          <span v-else-if="isSuccess" class="material-symbols-outlined">check_circle</span>
          <span v-else class="material-symbols-outlined">save</span>
          {{ submitLabel }}
        </button>
        <router-link to="/devices" class="px-xl border border-outline-variant text-on-surface-variant font-title-sm text-title-sm py-4 rounded-xl transition-all hover:bg-surface-container-high active:scale-95 text-center">
          Cancel
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { deviceApi } from '../api/devices'
import { sectorApi } from '../api/sectors'

const router = useRouter()
const route = useRoute()

const isSubmitting = ref(false)
const isSuccess = ref(false)
const isDeviceEdit = computed(() => !!route.params.id)

const sectors = ref([])
const form = ref({
  identifier: '',
  sector_id: 0,
  status: 'online',
  temperature: 20.0,
  humidity: 45.0
})

const fetchSectors = async () => {
    try {
        const response = await sectorApi.getSectors(1, 100)
        sectors.value = response.data.map(s => ({
            id: s.ID || s.id,
            name: s.name || s.Name
        }))
        if (sectors.value.length > 0 && !isDeviceEdit.value) {
            form.value.sector_id = sectors.value[0].id
        }
    } catch (err) {
        console.error('Failed to load sectors:', err)
    }
}

const fetchDeviceData = async () => {
    if (!isDeviceEdit.value) return
    try {
        const device = await deviceApi.getDevice(route.params.id)
        form.value = {
            identifier: device.identifier || device.Identifier || '',
            sector_id: device.sector_id || device.SectorID || 0,
            status: (device.status || device.Status || 'online').toLowerCase(),
            temperature: device.temperature || device.Temperature || 0,
            humidity: device.humidity || device.Humidity || 0
        }
    } catch (err) {
        console.error('Failed to load node data:', err)
    }
}

const submitLabel = computed(() => {
  if (isSuccess.value) return isDeviceEdit.value ? 'Node Updated' : 'Node Provisioned'
  if (isSubmitting.value) return 'Syncing...'
  return isDeviceEdit.value ? 'Update Configuration' : 'Provision Node'
})

const submitForm = async () => {
    isSubmitting.value = true
    try {
        if (isDeviceEdit.value) {
            await deviceApi.updateDevice(route.params.id, form.value)
        } else {
            await deviceApi.createDevice(form.value)
        }
        isSuccess.value = true
        setTimeout(() => router.push('/devices'), 1500)
    } catch (err) {
        console.error('Handshake failed:', err)
        alert('Node registration failed: Check identifier uniqueness.')
    } finally {
        isSubmitting.value = false
    }
}

onMounted(async () => {
    await fetchSectors()
    fetchDeviceData()
})
</script>

<style scoped>
.glass-panel {
    background: rgba(23, 31, 51, 0.7);
    backdrop-filter: blur(12px);
    border: 1px solid rgba(255, 255, 255, 0.08);
}
</style>
